package log

import (
	"sync"
)

var (
	dl        Logger
	l         Logger
	defaultMu sync.RWMutex
)

func init() {
	SetLogger(&Discard{})
}

func SetLogger(logger Logger) {
	defaultMu.Lock()
	defer defaultMu.Unlock()
	l = logger
	dl = l.SkipCaller(1)
}

func L() Logger {
	var logger Logger
	defaultMu.RLock()
	logger = l
	defaultMu.RUnlock()
	return logger
}

func dL() Logger {
	var logger Logger
	defaultMu.RLock()
	logger = dl
	defaultMu.RUnlock()
	return logger
}

func Debug(args ...any) {
	dL().Debug(args...)
}

func Debugf(template string, args ...any) {
	dL().Debugf(template, args...)
}

func DebugF(fields ...Field) {
	dL().DebugF(fields...)
}

func Info(args ...any) {
	dL().Info(args...)
}

func Infof(template string, args ...any) {
	dL().Infof(template, args...)
}

func InfoF(fields ...Field) {
	dL().InfoF(fields...)
}

func Warn(args ...any) {
	dL().Warn(args...)
}

func Warnf(template string, args ...any) {
	dL().Warnf(template, args...)
}

func WarnF(fields ...Field) {
	dL().WarnF(fields...)
}

func Error(args ...any) {
	dL().Error(args...)
}

func Errorf(template string, args ...any) {
	dL().Errorf(template, args...)
}

func ErrorF(fields ...Field) {
	dL().ErrorF(fields...)
}

func Panic(args ...any) {
	dL().Panic(args...)
}

func Panicf(template string, args ...any) {
	dL().Panicf(template, args...)
}

func PanicF(fields ...Field) {
	dL().PanicF(fields...)
}

func Fatal(args ...any) {
	dL().Fatal(args...)
}

func Fatalf(template string, args ...any) {
	dL().Fatalf(template, args...)
}

func FatalF(fields ...Field) {
	dL().FatalF(fields...)
}
