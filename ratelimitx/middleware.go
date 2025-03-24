package ratelimitx

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-leo/leo/v3/statusx"
)

var ErrRejected = statusx.ResourceExhausted(
	statusx.Message("ratelimitx: rejected by limiter"),
	statusx.Identifier("github.com/go-leo/leo/v3/ratelimitx.ErrRejected"),
)

func allowerMiddleware(limiter ratelimit.Allower) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
			// limit the request
			if !limiter.Allow() {
				return nil, ErrRejected
			}
			// continue handle the request
			return next(ctx, request)
		}
	}
}

func waiterMiddleware(limiter ratelimit.Waiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
			// limit the request
			if err := limiter.Wait(ctx); err != nil {
				if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, ratelimit.ErrLimited) {
					return nil, ErrRejected
				}
				return nil, statusx.Unknown(statusx.Message(err.Error()))
			}
			// continue handle the request
			return next(ctx, request)
		}
	}
}
