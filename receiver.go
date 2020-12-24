package sjmq

import (
	"context"
	"github.com/google/uuid"
	"github.com/kubemq-io/kubemq-go"
)

type Receiver struct {
	Host  string
	Group string
}

func NewReceiver(mqHost string, group string) *Receiver {
	r := &Receiver{
		Host:  mqHost,
		Group: group,
	}
	return r
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

func (r *Receiver) SubscribeEvent(c context.Context, sjEvent interface{}, handler func(*kubemq.EventStoreReceive)) error {
	receiver, err := newKubemqClient(c, r.Host)
	if err != nil {
		return err
	}
	defer receiver.Close()
	channel := getChannelName(sjEvent)
	errCh := make(chan error)
	eventsCh, err := receiver.SubscribeToEventsStore(c, channel, r.Group, errCh, kubemq.StartFromFirstEvent())
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case <-c.Done():
				return
			case <-errCh:
				return
			case event := <-eventsCh:
				go handler(event)
				// error handler
			}
		}
	}()
	return nil
}
