package log

import "context"

type logKey struct{}

// NewContext creates a new context with a Logger.
func NewContext(ctx context.Context, l Logger) context.Context {
	return context.WithValue(ctx, logKey{}, l)
}

// FromContext returns a Logger from ctx.
func FromContext(ctx context.Context) (Logger, bool) {
	l, ok := ctx.Value(logKey{}).(Logger)
	if !ok {
		return nil, false
	}
	return l, true
}

// FromContextOrDefault returns a Logger from ctx. If no Logger is found, this
// returns a default Logger.
func FromContextOrDefault(ctx context.Context) Logger {
	l, ok := ctx.Value(logKey{}).(Logger)
	if ok {
		return l
	}
	return defaultLogger{}
}
