package ratelimitx

import (
	"context"
	_ "embed"
	"github.com/go-kit/kit/ratelimit"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	//go:embed script/sliding_window_limiter.lua
	slidingWindowLimiterLua string

	//go:embed script/leaky_bucket_limiter.lua
	leakyBucketLimiterLua string

	//go:embed script/token_bucket_limiter.lua
	tokenBucketLimiterLua string
)

const (
	slidingWindowLimiterPrefix = "sliding_window_rate_limiter"

	leakyBucketLimiterPrefix = "leaky_bucket_rate_limiter"

	tokenBucketLimiterPrefix = "token_bucket_rate_limiter"
)

type RedisLimiter struct {
	// redisClient redis客户端
	redisClient *redis.Client
	// script 脚本内容
	script string
	// prefix key前缀
	prefix string
	// tokenKey 令牌key
	tokenKey string
	// timestampKey 时间戳key
	timestampKey string
	// rate 令牌桶算法->每秒生成令牌的数量, 漏桶算法->每秒通过请求数量, 滑动窗口算法->滑动窗口流量阈值，每秒请求数量
	rate float64
	// capacity 令牌桶算法->令牌桶容量, 漏桶算法->桶最大容量, 滑动窗口算法->限流窗口总请求数量
	capacity float64
}

func (limiter *RedisLimiter) setKey(endpointName string) {
	limiter.tokenKey = limiter.prefix + ":" + endpointName + ":token"
	limiter.timestampKey = limiter.prefix + ":" + endpointName + ":timestamp"
}

func (limiter *RedisLimiter) Wait(ctx context.Context) error {
	keys := []string{limiter.tokenKey, limiter.timestampKey}
	slice, err := limiter.redisClient.Eval(ctx, limiter.script, keys, limiter.rate, limiter.capacity, time.Now().Unix(), 1).Int64Slice()
	if err != nil {
		return err
	}
	if slice[0] != 1 {
		return ratelimit.ErrLimited
	}
	return nil
}

func NewRedisSlidingWindowLimiter(redisClient *redis.Client, rate float64, capacity float64) *RedisLimiter {
	return &RedisLimiter{
		redisClient: redisClient,
		script:      slidingWindowLimiterLua,
		prefix:      slidingWindowLimiterPrefix,
		rate:        rate,
		capacity:    capacity,
	}
}

func NewRedisLeakyBucketLimiter(redisClient *redis.Client, rate float64, capacity float64) *RedisLimiter {
	return &RedisLimiter{
		redisClient: redisClient,
		script:      leakyBucketLimiterLua,
		prefix:      leakyBucketLimiterPrefix,
		rate:        rate,
		capacity:    capacity,
	}
}

func NewRedisTokenBucketLimiter(redisClient *redis.Client, rate float64, capacity float64) *RedisLimiter {
	return &RedisLimiter{
		redisClient: redisClient,
		script:      tokenBucketLimiterLua,
		prefix:      tokenBucketLimiterPrefix,
		rate:        rate,
		capacity:    capacity,
	}
}
