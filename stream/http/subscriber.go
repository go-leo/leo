package http

//
//import (
//	"context"
//	"errors"
//	"fmt"
//	"sync"
//	"sync/atomic"
//
//	"github.com/go-leo/gox/errorx"
//
//	"codeup.aliyun.com/qimao/leo/leo/stream"
//
//	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
//)
//
//var _ stream.Subscriber = new(Subscriber)
//
//type Subscriber struct {
//	o          *options
//	consumer   *kafka.Consumer
//	topic      string
//	subscribed atomic.Bool
//	closed     atomic.Bool
//	closeC     chan struct{}
//	stopC      chan struct{}
//	wg         sync.WaitGroup
//}
//
//func (sub *Subscriber) Topic() string {
//	return sub.topic
//}
//
//func (sub *Subscriber) Queue() string {
//	return "kafka"
//}
//
//func (sub *Subscriber) Subscribe(ctx context.Context, msgC chan<- *stream.Message, errC chan<- error) error {
//	if sub.closed.Load() {
//		return stream.ErrSubscriberClosed
//	}
//	if !sub.subscribed.CompareAndSwap(false, true) {
//		return stream.ErrSubscriberAlreadySubscribed
//	}
//	if err := sub.consumer.Subscribe(sub.topic, sub.o.RebalanceCb); err != nil {
//		return err
//	}
//	defer func() {
//		defer close(sub.stopC)
//		sub.wg.Wait()
//		err := errors.Join(errorx.Concern(sub.consumer.Commit()), sub.consumer.Unsubscribe())
//		if err == nil {
//			return
//		}
//		if errC != nil {
//			errC <- err
//			return
//		}
//		sub.o.Logger.Error(err)
//	}()
//	for {
//		select {
//		case <-ctx.Done():
//			return nil
//		case <-sub.closeC:
//			return nil
//		default:
//			ev := sub.consumer.Poll(int(sub.o.PollTimeout.Milliseconds()))
//			if ev == nil {
//				continue
//			}
//			switch e := ev.(type) {
//			case *kafka.Message:
//				if e.TopicPartition.Error != nil {
//					err := fmt.Errorf(
//						"partition specific error, topic:%s, partition: %d, offset: %d, error: %w",
//						*e.TopicPartition.Topic,
//						e.TopicPartition.Partition,
//						e.TopicPartition.Offset,
//						e.TopicPartition.Error,
//					)
//					if errC != nil {
//						errC <- err
//						continue
//					}
//					sub.o.Logger.Error(err)
//					continue
//				}
//				sub.handleMsg(ctx, e, msgC, errC)
//			case kafka.Error:
//				if e.IsTimeout() {
//					continue
//				}
//				if e.IsRetriable() {
//					continue
//				}
//				if e.IsFatal() {
//					return e
//				}
//				if errC != nil {
//					errC <- e
//					continue
//				}
//				sub.o.Logger.Error(e)
//			default:
//				continue
//			}
//		}
//	}
//}
//
//func (sub *Subscriber) Close(ctx context.Context) error {
//	if !sub.closed.CompareAndSwap(false, true) {
//		return nil
//	}
//	close(sub.closeC)
//	select {
//	case <-ctx.Done():
//	case <-sub.stopC:
//	}
//	return sub.consumer.Close()
//}
//
//func (sub *Subscriber) handleMsg(ctx context.Context, kafkaMsg *kafka.Message, msgC chan<- *stream.Message, errC chan<- error) {
//	sub.wg.Add(1)
//	defer sub.wg.Done()
//	msg, err := sub.o.Marshaller.Unmarshal(sub.topic, kafkaMsg)
//	if err != nil {
//		errC <- fmt.Errorf("failed to unmarshal kafka message: %w", err)
//		return
//	}
//	if sub.o.OnMessageReceived != nil {
//		msg = sub.o.OnMessageReceived(msg, kafkaMsg)
//	}
//
//	ackC := make(chan struct{})
//	stream.NotifyAck(msg, ackC, func(ctx context.Context, msg *stream.Message) (any, error) {
//		partitions, err := sub.consumer.CommitMessage(kafkaMsg)
//		return partitions, err
//	})
//
//	nackC := make(chan struct{})
//	stream.NotifyNack(msg, nackC, func(ctx context.Context, msg *stream.Message) (any, error) {
//		return nil, nil
//	})
//
//	select {
//	case <-ctx.Done():
//		return
//	case <-sub.closeC:
//		return
//	case msgC <- msg:
//	}
//
//	select {
//	case <-ctx.Done():
//		return
//	case <-sub.closeC:
//		return
//	case <-ackC:
//		return
//	case <-nackC:
//		if sub.o.NackHandler != nil {
//			sub.o.NackHandler(msg)
//		}
//		return
//	}
//}
//
//func NewSubscriber(topic string, factory func() (*kafka.Consumer, error), opts ...Option) (*Subscriber, error) {
//	if factory == nil {
//		return nil, errors.New("factory is nil")
//	}
//	consumer, err := factory()
//	if err != nil {
//		return nil, err
//	}
//	o := &options{}
//	o.apply(opts...)
//	o.init()
//	return &Subscriber{
//		o:          o,
//		consumer:   consumer,
//		topic:      topic,
//		subscribed: atomic.Bool{},
//		closed:     atomic.Bool{},
//		closeC:     make(chan struct{}),
//		stopC:      make(chan struct{}),
//		wg:         sync.WaitGroup{},
//	}, nil
//}
