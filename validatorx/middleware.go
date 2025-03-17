package validatorx

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-leo/leo/v3/statusx"
)

type options struct {
	failFast    bool
	errCallback ErrCallbackFunc
}
type Option func(*options)

func (o *options) apply(opts ...Option) *options {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

type ErrCallbackFunc func(ctx context.Context, err error)

// ErrCallback registers function that will be invoked on validation error(s).
func ErrCallback(callback ErrCallbackFunc) Option {
	return func(o *options) {
		o.errCallback = callback
	}
}

// FailFast tells validator to immediately stop doing further validation after first validation error.
func FailFast() Option {
	return func(o *options) {
		o.failFast = true
	}
}

// Middleware returns a new endpoint.Endpoint that validates request.
// see: https://github.com/envoyproxy/protoc-gen-validate
func Middleware(opts ...Option) endpoint.Middleware {
	o := new(options).apply(opts...)
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			var err error

			if o.failFast {
				switch v := req.(type) {
				case interface{ Validate() error }:
					err = v.Validate()
				case interface{ ValidateAll() error }:
					err = v.ValidateAll()
				}
				return invoke(ctx, req, e, o, err)
			}

			switch v := req.(type) {
			case interface{ ValidateAll() error }:
				err = v.ValidateAll()
			case interface{ Validate() error }:
				err = v.Validate()
			}
			return invoke(ctx, req, e, o, err)
		}
	}
}

func invoke(ctx context.Context, req interface{}, e endpoint.Endpoint, o *options, err error) (interface{}, error) {
	if err != nil {
		if o.errCallback != nil {
			o.errCallback(ctx, err)
		}
		return nil, statusx.InvalidArgument(statusx.Message(err.Error()))
	}
	return e(ctx, req)
}
