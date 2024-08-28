package ratelimitx

import (
	"context"
	"github.com/go-kit/kit/ratelimit"
	"github.com/redis/go-redis/v9"
	"testing"
)

func TestNewRedisSlidingWindowLimiter(t *testing.T) {
	t.Parallel()
	for i := 0; i < 20; i++ {
		limiter := NewRedisSlidingWindowLimiter(redis.NewClient(&redis.Options{Addr: "localhost:6379"}), 1000, 1000)
		limiter.setKey("hello")
		go func(limiter ratelimit.Waiter) {
			for i := 0; i < 1000; i++ {
				if err := limiter.Wait(context.Background()); err != nil {
					panic(err)
				}
			}
		}(limiter)
	}
	select {}
}

func TestNewRedisLeakyBucketLimiter(t *testing.T) {
	t.Parallel()
	for i := 0; i < 20; i++ {
		limiter := NewRedisLeakyBucketLimiter(redis.NewClient(&redis.Options{Addr: "localhost:6379"}), 1000, 1000)
		limiter.setKey("hello")
		go func(limiter ratelimit.Waiter) {
			for i := 0; i < 1000; i++ {
				if err := limiter.Wait(context.Background()); err != nil {
					panic(err)
				}
			}
		}(limiter)
	}
	select {}
}

func TestNewRedisTokenBucketLimiter(t *testing.T) {
	t.Parallel()
	for i := 0; i < 20; i++ {
		limiter := NewRedisTokenBucketLimiter(redis.NewClient(&redis.Options{Addr: "localhost:6379"}), 1000, 1000)
		limiter.setKey("hello")
		go func(limiter ratelimit.Waiter) {
			for j := 0; j < 1000; j++ {
				t.Log(i, j)
				if err := limiter.Wait(context.Background()); err != nil {
					panic(err)
				}
			}
		}(limiter)
	}
	select {}
}
