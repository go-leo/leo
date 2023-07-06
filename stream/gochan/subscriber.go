package gochan

import (
	"codeup.aliyun.com/qimao/leo/leo/stream"
	"context"
	"errors"
	"fmt"
	"sync/atomic"
)

var _ stream.Subscriber = new(Subscriber)

type Subscriber struct {
	o          *options
	topic      string
	subscribed atomic.Bool
	closed     atomic.Bool
	closeC     chan struct{}
	goChan     <-chan []byte
}

func (sub *Subscriber) Topic() string {
	return sub.topic
}

func (sub *Subscriber) Queue() string {
	return "gochan"
}

func (sub *Subscriber) Subscribe(ctx context.Context, msgC chan<- *stream.Message, errC chan<- error) error {
	if sub.closed.Load() {
		return stream.ErrSubscriberClosed
	}
	if !sub.subscribed.CompareAndSwap(false, true) {
		return stream.ErrSubscriberAlreadySubscribed
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-sub.closeC:
			return nil
		case goChanMsg := <-sub.goChan:
			sub.handleMsg(ctx, goChanMsg, msgC, errC)
		}
	}
}

func (sub *Subscriber) Close(ctx context.Context) error {
	if !sub.closed.CompareAndSwap(false, true) {
		return nil
	}
	close(sub.closeC)
	return nil
}

func (sub *Subscriber) handleMsg(ctx context.Context, goChanMsg []byte, msgC chan<- *stream.Message, errC chan<- error) {
	msg, err := sub.o.Marshaler.Unmarshal(goChanMsg)
	if err != nil {
		if errC != nil {
			errC <- fmt.Errorf("failed to unmarshal kafka message: %w", err)
			return
		}
		sub.o.Logger.Error("failed to unmarshal message, error: %w", err)
		return
	}

	ackC := make(chan struct{})
	stream.NotifyAck(msg, ackC, func(ctx context.Context, msg *stream.Message) (any, error) {
		return nil, nil
	})
	nackC := make(chan struct{})
	stream.NotifyNack(msg, nackC, func(ctx context.Context, msg *stream.Message) (any, error) {
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

func NewSubscriber(topic string, goChan <-chan []byte, opts ...Option) (*Subscriber, error) {
	if goChan == nil {
		return nil, errors.New("factory is nil")
	}
	o := &options{}
	o.apply(opts...)
	o.init()
	return &Subscriber{goChan: goChan, topic: topic, o: o}, nil
}
