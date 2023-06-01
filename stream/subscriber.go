package stream

import (
	"context"
	"errors"
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
