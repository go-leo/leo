package stream

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"strings"

	"golang.org/x/exp/slices"
)

var ErrSubscriberClosed = errors.New("subscriber is closed")

var ErrSubscriberAlreadySubscribed = errors.New("subscriber is already subscribed")

// Subscriber is message queue subscriber
type Subscriber interface {
	Topic() string
	Queue() string
	Subscribe(ctx context.Context, msgC chan<- *Message, errC chan<- error) error
	Close(ctx context.Context) error
}

type multiSubscriber struct {
	subscribers []Subscriber
}

func (sub *multiSubscriber) Topic() string {
	var topics []string
	for _, publisher := range sub.subscribers {
		if slices.Contains(topics, publisher.Topic()) {
			continue
		}
		topics = append(topics, publisher.Topic())
	}
	return strings.Join(topics, ",")
}

func (sub *multiSubscriber) Queue() string {
	var queues []string
	for _, publisher := range sub.subscribers {
		if slices.Contains(queues, publisher.Queue()) {
			continue
		}
		queues = append(queues, publisher.Queue())
	}
	return strings.Join(queues, ",")
}

func (sub *multiSubscriber) Subscribe(ctx context.Context, msgC chan<- *Message, errC chan<- error) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, w := range sub.subscribers {
		eg.Go(func() error {
			return w.Subscribe(ctx, msgC, errC)
		})
	}
	return eg.Wait()
}

func (sub *multiSubscriber) Close(ctx context.Context) error {
	for _, w := range sub.subscribers {
		if err := w.Close(ctx); err != nil {
			return err
		}
	}
	return nil
}

func MultiSubscriber(subscribers ...Subscriber) Subscriber {
	allSubscribers := make([]Subscriber, 0, len(subscribers))
	for _, w := range subscribers {
		if mw, ok := w.(*multiSubscriber); ok {
			allSubscribers = append(allSubscribers, mw.subscribers...)
		} else {
			allSubscribers = append(allSubscribers, w)
		}
	}
	return &multiSubscriber{subscribers: allSubscribers}
}
