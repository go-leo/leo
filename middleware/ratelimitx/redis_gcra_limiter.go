package ratelimitx

import (
	"context"
	_ "embed"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-leo/leo/v3/statusx"
	"github.com/go-redis/redis_rate/v10"
)

type redisGCRALimiterWrapper struct {
	// limiter is redis_rate limiter
	limiter *redis_rate.Limiter

	// Limit is limit
	limit redis_rate.Limit

	// KeyFunc is key func
	KeyFunc func(ctx context.Context) string
}

func (limiter *redisGCRALimiterWrapper) Wait(ctx context.Context) error {
	allow, err := limiter.limiter.Allow(ctx, limiter.KeyFunc(ctx), limiter.limit)
	if err != nil {
		return statusx.ErrInternal.With(statusx.Wrap(err))
	}
	if allow.Allowed <= 0 {
		return ratelimit.ErrLimited
	}
	return nil
}
