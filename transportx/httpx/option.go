package httpx

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-leo/leo/v3/transportx/httpx/coder"
	"google.golang.org/protobuf/encoding/protojson"
)

type (
	ServerOptions interface {
		UnmarshalOptions() protojson.UnmarshalOptions
		MarshalOptions() protojson.MarshalOptions
		ResponseTransformer() coder.ResponseTransformer
		Middlewares() []endpoint.Middleware
	}
	serverOptions struct {
		unmarshalOptions    protojson.UnmarshalOptions
		marshalOptions      protojson.MarshalOptions
		responseTransformer coder.ResponseTransformer
		middlewares         []endpoint.Middleware
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

func (o *serverOptions) ResponseTransformer() coder.ResponseTransformer {
	return o.responseTransformer
}

func (o *serverOptions) Middlewares() []endpoint.Middleware {
	return o.middlewares
}

func WithUnmarshalOptions(opts protojson.UnmarshalOptions) ServerOption {
	return func(o *serverOptions) {
		o.unmarshalOptions = opts
	}
}

func WithMarshalOptions(opts protojson.MarshalOptions) ServerOption {
	return func(o *serverOptions) {
		o.marshalOptions = opts
	}
}

func WithResponseTransformer(transformer coder.ResponseTransformer) ServerOption {
	return func(o *serverOptions) {
		o.responseTransformer = transformer
	}
}

// WithMiddleware is a option that sets the go-kit endpoint middlewares.
func WithMiddleware(middlewares ...endpoint.Middleware) ServerOption {
	return func(o *serverOptions) {
		o.middlewares = append(o.middlewares, middlewares...)
	}
}

func NewServerOptions(opts ...ServerOption) ServerOptions {
	o := &serverOptions{
		unmarshalOptions:    protojson.UnmarshalOptions{},
		marshalOptions:      protojson.MarshalOptions{},
		responseTransformer: coder.DefaultResponseTransformer,
		middlewares:         nil,
	}
	o = o.Apply(opts...)
	return o
}
