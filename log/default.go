package log

import (
	"sync"
)

var (
	defaultLogger Logger
	defaultMu     sync.RWMutex
)

func init() {
	defaultLogger = &Discard{}
}

func SetLogger(l Logger) {
	defaultMu.Lock()
	defer defaultMu.Unlock()
	defaultLogger = l
}

func L() Logger {
	defaultMu.RLock()
	defer defaultMu.RUnlock()
	return defaultLogger
}

func Debug(args ...any) {
	L().Debug(args...)
}

func Debugf(template string, args ...any) {
	L().Debugf(template, args...)
}

func DebugF(fields ...Field) {
	L().DebugF(fields...)
}

func Info(args ...any) {
	L().Info(args...)
}

func Infof(template string, args ...any) {
	L().Infof(template, args...)
}

func InfoF(fields ...Field) {
	L().InfoF(fields...)
}

func Warn(args ...any) {
	L().Warn(args...)
}

func Warnf(template string, args ...any) {
	L().Warnf(template, args...)
}

func WarnF(fields ...Field) {
	L().WarnF(fields...)
}

func Error(args ...any) {
	L().Error(args...)
}

func Errorf(template string, args ...any) {
	L().Errorf(template, args...)
}

func ErrorF(fields ...Field) {
	L().ErrorF(fields...)
}

func Panic(args ...any) {
	L().Panic(args...)
}

func Panicf(template string, args ...any) {
	L().Panicf(template, args...)
}

func PanicF(fields ...Field) {
	L().PanicF(fields...)
}

func Fatal(args ...any) {
	L().Fatal(args...)
}

func Fatalf(template string, args ...any) {
	L().Fatalf(template, args...)
}

func FatalF(fields ...Field) {
	L().FatalF(fields...)
}
