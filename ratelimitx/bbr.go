package ratelimitx

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	kratosrate "github.com/go-kratos/aegis/ratelimit"
	"github.com/go-kratos/aegis/ratelimit/bbr"
)

// BBR creates a rate-limiting middleware using a bbrLimiter.
// see: https://github.com/alibaba/Sentinel/wiki/系统自适应限流
// see: https://go-kratos.dev/docs/component/middleware/ratelimit/
// see: https://github.com/go-kratos/aegis/tree/main/ratelimit/bbr
func BBR(limiter *bbr.BBR) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
			// limit the request
			doneFunc, err := limiter.Allow()
			if err != nil {
				return nil, ErrRejected
			}
			// continue handle the request
			response, err := next(ctx, request)
			doneFunc(kratosrate.DoneInfo{Err: err})
			return response, err
		}
	}
}
