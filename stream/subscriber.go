package stream

import (
	"context"
	"errors"
)

var ErrSubscriberClosed = errors.New("subscriber is closed")

var ErrSubscriberAlreadySubscribed = errors.New("subscriber is already subscribed")

// Subscriber is message queue subscriber
type Subscriber interface {
	Subscribe(ctx context.Context, topic string, msgC chan<- *Message, errC chan<- error) error
	Close(ctx context.Context) error
}
