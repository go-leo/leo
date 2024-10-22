package ratelimitx

import (
	"context"
	_ "embed"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-leo/gox/convx"
	"github.com/go-leo/leo/v3/statusx"
	"github.com/redis/go-redis/v9"
	"time"
)

//go:embed script/acquire.lua
var acquireScript string

// redisLimiter
type redisLimiter struct {
	// rate is rate
	rate int
	// rateInterval is rate time interval
	rateInterval time.Duration
	// scripter is redis scripter
	scripter redis.Scripter
	// KeyFunc is key func
	KeyFunc func(ctx context.Context) string
}

func (limiter *redisLimiter) Wait(ctx context.Context) error {
	result, err := limiter.scripter.Eval(ctx, acquireScript, []string{limiter.KeyFunc(ctx)}, limiter.rate, limiter.rateInterval/time.Millisecond, 1).Result()
	if err != nil {
		return statusx.ErrInternal.With(statusx.Wrap(err))
	}
	allow, ok := result.([]any)
	if !ok {
		return statusx.ErrInternal.With(statusx.Message("invalid result"))
	}
	if convx.ToInt(allow[0]) <= 0 {
		return ratelimit.ErrLimited
	}
	return nil
}
