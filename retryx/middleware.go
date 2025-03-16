package retryx

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-leo/gox/backoff"
	"github.com/go-leo/leo/v3/statusx"
	"google.golang.org/protobuf/types/known/durationpb"
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
// 只有当服务端返回含有 errdetails.RetryInfo的错误时才会重试.
// 如果RetryDelay为负数，则全链路禁止重试。
func Middleware(opts ...Option) endpoint.Middleware {
	o := &options{
		maxAttempts: 3,
		backoff:     backoff.LinearFactory(),
	}
	o = o.apply(opts...)
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
			resp, err := next(ctx, request)
			if o.maxAttempts <= 0 {
				// 最大重试次数为0，不重试
				return resp, err
			}
			if err == nil {
				// 没有错误，不重试
				return resp, nil
			}

			for attempt := uint(0); attempt < o.maxAttempts; attempt++ {
				// 还原出 statusx.Status
				st, ok := statusx.From(err)
				if !ok {
					// 不是 statusx.Status 错误，不重试
					return nil, err
				}

				// 取出重试信息, 如果没有重试信息，不重试
				info := st.RetryInfo()
				if info == nil {
					return nil, st
				}

				// 取出重试延迟时间, 如果重试延迟时间为负数，不重试
				delay := info.GetRetryDelay().AsDuration()
				if delay < 0 {
					return nil, st
				}

				select {
				case <-ctx.Done():
					// context 被取消，不重试
					return nil, st
				case <-time.After(o.backoff(delay)(ctx, attempt)):
					// 等待重试
					resp, err = next(ctx, request)
					if err == nil {
						// 没有错误，不重试
						return resp, nil
					}
				}
			}

			// 还原出 statusx.Status, 设置RetryDelay为负数，禁止全链路重试
			st, ok := statusx.From(err)
			if !ok {
				return nil, err
			}
			info := st.RetryInfo()
			if info == nil {
				return nil, st
			}
			st.RetryInfo().RetryDelay = durationpb.New(-1)
			return nil, st
		}
	}
}
