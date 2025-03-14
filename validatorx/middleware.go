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
func Middleware(opts ...Option) endpoint.Middleware {
	o := new(options).apply(opts...)
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			var err error
			switch v := req.(type) {
			case interface{ ValidateAll() error }:
				if o.failFast {
					err = v.ValidateAll()
				}
			case interface{ Validate(all bool) error }:
				err = v.Validate(!o.failFast)
			case interface{ Validate() error }:
				err = v.Validate()
			}
			if err != nil {
				if o.errCallback != nil {
					o.errCallback(ctx, err)
				}
				return nil, statusx.InvalidArgument(statusx.Message(err.Error()))
			}
			return e(ctx, req)
		}
	}
}
