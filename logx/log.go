package logx

import (
	"context"
	"fmt"
	kitlog "github.com/go-kit/log"
	kitloglevel "github.com/go-kit/log/level"
	stdlog "log"
	"os"
	"slices"
	"sync/atomic"
)

var (
	l atomic.Value
)

func init() {
	l.Store(New(os.Stdout, Logfmt(), Level(kitloglevel.InfoValue()), Timestamp(), Caller(2), Sync()))
}

func Replace(newL kitlog.Logger) kitlog.Logger {
	if newL == nil {
		panic("logx: can not replace default logger with nil")
	}
	oldL := l.Swap(newL)
	return oldL.(kitlog.Logger)
}

func get() kitlog.Logger {
	return l.Load().(kitlog.Logger)
}

var debugLevel = []interface{}{kitloglevel.Key(), kitloglevel.DebugValue()}
var infoLevel = []interface{}{kitloglevel.Key(), kitloglevel.InfoValue()}
var warnLevel = []interface{}{kitloglevel.Key(), kitloglevel.WarnValue()}
var errorLevel = []interface{}{kitloglevel.Key(), kitloglevel.ErrorValue()}

func log(ctx context.Context, level []interface{}, keyvals ...interface{}) {
	if err := get().Log(slices.Concat(FetchKeyValsExtractor(ctx), level, keyvals)...); err != nil {
		stdlog.Println(err.Error())
	}
}

func print(ctx context.Context, level []interface{}, args ...any) {
	if err := get().Log(slices.Concat(FetchKeyValsExtractor(ctx), level, []any{"msg", fmt.Sprint(args...)})...); err != nil {
		stdlog.Println(err.Error())
	}
}

func printf(ctx context.Context, level []interface{}, format string, args ...any) {
	if err := get().Log(slices.Concat(FetchKeyValsExtractor(ctx), level, []any{"msg", fmt.Sprintf(format, args...)})...); err != nil {
		stdlog.Println(err.Error())
	}
}

func Debug(ctx context.Context, keyvals ...interface{}) {
	log(ctx, debugLevel, keyvals...)
}

func Debugf(ctx context.Context, format string, args ...any) {
	printf(ctx, debugLevel, format, args...)
}

func Debugln(ctx context.Context, args ...any) {
	print(ctx, debugLevel, args...)
}

func Info(ctx context.Context, keyvals ...interface{}) {
	log(ctx, infoLevel, keyvals...)
}

func Infof(ctx context.Context, format string, args ...any) {
	printf(ctx, infoLevel, format, args...)
}

func Infoln(ctx context.Context, args ...any) {
	print(ctx, infoLevel, args...)
}

func Warn(ctx context.Context, keyvals ...interface{}) {
	log(ctx, warnLevel, keyvals...)
}

func Warnf(ctx context.Context, format string, args ...any) {
	printf(ctx, warnLevel, format, args...)
}

func Warnln(ctx context.Context, args ...any) {
	print(ctx, warnLevel, args...)
}

func Error(ctx context.Context, keyvals ...interface{}) {
	log(ctx, errorLevel, keyvals...)
}

func Errorf(ctx context.Context, format string, args ...any) {
	printf(ctx, errorLevel, format, args...)
}

func Errorln(ctx context.Context, args ...any) {
	print(ctx, errorLevel, args...)
}
