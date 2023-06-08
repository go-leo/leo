//go:build !go1.20

package contextx

import (
	"context"
	"os"
	"os/signal"
)

// Signal creates a new context that cancels on the given signals.
func Signal(signals ...os.Signal) (context.Context, context.CancelFunc) {
	return WithSignal(context.Background(), signals...)
}

// WithSignal like signal.NotifyContext.
func WithSignal(ctx context.Context, signals ...os.Signal) (context.Context, context.CancelFunc) {
	return signal.NotifyContext(ctx, signals...)
}
