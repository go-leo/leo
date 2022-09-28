package backoffx

import (
	"context"
	"math"
	"time"
)

// Exponential2 it waits for "delta * 2^attempts" time between calls.
// Deprecated: Do not use. use github.com/go-leo/backoffx instead.
func Exponential2(delta time.Duration) BackoffFunc {
	return func(ctx context.Context, attempt uint) time.Duration {
		return exponential2(delta, attempt)
	}
}

// exponential return "delta * 2^attempts" time.duration
// Deprecated: Do not use. use github.com/go-leo/backoffx instead.
func exponential2(delta time.Duration, attempt uint) time.Duration {
	return delta * time.Duration(math.Exp2(float64(attempt)))
}
