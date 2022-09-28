package backoffx

import (
	"context"
	"time"
)

// Constant it waits for a fixed period of time between calls.
// Deprecated: Do not use. use github.com/go-leo/backoffx instead.
func Constant(delta time.Duration) BackoffFunc {
	return func(ctx context.Context, attempt uint) time.Duration {
		return delta
	}
}
