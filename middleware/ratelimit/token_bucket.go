package ratelimit

import (
	"context"
	"errors"
)

// ErrLimited is returned in the request path when the rate limiter is
// triggered and the request is rejected.
var ErrLimited = errors.New("rate limit exceeded")

// Allower dictates whether or not a request is acceptable to run.
// The Limiter from "golang.org/x/time/rate" already implements this interface,
// one is able to use that in NewErroringLimiter without any modifications.
type Allower interface {
	Allow() bool
}

// Waiter dictates how long a request must be delayed.
// The Limiter from "golang.org/x/time/rate" already implements this interface,
// one is able to use that in NewDelayingLimiter without any modifications.
type Waiter interface {
	Wait(ctx context.Context) error
}
