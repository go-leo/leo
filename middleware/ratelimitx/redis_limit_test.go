package ratelimitx

import (
	"context"
	_ "embed"
	"github.com/go-kit/kit/ratelimit"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestAcquireLua(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("ADDR"),
	})
	limiter := redisLimiter{
		rate:         5,
		rateInterval: 5 * time.Second,
		scripter:     client,
		KeyFunc: func(ctx context.Context) string {
			return "rate_limit:demo"
		},
	}
	ctx := context.Background()

	err := limiter.Wait(ctx)
	assert.NoError(t, err)

	err = limiter.Wait(ctx)
	assert.NoError(t, err)

	err = limiter.Wait(ctx)
	assert.NoError(t, err)

	err = limiter.Wait(ctx)
	assert.NoError(t, err)

	err = limiter.Wait(ctx)
	assert.NoError(t, err)

	err = limiter.Wait(ctx)
	assert.ErrorIs(t, err, ratelimit.ErrLimited)

}
