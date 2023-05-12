package message

import (
	"context"
	"errors"
)

type Ack func(ctx context.Context, msg *Message) error

type Nack func(ctx context.Context, msg *Message) error

type Handler interface {
	Handle(ctx context.Context, msg *Message, ack Ack, nack Nack) error
}

type HandlerFunc func(ctx context.Context, msg *Message, ack Ack, nack Nack) error

func (f HandlerFunc) Handle(ctx context.Context, msg *Message, ack Ack, nack Nack) error {
	return f(ctx, msg, ack, nack)
}

type PublishHandler struct {
	Publisher    Publisher
	PublishTopic string
	Handler      func(ctx context.Context, msg *Message) ([]*Message, error)
}

func (h PublishHandler) Handle(ctx context.Context, msg *Message, ack Ack, nack Nack) error {
	if h.Publisher == nil {
		return errors.New("publisher is nil")
	}
	msgs, err := h.Handler(ctx, msg)
	if err != nil {
		return nack(ctx, msg)
	}
	_, err = h.Publisher.Publish(ctx, h.PublishTopic, msgs...)
	if err != nil {
		return nack(ctx, msg)
	}
	return ack(ctx, msg)
}
