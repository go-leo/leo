package stream

import "context"

type Publisher interface {
	Topic() string
	Publish(ctx context.Context, messages ...Message) (any, error)
	Close() error
}
