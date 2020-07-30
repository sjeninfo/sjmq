package sjmq

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/google/uuid"
	"github.com/kubemq-io/kubemq-go"
)

type IClient interface {
	SendEvent(interface{}) error
}

type Client struct {
	KubemqClient *kubemq.Client
	Ctx          context.Context
	Group        string
}

func NewSjmqClient(host string, group string, ctx context.Context) (*Client, error) {
	clientId := uuid.New()
	kubemqClient, err := kubemq.NewClient(ctx,
		kubemq.WithAddress(host, 50000),
		kubemq.WithClientId(clientId.String()),
		kubemq.WithTransportType(kubemq.TransportTypeGRPC),
		kubemq.WithAutoReconnect(true),
		kubemq.WithMaxReconnects(2))
	if err != nil {
		return nil, err
	}
	sjmqClient := &Client{
		KubemqClient: kubemqClient,
		Ctx:          ctx,
		Group:        group,
	}
	return sjmqClient, err
}

func (this *Client) SendEvent(sjEvent interface{}) error {
	channel := getChannelName(sjEvent)
	data, err := json.Marshal(sjEvent)
	if err != nil {
		return err
	}
	_, err = this.KubemqClient.ES().
		SetChannel(channel).
		SetBody(data).
		Send(this.Ctx)
	return err
}

func (this *Client) SubscribeEvent(sjEvent interface{}, handler func(*kubemq.EventStoreReceive), errCh chan error) {
	channel := getChannelName(sjEvent)
	eventsCh, err := this.KubemqClient.SubscribeToEventsStore(this.Ctx, channel, this.Group, errCh, kubemq.StartFromFirstEvent())
	if err != nil {
		errCh <- err
		return
	}
	go func() {
		for {
			select {
			case <-errCh:
				return
			case event := <-eventsCh:
				go handler(event)
			case <-this.Ctx.Done():
				return
			}
		}
	}()
}

func getChannelName(v interface{}) string {
	var channel string
	if reflect.TypeOf(v).Kind() == reflect.Ptr {
		channel = reflect.TypeOf(v).Elem().Name()
	} else {
		channel = reflect.TypeOf(v).Name()
	}
	return channel
}

func (this *Client) Close() error {
	return this.KubemqClient.Close()
}
