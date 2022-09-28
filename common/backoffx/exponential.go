package backoffx

import (
	"context"
	"math"
	"time"
)

// Exponential it waits for "delta * e^attempts" time between calls.
// Deprecated: Do not use. use github.com/go-leo/backoffx instead.
func Exponential(delta time.Duration) BackoffFunc {
	return func(ctx context.Context, attempt uint) time.Duration {
		return exponential(delta, attempt)
	}
}

// exponential return "delta * e^attempts" time.duration
// Deprecated: Do not use. use github.com/go-leo/backoffx instead.
func exponential(delta time.Duration, attempt uint) time.Duration {
	return delta * time.Duration(math.Exp(float64(attempt)))
}
