package kafka

import (
	"codeup.aliyun.com/qimao/leo/leo/stream"
	"context"
	"errors"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"sync/atomic"
)

var _ stream.Subscriber = new(Subscriber)

type Subscriber struct {
	o          *options
	consumer   *kafka.Consumer
	topic      string
	subscribed atomic.Bool
	closed     atomic.Bool
	closeC     chan struct{}
}

func (sub *Subscriber) Topic() string {
	return sub.topic
}

func (sub *Subscriber) Queue() string {
	return "kafka"
}

func (sub *Subscriber) Subscribe(ctx context.Context, msgC chan<- *stream.Message, errC chan<- error) error {
	if sub.closed.Load() {
		return stream.ErrSubscriberClosed
	}
	if !sub.subscribed.CompareAndSwap(false, true) {
		return stream.ErrSubscriberAlreadySubscribed
	}
	if err := sub.consumer.Subscribe(sub.topic, sub.o.RebalanceCb); err != nil {
		return err
	}
	go sub.consumeMsg(ctx, msgC, errC)
	return nil
}

func (sub *Subscriber) Close(ctx context.Context) error {
	if !sub.closed.CompareAndSwap(false, true) {
		return nil
	}
	close(sub.closeC)
	return sub.consumer.Close()
}

func (sub *Subscriber) consumeMsg(ctx context.Context, msgC chan<- *stream.Message, errC chan<- error) {
	defer func() {
		if _, err := sub.consumer.Commit(); err != nil {
			errC <- err
		}
		if err := sub.consumer.Unsubscribe(); err != nil {
			errC <- err
		}
	}()
	for {
		select {
		case <-ctx.Done():
			return
		case <-sub.closeC:
			return
		default:
			ev := sub.consumer.Poll(100)
			if ev == nil {
				continue
			}
			switch e := ev.(type) {
			case *kafka.Message:
				if e.TopicPartition.Error != nil {
					err := fmt.Errorf(
						"partition specific error, topic:%s, partition: %d, offset: %d, error: %w",
						*e.TopicPartition.Topic,
						e.TopicPartition.Partition,
						e.TopicPartition.Offset,
						e.TopicPartition.Error,
					)
					errC <- err
					continue
				}
				sub.handleMsg(ctx, e, msgC, errC)
			case kafka.Error:
				if e.IsTimeout() {
					continue
				}
				errC <- e
			default:
				continue
			}
		}
	}

}

func (sub *Subscriber) handleMsg(ctx context.Context, kafkaMsg *kafka.Message, msgC chan<- *stream.Message, errC chan<- error) {
	msg, err := sub.o.Marshaler.Unmarshal(kafkaMsg)
	if err != nil {
		errC <- fmt.Errorf("failed to unmarshal kafka message: %w", err)
		return
	}

	ackC := make(chan struct{})
	msg.NotifyAck(ackC, func(ctx context.Context, msg *stream.Message) (any, error) {
		return sub.consumer.CommitMessage(kafkaMsg)
	})

	nackC := make(chan struct{})
	msg.NotifyNack(nackC, func(ctx context.Context, msg *stream.Message) (any, error) {
		return nil, nil
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

func NewSubscriber(topic string, factory func() (*kafka.Consumer, error), opts ...Option) (*Subscriber, error) {
	if factory == nil {
		return nil, errors.New("factory is nil")
	}
	consumer, err := factory()
	if err != nil {
		return nil, err
	}
	o := &options{}
	o.apply(opts...)
	o.init()
	return &Subscriber{consumer: consumer, topic: topic, o: o}, nil
}
