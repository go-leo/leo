package backoffx

import (
	"context"
	"math/rand"
	"time"
)

// JitterUp adds random jitter to the interval.
//
// This adds or subtracts time from the interval within a given jitter fraction.
// For example for 10s and jitter 0.1, it will return a time within [9s, 11s])
// Deprecated: Do not use. use github.com/go-leo/backoffx instead.
func JitterUp(backoff BackoffFunc, jitter float64) BackoffFunc {
	return func(ctx context.Context, attempt uint) time.Duration {
		interval := backoff(ctx, attempt)
		return jitterUp(interval, jitter)
	}
}

func jitterUp(interval time.Duration, jitter float64) time.Duration {
	multiplier := jitter * (rand.Float64()*2 - 1)
	return time.Duration(float64(interval) * (1 + multiplier))
}
