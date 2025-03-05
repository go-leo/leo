package httptransportx

import (
	"github.com/go-kit/kit/endpoint"
	"google.golang.org/protobuf/encoding/protojson"
)

type (
	ServerOptions interface {
		// UnmarshalOptions returns the protojson.UnmarshalOptions.
		UnmarshalOptions() protojson.UnmarshalOptions
		// MarshalOptions returns the protojson.MarshalOptions.
		MarshalOptions() protojson.MarshalOptions
		// Middlewares returns the go-kit endpoint middlewares.
		Middlewares() []endpoint.Middleware
	}
	serverOptions struct {
		unmarshalOptions protojson.UnmarshalOptions
		marshalOptions   protojson.MarshalOptions
		middlewares      []endpoint.Middleware
	}
	ServerOption func(o *serverOptions)
)

func (o *serverOptions) Apply(opts ...ServerOption) *serverOptions {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func (o *serverOptions) UnmarshalOptions() protojson.UnmarshalOptions {
	return o.unmarshalOptions
}

func (o *serverOptions) MarshalOptions() protojson.MarshalOptions {
	return o.marshalOptions
}

func (o *serverOptions) Middlewares() []endpoint.Middleware {
	return o.middlewares
}

// UnmarshalOptions is a option that sets the protojson.UnmarshalOptions.
func UnmarshalOptions(opts protojson.UnmarshalOptions) ServerOption {
	return func(o *serverOptions) {
		o.unmarshalOptions = opts
	}
}

// MarshalOptions is a option that sets the protojson.MarshalOptions.
func MarshalOptions(opts protojson.MarshalOptions) ServerOption {
	return func(o *serverOptions) {
		o.marshalOptions = opts
	}
}

// Middleware is a option that sets the go-kit endpoint middlewares.
func Middleware(middlewares ...endpoint.Middleware) ServerOption {
	return func(o *serverOptions) {
		o.middlewares = append(o.middlewares, middlewares...)
	}
}

func NewServerOptions(opts ...ServerOption) ServerOptions {
	o := &serverOptions{
		unmarshalOptions: protojson.UnmarshalOptions{},
		marshalOptions:   protojson.MarshalOptions{},
		middlewares:      nil,
	}
	o = o.Apply(opts...)
	return o
}
