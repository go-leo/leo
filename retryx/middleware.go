package retryx

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-leo/gox/backoff"
	"github.com/go-leo/leo/v3/statusx"
	"time"
)

type options struct {
	maxAttempts uint
	backoff     BackoffFunc
}

type Option func(*options)

func (o *options) apply(opts ...Option) *options {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// MaxAttempts set max attempts.
func MaxAttempts(maxAttempts uint) Option {
	return func(o *options) {
		o.maxAttempts = maxAttempts
	}
}

// Backoff set backoff.
func Backoff(backoff BackoffFunc) Option {
	return func(o *options) {
		o.backoff = backoff
	}
}

// Middleware returns a retry middleware.
// 只有当服务端返回错误时，并且含有 errdetails.RetryInfo 才会重试.
func Middleware(opts ...Option) endpoint.Middleware {
	o := &options{
		maxAttempts: 3,
		backoff:     backoff.LinearFactory(),
	}
	o = o.apply(opts...)
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
			var attempt uint
			for attempt < o.maxAttempts {
				// execute cmd
				resp, err := next(ctx, request)
				if err == nil {
					// return if err is nil.
					return resp, nil
				}
				st, ok := statusx.From(err)
				if !ok {
					return nil, err
				}
				info := st.RetryInfo()
				if info == nil {
					return nil, err
				}
				// increase the number of attempts
				attempt++
				select {
				case <-ctx.Done():
					// return if context is done, return error, remove retry info.
					return nil, st.WithoutDetail(info)
				case <-time.After(o.backoff(info.GetRetryDelay().AsDuration())(ctx, attempt)):
					// sleep and wait retry
					continue
				}
			}
			// perform the execution
			return next(ctx, request)
		}
	}
}
