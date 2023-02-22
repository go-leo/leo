package stream

import "context"

// Subscriber is message queue subscriber
type Subscriber interface {
	Topic() string
	Subscribe(ctx context.Context) (<-chan Message, error)
	Close() error
}
