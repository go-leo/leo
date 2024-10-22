package ratelimitx

import (
	"context"
	uberrate "go.uber.org/ratelimit"
	"time"
)

// uberRateLimiterWrapper is a wrapper of uberrate.Limiter
type uberRateLimiterWrapper struct {

	// limiter is a uberrate.Limiter
	limiter uberrate.Limiter

	// timeC is a channel used to control the rate limiter's token issuance.
	timeC chan time.Time

	// exitC is a channel used to disable the limiter.
	exitC <-chan struct{}
}

// start a goroutine to control the rate limiter's token issuance
func (w uberRateLimiterWrapper) start() {
	go func() {
		for {
			select {
			case w.timeC <- w.limiter.Take():
				// When a token is obtained from w.limiter.Take(), it is sent to the w.timeC channel.
			case <-w.exitC:
				// When a close signal is received from the w.exitC channel, the loop exits and the goroutine terminates.
				return
			}
		}
	}()
}

// Wait is used to control the rate of requests.
func (w uberRateLimiterWrapper) Wait(ctx context.Context) error {
	select {
	case <-ctx.Done():
		// If the context is canceled or the deadline is exceeded, it returns an error.
		// the request is restricted
		return ctx.Err()
	case <-w.timeC:
		// If a signal is received, it indicates that the request is allowed, and it returns no error.
		// request is allowed
		return nil
	case <-w.exitC:
		// If a signal is received, it indicates that the rate limiter is disabled, and it returns no error.
		// disable rate limiter
		return nil
	}
}
