package log

import (
	"context"
)

type logKey struct{}

// NewContext creates a new context with a Logger.
func NewContext(ctx context.Context, l Logger, fieldFuncs ...func(ctx context.Context) []Field) context.Context {
	creators := make([]FieldCreator, 0, len(fieldFuncs))
	for _, f := range fieldFuncs {
		creators = append(creators, FieldCreatorFunc(f))
	}
	return context.WithValue(ctx, logKey{}, l.WithContext(ctx, creators...))
}

// 用于给ctx添加指定kv的集合，业务中常在中间件中使用
func NewContextClosure(l Logger, fieldFuncs ...func(ctx context.Context) []Field) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return NewContext(ctx, l, fieldFuncs...)
	}
}

// FromContext returns a Logger from ctx.
func FromContext(ctx context.Context) (Logger, bool) {
	l, ok := ctx.Value(logKey{}).(Logger)
	if !ok {
		return nil, false
	}
	return l, true
}

// FromContextOrDiscard returns a Logger from ctx.  If no Logger is found, this
// returns a Logger that discards all log messages.
func FromContextOrDiscard(ctx context.Context) Logger {
	l, ok := ctx.Value(logKey{}).(Logger)
	if ok {
		return l
	}
	return Discard{}
}
