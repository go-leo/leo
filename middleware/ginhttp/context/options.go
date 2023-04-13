package context

import (
	"context"
)

type options struct {
	contextFunc func(ctx context.Context) context.Context
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

type Option func(o *options)

func defaultOptions() *options {
	return &options{
		contextFunc: func(ctx context.Context) context.Context { return ctx },
	}
}

func ContextFunc(contextFunc func(ctx context.Context) context.Context) Option {
	return func(o *options) {
		o.contextFunc = contextFunc
	}
}
