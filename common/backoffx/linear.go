package backoffx

import (
	"context"
	"time"
)

// Linear it waits for "delta * attempt" time between calls.
func Linear(delta time.Duration) BackoffFunc {
	return func(ctx context.Context, attempt uint) time.Duration {
		return linear(delta, attempt)
	}
}

func linear(delta time.Duration, attempt uint) time.Duration {
	return delta * time.Duration(attempt)
}
