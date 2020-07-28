package sjmq

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/google/uuid"
	"github.com/kubemq-io/kubemq-go"
)

func NewSjmqClient(host string, ctx context.Context) (*Client, error) {
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
	}
	return sjmqClient, err
}

type Client struct {
	KubemqClient *kubemq.Client
	Ctx          context.Context
}

func (this *Client) SendEvent(event interface{}) error {
	channel := reflect.TypeOf(event).Name()
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	_, err = this.KubemqClient.ES().
		SetChannel(channel).
		SetBody(data).
		Send(this.Ctx)
	return err
}

func (this *Client) SubscribeEvent(sjEvent interface{}, group string, handler func(*kubemq.EventStoreReceive)) error {
	errCh := make(chan error)
	channel := reflect.TypeOf(sjEvent).Name()
	eventsCh, err := this.KubemqClient.SubscribeToEventsStore(this.Ctx, channel, group, errCh, kubemq.StartFromNewEvents())
	if err != nil {
		return err
	}

	for {
		select {
		case err := <-errCh:
			return err
		case event := <-eventsCh:
			go handler(event)
		case <-this.Ctx.Done():
			return nil
		}
	}
}

func (this *Client) Close() error {
	return this.KubemqClient.Close()
}
