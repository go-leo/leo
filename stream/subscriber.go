package stream

import (
	"context"
)

// Subscriber is message queue subscriber
type Subscriber interface {
	Subscribe(ctx context.Context, topic string, msgC chan<- *Message, errC chan<- error) error
	Close(ctx context.Context) error
}
