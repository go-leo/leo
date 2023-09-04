package pubsub

import (
	"time"

	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/hmldd/leo/log"
)

type options struct {
	Logger               log.Logger
	CloseTimeout         time.Duration
	Middlewares          []message.HandlerMiddleware
	Plugins              []message.RouterPlugin
	PublisherDecorators  []message.PublisherDecorator
	SubscriberDecorators []message.SubscriberDecorator
}

func (o *options) init() {
	if o.Logger == nil {
		o.Logger = &log.Discard{}
	}
	if o.CloseTimeout == 0 {
		o.CloseTimeout = time.Second * 30
	}
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

type Option func(o *options)

func Logger(l log.Logger) Option {
	return func(o *options) {
		o.Logger = l
	}
}

func CloseTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.CloseTimeout = timeout
	}
}

func Middleware(mdw ...message.HandlerMiddleware) Option {
	return func(o *options) {
		o.Middlewares = append(o.Middlewares, mdw...)
	}
}

func Plugin(p ...message.RouterPlugin) Option {
	return func(o *options) {
		o.Plugins = append(o.Plugins, p...)
	}
}

func PublisherDecorator(d ...message.PublisherDecorator) Option {
	return func(o *options) {
		o.PublisherDecorators = append(o.PublisherDecorators, d...)
	}
}
func SubscriberDecorator(d ...message.SubscriberDecorator) Option {
	return func(o *options) {
		o.SubscriberDecorators = append(o.SubscriberDecorators, d...)
	}
}
