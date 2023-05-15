package slog

import (
	"context"

	"codeup.aliyun.com/qimao/leo/leo/log"
)

type logKey struct{}

// NewContext creates a new context with a Logger.
func NewContext(ctx context.Context, l log.Logger, fieldFuncs ...func(ctx context.Context) []log.Field) context.Context {
	creators := make([]log.FieldCreator, 0, len(fieldFuncs))
	for _, f := range fieldFuncs {
		creators = append(creators, log.FieldCreatorFunc(f))
	}
	return context.WithValue(ctx, logKey{}, l.WithContext(ctx, creators...))
}

// 用于给ctx添加指定kv的集合，业务中常在中间件中使用
func NewContextClosure(l log.Logger, fieldFuncs ...func(ctx context.Context) []log.Field) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return NewContext(ctx, l, fieldFuncs...)
	}
}

// FromContext returns a Logger from ctx.
func FromContext(ctx context.Context) (log.Logger, bool) {
	l, ok := ctx.Value(logKey{}).(log.Logger)
	if !ok {
		return nil, false
	}
	return l, true
}

// returns a Logger that discards all log messages.
func FromContextOrDiscard(ctx context.Context) log.Logger {
	l, ok := ctx.Value(logKey{}).(log.Logger)
	if ok {
		return l
	}
	return log.Discard{}
}

// WithCtx alias FromContextOrDiscard
func WithCtx(ctx context.Context) log.Logger {
	return FromContextOrDiscard(ctx)
}
