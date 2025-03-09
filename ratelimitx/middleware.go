package ratelimitx

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
	kratosrate "github.com/go-kratos/aegis/ratelimit"
	"github.com/go-kratos/aegis/ratelimit/bbr"
	"github.com/go-leo/leo/v3/endpointx"
	"github.com/go-leo/leo/v3/statusx"
)

const (
	errMessageFormat = "ratelimitx: %s is rejected by limiter"
)

func allowerMiddleware(limiter ratelimit.Allower) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
			// extract endpoint name
			endpointName, ok := endpointx.ExtractName(ctx)
			if !ok {
				return next(ctx, request)
			}
			// limit the request
			if !limiter.Allow() {
				return nil, statusx.ResourceExhausted(statusx.Message(errMessageFormat, endpointName))
			}
			// continue handle the request
			return next(ctx, request)
		}
	}
}

func waiterMiddleware(limiter ratelimit.Waiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
			// extract endpoint name
			endpointName, ok := endpointx.ExtractName(ctx)
			if !ok {
				return next(ctx, request)
			}
			// limit the request
			if err := limiter.Wait(ctx); err != nil {
				if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, ratelimit.ErrLimited) {
					return nil, statusx.ResourceExhausted(statusx.Message(errMessageFormat, endpointName))
				}
				return nil, statusx.Unknown(statusx.Message(err.Error()))
			}
			// continue handle the request
			return next(ctx, request)
		}
	}
}
