package signalx

import (
	"context"
	"os"
	"os/signal"
)

// SignalContext creates a new context that cancels on the given signals.
// Deprecated: Do not use. use github.com/go-leo/osx/signalx instead.
func SignalContext(signals ...os.Signal) (context.Context, context.CancelFunc) {
	return ContextWithSignal(context.Background(), signals...)
}

// ContextWithSignal creates a new context that cancels on the given signals.
// Deprecated: Do not use. use github.com/go-leo/osx/signalx instead.
func ContextWithSignal(ctx context.Context, signals ...os.Signal) (context.Context, context.CancelFunc) {
	ctx, closer := context.WithCancel(ctx)

	c := make(chan os.Signal, 1)
	signal.Notify(c, signals...)

	go func() {
		select {
		case <-c:
			closer()
		case <-ctx.Done():
			//signal.Stop(c)
		}
	}()

	return ctx, closer
}
