package backoffx

import (
	"context"
	"time"
)

// Fibonacci it waits for "delta * fibonacci(attempt)" time between calls.
// Deprecated: Do not use. use github.com/go-leo/backoffx instead.
func Fibonacci(delta time.Duration) BackoffFunc {
	return func(ctx context.Context, attempt uint) time.Duration {
		return fibonacci(delta, attempt)
	}
}

func fibonacci(delta time.Duration, attempt uint) time.Duration {
	var (
		pre int64
		cur int64
		i   uint
	)
	for pre, cur, i = 0, 1, 0; i < attempt; i++ {
		pre, cur = cur, pre+cur
	}
	return delta * time.Duration(pre)
}
