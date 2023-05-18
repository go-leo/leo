package zaplog

import (
	"context"

	"codeup.aliyun.com/qimao/leo/leo/log"
)

// NewContext creates a new context with a Logger.
var NewContext = log.NewContext

// 用于给ctx添加指定kv的集合，业务中常在中间件中使用
var NewContextClosure = log.NewContextClosure

// FromContext returns a Logger from ctx.
var FromContext = log.FromContext

// returns a Logger that discards all log messages.
var FromContextOrDiscard = log.FromContextOrDiscard

// WithCtx alias FromContextOrDiscard
func WithCtx(ctx context.Context) log.Logger {
	l, ok := FromContext(ctx)
	if ok {
		return l
	}
	return Logger()
}
