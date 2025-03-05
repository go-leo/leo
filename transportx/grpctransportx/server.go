package grpctransportx

import (
	"github.com/go-kit/kit/endpoint"
)

type (
	ServerOptions interface {
		// Middlewares returns the go-kit endpoint middlewares.
		Middlewares() []endpoint.Middleware
	}
	serverOptions struct {
		middlewares []endpoint.Middleware
	}
	ServerOption func(o *serverOptions)
)

func (o *serverOptions) Apply(opts ...ServerOption) *serverOptions {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func (o *serverOptions) Middlewares() []endpoint.Middleware {
	return o.middlewares
}

// Middleware is a option that sets the go-kit endpoint middlewares.
func Middleware(middlewares ...endpoint.Middleware) ServerOption {
	return func(o *serverOptions) {
		o.middlewares = append(o.middlewares, middlewares...)
	}
}

func NewServerOptions(opts ...ServerOption) ServerOptions {
	o := &serverOptions{
		middlewares: nil,
	}
	o = o.Apply(opts...)
	return o
}
