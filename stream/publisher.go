package stream

import "context"

type Publisher interface {
	Topic() string
	Publish(ctx context.Context, messages ...Message) error
	Close() error
}
