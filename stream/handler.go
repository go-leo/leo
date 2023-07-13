package stream

import (
	"context"
)

type PubSubHandler interface {
	Subscriber() (Subscriber, error)
	Handle(ctx context.Context, msg *Message) ([]*Message, error)
	Publisher() (Publisher, error)
}

type Handler interface {
	Subscriber() (Subscriber, error)
	Handle(ctx context.Context, msg *Message) error
}
