package amqp

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"

	"codeup.aliyun.com/qimao/leo/leo/stream"

	"github.com/rabbitmq/amqp091-go"
)

var _ stream.Subscriber = new(Subscriber)

type Subscriber struct {
	o          *options
	connect    *amqp091.Connection
	channel    *amqp091.Channel
	topic      string
	subscribed atomic.Bool
	closed     atomic.Bool
	closeC     chan struct{}
	stopC      chan struct{}
	wg         sync.WaitGroup
}

func (sub *Subscriber) Topic() string {
	return sub.topic
}

func (sub *Subscriber) Queue() string {
	return "amqp"
}

func (sub *Subscriber) Subscribe(ctx context.Context, msgC chan<- *stream.Message, errC chan<- error) error {
	if sub.closed.Load() {
		return stream.ErrSubscriberClosed
	}
	if !sub.subscribed.CompareAndSwap(false, true) {
		return stream.ErrSubscriberAlreadySubscribed
	}
	queue := sub.o.QueueName(sub.topic)
	deliveryC, err := sub.channel.Consume(queue, sub.o.ConsumeOption.Tag, sub.o.ConsumeOption.AutoAck, sub.o.ConsumeOption.Exclusive, sub.o.ConsumeOption.NoLocal, sub.o.ConsumeOption.NoWait, sub.o.ConsumeOption.Table)
	if err != nil {
		return err
	}
	defer func() {
		defer close(sub.stopC)
		sub.wg.Wait()
		err := sub.channel.Cancel(sub.o.ConsumeOption.Tag, sub.o.ConsumeOption.NoWait)
		if err == nil {
			return
		}
		if errC != nil {
			errC <- err
		} else {
			sub.o.Logger.Error(err)
		}
	}()
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-sub.closeC:
			return nil
		case delivery, ok := <-deliveryC:
			if !ok {
				return nil
			}
			sub.handleMsg(ctx, delivery, msgC, errC)
		}
	}
}

func (sub *Subscriber) Close(ctx context.Context) error {
	if !sub.closed.CompareAndSwap(false, true) {
		return nil
	}
	close(sub.closeC)
	select {
	case <-ctx.Done():
	case <-sub.stopC:
	}
	return errors.Join(sub.channel.Close(), sub.connect.Close())
}

func (sub *Subscriber) handleMsg(ctx context.Context, delivery amqp091.Delivery, msgC chan<- *stream.Message, errC chan<- error) {
	sub.wg.Add(1)
	defer sub.wg.Done()
	msg, err := sub.o.Marshaller.Unmarshal(sub.topic, delivery)
	if err != nil {
		errC <- fmt.Errorf("failed to unmarshal amqp091 delivery: %w", err)
		return
	}
	if sub.o.OnMessageReceived != nil {
		msg = sub.o.OnMessageReceived(msg, delivery)
	}

	ackC := make(chan struct{})
	stream.NotifyAck(msg, ackC, func(ctx context.Context, msg *stream.Message) (any, error) {
		if sub.o.ConsumeOption.AutoAck {
			return nil, nil
		}
		return nil, delivery.Ack(sub.o.AckOption.Multiple)
	})

	nackC := make(chan struct{})
	stream.NotifyNack(msg, nackC, func(ctx context.Context, msg *stream.Message) (any, error) {
		if sub.o.ConsumeOption.AutoAck {
			return nil, nil
		}
		return nil, delivery.Nack(sub.o.AckOption.Multiple, sub.o.AckOption.Requeue)
	})

	select {
	case <-ctx.Done():
		return
	case <-sub.closeC:
		return
	case msgC <- msg:
	}

	select {
	case <-ctx.Done():
		return
	case <-sub.closeC:
		return
	case <-ackC:
		return
	case <-nackC:
		if sub.o.NackHandler != nil {
			sub.o.NackHandler(msg)
		}
		return
	}
}

func NewSubscriber(topic string, factory func(topic string) (*amqp091.Connection, error), opts ...Option) (*Subscriber, error) {
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
	return &Subscriber{
		o:          o,
		connect:    connect,
		channel:    channel,
		topic:      topic,
		subscribed: atomic.Bool{},
		closed:     atomic.Bool{},
		closeC:     make(chan struct{}),
		stopC:      make(chan struct{}),
		wg:         sync.WaitGroup{},
	}, nil
}
