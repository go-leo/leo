package zap

import (
	"net/http"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/go-leo/leo/v2/log"
)

var _ log.Logger = new(Logger)

type Logger struct {
	level zap.AtomicLevel
	zl    *zap.Logger
	zsl   *zap.SugaredLogger
}

func (l *Logger) SetLevel(level log.Level) {
	switch level {
	case log.Debug:
		l.EnableDebug()
	case log.Info:
		l.EnableInfo()
	case log.Warn:
		l.EnableWarn()
	case log.Error:
		l.EnableError()
	case log.Panic:
		l.EnablePanic()
	case log.Fatal:
		l.EnableFatal()
	}
}

func (l *Logger) GetLevel() log.Level {
	switch l.level.Level() {
	case zapcore.DebugLevel:
		return log.Debug
	case zapcore.InfoLevel:
		return log.Info
	case zapcore.WarnLevel:
		return log.Warn
	case zapcore.ErrorLevel:
		return log.Error
	case zapcore.PanicLevel:
		return log.Panic
	case zapcore.FatalLevel:
		return log.Fatal
	}
	return ""
}

func (l *Logger) EnableDebug() {
	l.level.SetLevel(zapcore.DebugLevel)
}

func (l *Logger) IsDebugEnabled() bool {
	return l.level.Enabled(zapcore.DebugLevel)
}

func (l *Logger) EnableInfo() {
	l.level.SetLevel(zapcore.InfoLevel)
}

func (l *Logger) IsInfoEnabled() bool {
	return l.level.Enabled(zapcore.InfoLevel)
}

func (l *Logger) EnableWarn() {
	l.level.SetLevel(zapcore.WarnLevel)
}

func (l *Logger) IsWarnEnabled() bool {
	return l.level.Enabled(zapcore.WarnLevel)
}

func (l *Logger) EnableError() {
	l.level.SetLevel(zapcore.ErrorLevel)
}

func (l *Logger) IsErrorEnabled() bool {
	return l.level.Enabled(zap.ErrorLevel)
}

func (l *Logger) EnablePanic() {
	l.level.Enabled(zapcore.PanicLevel)
}

func (l *Logger) IsPanicEnabled() bool {
	return l.level.Enabled(zapcore.PanicLevel)
}

func (l *Logger) EnableFatal() {
	l.level.SetLevel(zapcore.FatalLevel)
}

func (l *Logger) IsFatalEnabled() bool {
	return l.level.Enabled(zapcore.FatalLevel)
}

func (l *Logger) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	l.level.ServeHTTP(resp, req)
}

func (l *Logger) SkipCaller(depth int) log.Logger {
	return l.clone(zap.WithCaller(true), zap.AddCallerSkip(depth))
}

func (l *Logger) With(fields ...log.Field) log.Logger {
	return l.clone(zap.Fields(toZapFields(fields...)...))
}

func (l *Logger) Clone() log.Logger {
	return l.clone()
}

func (l *Logger) Debug(args ...any) {
	l.zsl.Debug(args...)
}

func (l *Logger) Debugf(template string, args ...any) {
	l.zsl.Debugf(template, args...)
}

func (l *Logger) DebugF(fields ...log.Field) {
	fs := toZapFields(fields...)
	l.zl.Debug("", fs...)
}

func (l *Logger) Info(args ...any) {
	l.zsl.Info(args...)
}

func (l *Logger) Infof(template string, args ...any) {
	l.zsl.Infof(template, args...)
}

func (l *Logger) InfoF(fields ...log.Field) {
	fs := toZapFields(fields...)
	l.zl.Info("", fs...)
}

func (l *Logger) Warn(args ...any) {
	l.zsl.Warn(args...)
}

func (l *Logger) Warnf(template string, args ...any) {
	l.zsl.Warnf(template, args...)
}

func (l *Logger) WarnF(fields ...log.Field) {
	fs := toZapFields(fields...)
	l.zl.Warn("", fs...)
}

func (l *Logger) Error(args ...any) {
	l.zsl.Error(args...)
}

func (l *Logger) Errorf(template string, args ...any) {
	l.zsl.Errorf(template, args...)
}

func (l *Logger) ErrorF(fields ...log.Field) {
	fs := toZapFields(fields...)
	l.zl.Error("", fs...)
}

func (l *Logger) Panic(args ...any) {
	l.zsl.Panic(args...)
}

func (l *Logger) Panicf(template string, args ...any) {
	l.zsl.Panicf(template, args...)
}

func (l *Logger) PanicF(fields ...log.Field) {
	fs := toZapFields(fields...)
	l.zl.Panic("", fs...)
}

func (l *Logger) Fatal(args ...any) {
	l.zsl.Fatal(args...)
}

func (l *Logger) Fatalf(template string, args ...any) {
	l.zsl.Fatalf(template, args...)
}

func (l *Logger) FatalF(fields ...log.Field) {
	fs := toZapFields(fields...)
	l.zl.Fatal("", fs...)
}

func (l *Logger) clone(opts ...zap.Option) log.Logger {
	cloned := *l
	cloned.zl = l.zl.WithOptions(opts...)
	cloned.zsl = cloned.zl.Sugar()
	return &cloned
}

func toZapFields(fields ...log.Field) []zap.Field {
	zapFields := make([]zap.Field, 0, len(fields))
	for _, field := range fields {
		zapFields = append(zapFields, zap.Any(field.Key(), field.Value()))
	}
	return zapFields
}
