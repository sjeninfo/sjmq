package sjmq

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"reflect"
	"time"

	"github.com/google/uuid"
	"github.com/kubemq-io/kubemq-go"
	"gorm.io/gorm"
)

type EventConfig struct {
	Event   interface{}
	Handler interface{}
}

type Receiver struct {
	Host      string
	Group     string
	db        *gorm.DB
	inboxRepo *InboxRepository

	channelToType    map[string]reflect.Type
	channelToHandler map[string]reflect.Value
	notifyCh         chan bool
}

func NewReceiver(db *gorm.DB, mqHost string, group string) *Receiver {
	r := &Receiver{
		db:               db,
		inboxRepo:        NewInboxRepository(NewGormDBContext(db)),
		Host:             mqHost,
		Group:            group,
		channelToType:    make(map[string]reflect.Type),
		channelToHandler: make(map[string]reflect.Value),
		notifyCh:         make(chan bool, 1),
	}
	go r.eventHandler()
	return r
}

func (r *Receiver) notify() {
	r.notifyCh <- true
}

func (r *Receiver) SubscribeEvents(eventConfigs []EventConfig) {
	errorModel := func() error { return nil }
	for _, eventConfig := range eventConfigs {
		event := eventConfig.Event
		handler := eventConfig.Handler
		checker := reflect.FuncOf(
			[]reflect.Type{reflect.TypeOf(event)},
			[]reflect.Type{reflect.TypeOf(errorModel).Out(0)},
			false,
		)
		t := reflect.TypeOf(handler)
		if t != checker {
			panic("event and handler input type is not matching")
		}

		channel := getChannelName(event)
		r.channelToType[channel] = reflect.TypeOf(event)
		r.channelToHandler[channel] = reflect.ValueOf(handler)
		r.subscribeMqEvent(context.Background(), event)
	}
	r.notify()
}

func (r *Receiver) eventHandler() {
	enable := false
	fn := func() {
		for enable {
			func() {
				dbCtx := NewGormDBContext(r.db)
				inbox := NewInboxRepository(dbCtx)
				err := dbCtx.Begin()
				if err != nil {
					panic(err)
				}
				defer dbCtx.Commit()
				message, err := inbox.GetTopMessage()
				if err != nil {
					panic(err) //TODO:
				}
				if message == nil {
					enable = false
					return
				}

				t, ok := r.channelToType[message.Channel]
				if !ok {
					panic(errors.New("received an unsubscribed event"))
				}
				eventPtr := reflect.New(t).Interface()
				err = json.Unmarshal([]byte(message.Event), eventPtr)
				if err != nil {
					panic(err) //TODO: handle error
				}
				handler := r.channelToHandler[message.Channel]
				out := handler.Call([]reflect.Value{reflect.ValueOf(eventPtr).Elem()})
				outValue := out[0].Interface()
				if outValue != nil {
					err := outValue.(error)
					_ = inbox.AddTimes(message)
					panic(err)
					time.Sleep(time.Second)
					return
				}

				err = inbox.DeleteMessage(message)
				if err != nil {
					panic(err) //todo:
				}
			}()
		}
	}

	for {
		<-r.notifyCh
		if enable {
			continue
		}
		enable = true
		go fn()
	}
}

func (r *Receiver) subscribeMqEvent(c context.Context, sjEvent interface{}) error {
	client, err := newKubemqClient(c, r.Host)
	if err != nil {
		return err
	}
	//defer receiver.Close()
	channel := getChannelName(sjEvent)
	errCh := make(chan error)
	eventsCh, err := client.SubscribeToEventsStore(c, channel, r.Group, errCh, kubemq.StartFromFirstEvent())
	if err != nil {
		return err
	}

	go r.EventLisenter(eventsCh, errCh)
	return nil
}

func (r *Receiver) EventLisenter(eventCh <-chan *kubemq.EventStoreReceive, errCh <-chan error) {
	for {
		select {
		case <-errCh:
			return
		case event := <-eventCh:
			err := r.inboxRepo.InsertMessage(event.Id, event.Channel, string(event.Body))
			if err != nil {
				panic("") //todo
			}
			r.notify()
		}
	}
}

func newKubemqClient(ctx context.Context, host string) (*kubemq.Client, error) {
	clientId := uuid.New()
	return kubemq.NewClient(ctx,
		kubemq.WithAddress(host, 50000),
		kubemq.WithClientId(clientId.String()),
		kubemq.WithTransportType(kubemq.TransportTypeGRPC),
		kubemq.WithAutoReconnect(true),
		kubemq.WithMaxReconnects(2))
}
