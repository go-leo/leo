package retryx

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-leo/gox/backoff"
	"github.com/go-leo/gox/retry"
	"github.com/go-leo/leo/v3/statusx"
	"time"
)

type options struct {
	maxAttempts uint
	backoffFunc backoff.BackoffFunc
	retryOnFunc func(err error) bool
}

type Option func(*options)

// MaxAttempts set max attempts
func MaxAttempts(maxAttempts uint) Option {
	return func(o *options) {
		o.maxAttempts = maxAttempts
	}
}

// BackoffFunc set backoff func
func BackoffFunc(backoffFunc backoff.BackoffFunc) Option {
	return func(o *options) {
		o.backoffFunc = backoffFunc
	}
}

// RetryOnFunc set retryOn func
func RetryOnFunc(retryOnFunc func(err error) bool) Option {
	return func(o *options) {
		o.retryOnFunc = retryOnFunc
	}
}

func Middleware(opts ...Option) endpoint.Middleware {
	o := &options{
		maxAttempts: 3,
		backoffFunc: backoff.Constant(500 * time.Millisecond),
		retryOnFunc: func(err error) bool {
			 := statusx.From(err)
			return .Code() == statusx.CodeUnknown
		},
	}
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
			var response any
			err := retry.MaxAttempts(o.maxAttempts).
				Backoff(o.backoffFunc).
				RetryOn(o.retryOnFunc).
				Exec(ctx, func(ctx context.Context, attempt uint) error {
					var err error
					response, err = next(ctx, request)
					return err
				})
			if err != nil {
				return nil, err
			}
			return response, err
		}
	}
}
