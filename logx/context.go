package logx

import (
	"context"
	"github.com/go-kit/log"
)

type loggerKey struct{}

func FromContext(ctx context.Context) log.Logger {
	v, ok := ctx.Value(loggerKey{}).(log.Logger)
	if ok {
		return v
	}
	return L()
}
