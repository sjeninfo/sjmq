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
	Sender    *kubemq.Client
	Receivers []*kubemq.Client
	Ctx       context.Context
	Host      string
	Group     string
}

func newMqClient(ctx context.Context, host string) (*kubemq.Client, error) {
	clientId := uuid.New()
	return kubemq.NewClient(ctx,
		kubemq.WithAddress(host, 50000),
		kubemq.WithClientId(clientId.String()),
		kubemq.WithTransportType(kubemq.TransportTypeGRPC),
		kubemq.WithAutoReconnect(true),
		kubemq.WithMaxReconnects(2))
}

func NewSjmqClient(ctx context.Context, host string, group string) (*Client, error) {
	sender, err := newMqClient(ctx, host)
	if err != nil {
		return nil, err
	}
	sjmqClient := &Client{
		Sender: sender,
		Ctx:    ctx,
		Group:  group,
	}
	return sjmqClient, err
}

func (c *Client) SendEvent(sjEvent interface{}) error {
	channel := getChannelName(sjEvent)
	data, err := json.Marshal(sjEvent)
	if err != nil {
		return err
	}
	_, err = c.Sender.ES().
		SetChannel(channel).
		SetBody(data).
		Send(c.Ctx)
	return err
}

func (c *Client) SubscribeEvent(sjEvent interface{}, handler func(*kubemq.EventStoreReceive), errCh chan error) {
	receiver, err := newMqClient(c.Ctx, c.Host)
	if err != nil {
		errCh <- err
		return
	}
	channel := getChannelName(sjEvent)
	eventsCh, err := receiver.SubscribeToEventsStore(c.Ctx, channel, c.Group, errCh, kubemq.StartFromFirstEvent())
	if err != nil {
		_ = receiver.Close()
		errCh <- err
		return
	}
	c.Receivers = append(c.Receivers, receiver)
	
	go func() {
		for {
			select {
			case <-errCh:
				return
			case event := <-eventsCh:
				go handler(event)
			case <-c.Ctx.Done():
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

func (c *Client) Close() {
	_ = c.Sender.Close()
	for _, receiver := range c.Receivers {
		_ = receiver.Close()
	}
}
