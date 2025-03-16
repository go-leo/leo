package ratelimitx

import (
	"context"
	_ "embed"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-redis/redis_rate/v10"
)

// Redis creates a rate-limiting middleware using a redis_rate.Limiter.
// see: https://en.wikipedia.org/wiki/Generic_cell_rate_algorithm
// see: https://github.com/go-redis/redis_rate
func Redis(limiter *redis_rate.Limiter, limit redis_rate.Limit, keyFunc func(ctx context.Context) string) endpoint.Middleware {
	return waiterMiddleware(&redisLimiterWrapper{
		limiter: limiter,
		limit:   limit,
		KeyFunc: keyFunc,
	})
}

type redisLimiterWrapper struct {
	// limiter is redis_rate limiter
	limiter *redis_rate.Limiter

	// Limit is limit
	limit redis_rate.Limit

	// KeyFunc is key func
	KeyFunc func(ctx context.Context) string
}

func (limiter *redisLimiterWrapper) Wait(ctx context.Context) error {
	allow, err := limiter.limiter.Allow(ctx, limiter.KeyFunc(ctx), limiter.limit)
	if err != nil {
		return err
	}
	if allow.Allowed <= 0 {
		return ratelimit.ErrLimited
	}
	return nil
}
