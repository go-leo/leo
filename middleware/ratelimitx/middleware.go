package ratelimitx

import (
	"context"
	"errors"
	"github.com/RussellLuo/slidingwindow"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kratos/aegis/ratelimit/bbr"
	"github.com/go-leo/leo/v3/endpointx"
	"github.com/go-leo/leo/v3/statusx"
	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
	uberrate "go.uber.org/ratelimit"
	"golang.org/x/time/rate"
	"time"
)

// SlideWindow accepts a sliding window rate limiter limiter as a parameter and returns a middleware function
// allowerMiddleware(limiter). This middleware is typically used for service endpoints to control the frequency of
// requests using a sliding window algorithm, ensuring the system does not become overloaded due to too many requests.
// see: https://github.com/RussellLuo/slidingwindow
func SlideWindow(limiter *slidingwindow.Limiter) endpoint.Middleware {
	return allowerMiddleware(limiter)
}

// LeakyBucket creates a rate-limiting middleware using an uberRateLimiterWrapper. It initializes the wrapper with a
// limiter and starts the rate-limiting logic, then returns the middleware.
// see: https://github.com/uber-go/ratelimit
func LeakyBucket(limiter uberrate.Limiter, exitC <-chan struct{}) endpoint.Middleware {
	wrapper := uberRateLimiterWrapper{
		limiter: limiter,
		timeC:   make(chan time.Time, 1),
		exitC:   exitC,
	}
	wrapper.start()
	return waiterMiddleware(wrapper)
}

// TokenBucket accepts a parameter limiter of type *rate.Limiter and returns an endpoint.Middleware. It converts the
// rate limiter into a middleware by calling the waiterMiddleware function, which is used to control the request rate.
// see: https://pkg.go.dev/golang.org/x/time/rate
func TokenBucket(limiter *rate.Limiter) endpoint.Middleware {
	return waiterMiddleware(limiter)
}

// Redis creates a rate-limiting middleware using a redisLimiter.
// see: https://redis.io/glossary/rate-limiting/
func Redis(rate int,
	rateInterval time.Duration,
	client redis.Scripter,
	KeyFunc func(ctx context.Context) string,
) endpoint.Middleware {
	return waiterMiddleware(&redisLimiter{
		rate:         rate,
		rateInterval: rateInterval,
		scripter:     client,
		KeyFunc:      KeyFunc,
	})
}

// RedisGCRA creates a rate-limiting middleware using a redis_rate.Limiter.
// see: https://en.wikipedia.org/wiki/Generic_cell_rate_algorithm
// see: https://github.com/go-redis/redis_rate
func RedisGCRA(limiter *redis_rate.Limiter, limit redis_rate.Limit, keyFunc func(ctx context.Context) string) endpoint.Middleware {
	return waiterMiddleware(&redisGCRALimiterWrapper{
		limiter: limiter,
		limit:   limit,
		KeyFunc: keyFunc,
	})
}

// BBR creates a rate-limiting middleware using a bbrLimiter.
// see: https://github.com/alibaba/Sentinel/wiki/系统自适应限流
// see: https://go-kratos.dev/docs/component/middleware/ratelimit/
// see: https://github.com/go-kratos/aegis/tree/main/ratelimit/bbr
func BBR(limiter *bbr.BBR) endpoint.Middleware {
	return allowerMiddleware(&bbrLimiter{limiter: limiter})
}

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
				return nil, statusx.ErrResourceExhausted.With(statusx.Message("ratelimitx: %s is rejected by limiter", endpointName))
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
					return nil, statusx.ErrResourceExhausted.With(statusx.Message("ratelimitx: %s is rejected by limiter", endpointName))
				}
				return nil, statusx.ErrUnknown.With(statusx.Wrap(err))
			}
			// continue handle the request
			return next(ctx, request)
		}
	}
}
