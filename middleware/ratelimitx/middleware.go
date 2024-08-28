package ratelimitx

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-leo/leo/v3/endpointx"
	"github.com/go-leo/leo/v3/statusx"
	"golang.org/x/time/rate"
)

func SlideWindow(limiters map[string]*SlideWindowLimiter) endpoint.Middleware {
	allowableLimiters := make(map[string]ratelimit.Allower)
	for key, limiter := range limiters {
		allowableLimiters[key] = limiter
	}
	return allowerMiddleware(allowableLimiters)
}

func LeakyBucket(limiters map[string]*rate.Limiter) endpoint.Middleware {
	waiterLimiters := make(map[string]ratelimit.Waiter)
	for key, limiter := range limiters {
		waiterLimiters[key] = limiter
	}
	return waiterMiddleware(waiterLimiters)
}

func TokenBucket(limiters map[string]*rate.Limiter) endpoint.Middleware {
	allowerLimiters := make(map[string]ratelimit.Allower)
	for key, limiter := range limiters {
		allowerLimiters[key] = limiter
	}
	return allowerMiddleware(allowerLimiters)
}

func Redis(limiters map[string]*RedisLimiter) endpoint.Middleware {
	waiterLimiters := make(map[string]ratelimit.Waiter)
	for key, limiter := range limiters {
		waiterLimiters[key] = limiter
	}
	return waiterMiddleware(waiterLimiters)
}

func allowerMiddleware(limiters map[string]ratelimit.Allower) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
			// extract endpoint name
			endpointName, ok := endpointx.ExtractName(ctx)
			if !ok {
				return next(ctx, request)
			}
			// get limiter
			limiter, ok := limiters[endpointName]
			if !ok {
				limiter, ok = limiters[All]
				if !ok {
					return next(ctx, request)
				}
			}
			// limit the request
			if !limiter.Allow() {
				return nil, statusx.ErrResourceExhausted.With(statusx.Message("ratelimitx: %s is rejected by limiter", endpointName))
			}
			// continue handle the request
			return next(ctx, request)
		}
	}
}

func waiterMiddleware(limiters map[string]ratelimit.Waiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
			// extract endpoint name
			endpointName, ok := endpointx.ExtractName(ctx)
			if !ok {
				return next(ctx, request)
			}
			// get limiter
			limiter, ok := limiters[endpointName]
			if !ok {
				return next(ctx, request)
			}
			// limit the request
			if err := limiter.Wait(ctx); err != nil {
				if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, ratelimit.ErrLimited) {
					return nil, statusx.ErrResourceExhausted.With(statusx.Message("ratelimitx: %s is rejected by limiter", endpointName))
				}
				return nil, statusx.ErrUnknown.With(statusx.Wrap(err))
			}
			// continue handle the request
			return next(ctx, request)
		}
	}
}
