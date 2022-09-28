package backoffx

import (
	"context"
	"time"
)

// BackoffFunc denotes a family of functions that control the backoff duration between call retries.
//
// They are called with an identifier of the attempt, and should return a time the system client should
// hold off for. If the time returned is longer than the `context.Context.Deadline` of the request
// the deadline of the request takes precedence and the wait will be interrupted before proceeding
// with the next iteration.
// The context can be used to extract context values.
// Deprecated: Do not use. use github.com/go-leo/backoffx instead.
type BackoffFunc func(ctx context.Context, attempt uint) time.Duration
