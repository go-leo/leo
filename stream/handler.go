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

type pubSubHandler struct {
	sub Subscriber
	pub Publisher
	h   func(ctx context.Context, msg *Message) ([]*Message, error)
}

func (p pubSubHandler) Subscriber() (Subscriber, error) {
	return p.sub, nil
}

func (p pubSubHandler) Handle(ctx context.Context, msg *Message) ([]*Message, error) {
	return p.h(ctx, msg)
}

func (p pubSubHandler) Publisher() (Publisher, error) {
	return p.pub, nil
}

type handler struct {
	sub Subscriber
	h   func(ctx context.Context, msg *Message) error
}

func (h handler) Subscriber() (Subscriber, error) {
	return h.sub, nil
}

func (h handler) Handle(ctx context.Context, msg *Message) error {
	return h.h(ctx, msg)
}

func NewPubSubHandler(sub Subscriber, pub Publisher, h func(ctx context.Context, msg *Message) ([]*Message, error)) PubSubHandler {
	return &pubSubHandler{sub: sub, pub: pub, h: h}
}

func NewHandler(sub Subscriber, h func(ctx context.Context, msg *Message) error) Handler {
	return &handler{sub: sub, h: h}
}
