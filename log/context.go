package log

import (
	"context"
	"github.com/gin-gonic/gin"
)

type logKey struct{}

// NewContext creates a new context with a Logger.
func NewContext(ctx context.Context, l Logger, creators ...FieldCreator) context.Context {
	return context.WithValue(ctx, logKey{}, l.WithContext(ctx, creators...))
}

func NewContextClosure(l Logger, creators ...FieldCreator) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return NewContext(ctx, l, creators...)
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

// WithCtx returns a Logger from ctx. 主要用于gin中 ctx 存在内置情况，需要从Request.Context()中获取，所以需要这个方法
func WithCtx(ctx context.Context) Logger {
	if ginCtx, ok := ctx.(*gin.Context); ok {
		ctx = ginCtx.Request.Context()
	}
	l, ok := ctx.Value(logKey{}).(Logger)
	if ok {
		return l
	}
	return Discard{}

}
