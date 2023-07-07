package amqp

import (
	"codeup.aliyun.com/qimao/leo/leo/stream"
	"context"
	"errors"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"sync"
	"sync/atomic"
)

var _ stream.Publisher = new(Publisher)

type Publisher struct {
	o       *options
	wg      sync.WaitGroup
	closed  atomic.Bool
	topic   string
	channel *amqp091.Channel
	connect *amqp091.Connection
}

func (pub *Publisher) Topic() string {
	return pub.topic
}

func (pub *Publisher) Queue() string {
	return "amqp"
}

func (pub *Publisher) Publish(ctx context.Context, messages ...*stream.Message) (stream.PublishResult, error) {
	if len(messages) == 0 {
		return nil, nil
	}
	if pub.closed.Load() {
		return nil, stream.ErrPublisherClosed
	}
	pub.wg.Add(1)
	defer pub.wg.Done()
	return pub.publishMessages(ctx, messages, pub.o.ExchangeName(pub.topic), pub.o.RoutingKeys(pub.topic))
}

func (pub *Publisher) publishMessages(ctx context.Context, messages []*stream.Message, exchangeName string, routingKeys []string) (stream.PublishResult, error) {
	var results stream.PublishResults
	for _, routingKey := range routingKeys {
		for _, msg := range messages {
			publishing, err := pub.convertMessage(ctx, msg)
			if err != nil {
				return nil, err
			}
			result, err := pub.publishing(ctx, publishing, exchangeName, routingKey)
			if err != nil {
				return nil, err
			}
			results = append(results, result)
		}
	}
	return results, nil
}

func (pub *Publisher) convertMessage(ctx context.Context, msg *stream.Message) (amqp091.Publishing, error) {
	publishing, err := pub.o.Marshaller.Marshal(ctx, pub.topic, msg)
	if err != nil {
		return amqp091.Publishing{}, err
	}
	if pub.o.PublishingDecorator != nil {
		publishing = pub.o.PublishingDecorator(msg, publishing)
	}
	publishing.DeliveryMode = pub.o.PublishOption.DeliveryMode
	return publishing, nil
}

func (pub *Publisher) publishing(ctx context.Context, publishing amqp091.Publishing, exchangeName string, routingKey string) (*PublishResult, error) {
	if pub.o.ConfirmOption == nil {
		if err := pub.channel.PublishWithContext(ctx, exchangeName, routingKey, pub.o.PublishOption.Mandatory, pub.o.PublishOption.Immediate, publishing); err != nil {
			return nil, err
		}
		return &PublishResult{Msg: publishing, ExchangeName: exchangeName, RoutingKey: routingKey}, nil
	}
	confirmation, err := pub.channel.PublishWithDeferredConfirmWithContext(ctx, exchangeName, routingKey, pub.o.PublishOption.Mandatory, pub.o.PublishOption.Immediate, publishing)
	if err != nil {
		return nil, err
	}
	acked, err := confirmation.WaitContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("waiting for confirmation, %w", err)
	}
	return &PublishResult{Msg: publishing, Acked: acked, ExchangeName: exchangeName, RoutingKey: routingKey}, nil
}

func (pub *Publisher) Close(_ context.Context) error {
	if !pub.closed.CompareAndSwap(false, true) {
		return nil
	}
	pub.wg.Wait()
	return errors.Join(pub.channel.Close(), pub.connect.Close())
}

type PublishResult struct {
	Msg          amqp091.Publishing
	Acked        bool
	ExchangeName string
	RoutingKey   string
}

func (p PublishResult) String() string {
	return p.ExchangeName + "/" + p.RoutingKey
}

func NewPublisher(topic string, factory func(topic string) (*amqp091.Connection, error), opts ...Option) (*Publisher, error) {
	if factory == nil {
		return nil, errors.New("factory is nil")
	}
	connect, err := factory(topic)
	if err != nil {
		return nil, err
	}
	channel, err := connect.Channel()
	if err != nil {
		return nil, err
	}
	o := &options{}
	o.apply(opts...)
	if err := o.init(topic, channel); err != nil {
		return nil, err
	}
	pub := &Publisher{
		o:       o,
		connect: connect,
		channel: channel,
		topic:   topic,
	}
	return pub, nil
}
