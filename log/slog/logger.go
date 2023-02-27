package slog

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/exp/slog"

	"codeup.aliyun.com/qimao/leo/leo/log"
)

var _ log.Logger = new(Logger)

type Logger struct {
	o         *options
	level     *slog.LevelVar
	Logger    *slog.Logger
	field     []log.Field
	callDepth int
}

func (l *Logger) SetLevel(lvl log.Level) {
	switch lvl.Code() {
	case log.LevelDebugCode:
		l.level.Set(sLogLevelDebug)
	case log.LevelInfoCode:
		l.level.Set(sLogLevelInfo)
	case log.LevelWarnCode:
		l.level.Set(sLogLevelWarn)
	case log.LevelErrorCode:
		l.level.Set(sLogLevelError)
	case log.LevelPanicCode:
		l.level.Set(sLogLevelPanic)
	case log.LevelFatalCode:
		l.level.Set(sLogLevelFatal)
	}
}

func (l *Logger) GetLevel() log.Level {
	switch l.level.Level() {
	case sLogLevelDebug:
		return log.LevelDebug
	case sLogLevelInfo:
		return log.LevelInfo
	case sLogLevelWarn:
		return log.LevelWarn
	case sLogLevelError:
		return log.LevelError
	case sLogLevelPanic:
		return log.LevelPanic
	case sLogLevelFatal:
		return log.LevelFatal
	}
	return nil
}

func (l *Logger) EnableDebug() {
	l.level.Set(sLogLevelDebug)
}

func (l *Logger) IsDebugEnabled() bool {
	return l.Logger.Enabled(sLogLevelDebug)
}

func (l *Logger) EnableInfo() {
	l.level.Set(sLogLevelInfo)
}

func (l *Logger) IsInfoEnabled() bool {
	return l.Logger.Enabled(sLogLevelInfo)
}

func (l *Logger) EnableWarn() {
	l.level.Set(sLogLevelWarn)
}

func (l *Logger) IsWarnEnabled() bool {
	return l.Logger.Enabled(sLogLevelWarn)
}

func (l *Logger) EnableError() {
	l.level.Set(sLogLevelError)
}

func (l *Logger) IsErrorEnabled() bool {
	return l.Logger.Enabled(sLogLevelError)
}

func (l *Logger) EnablePanic() {
	l.level.Set(sLogLevelPanic)
}

func (l *Logger) IsPanicEnabled() bool {
	return l.Logger.Enabled(sLogLevelPanic)
}

func (l *Logger) EnableFatal() {
	l.level.Set(sLogLevelFatal)
}

func (l *Logger) IsFatalEnabled() bool {
	return l.Logger.Enabled(sLogLevelFatal)
}

func (l *Logger) Debug(a ...any) {
	l.log(sLogLevelDebug, a...)
}

func (l *Logger) Debugf(format string, a ...any) {
	l.logf(sLogLevelDebug, format, a...)
}

func (l *Logger) DebugF(fs ...log.Field) {
	l.logF(sLogLevelDebug, fs...)
}

func (l *Logger) Info(a ...any) {
	l.log(sLogLevelInfo, a...)
}

func (l *Logger) Infof(format string, a ...any) {
	l.logf(sLogLevelInfo, format, a...)
}

func (l *Logger) InfoF(fs ...log.Field) {
	l.logF(sLogLevelInfo, fs...)
}

func (l *Logger) Warn(a ...any) {
	l.log(sLogLevelWarn, a...)
}

func (l *Logger) Warnf(format string, a ...any) {
	l.logf(sLogLevelWarn, format, a...)
}

func (l *Logger) WarnF(fs ...log.Field) {
	l.logF(sLogLevelWarn, fs...)
}

func (l *Logger) Error(a ...any) {
	l.log(sLogLevelError, a...)
}

func (l *Logger) Errorf(format string, a ...any) {
	l.logf(sLogLevelError, format, a...)
}

func (l *Logger) ErrorF(fs ...log.Field) {
	l.logF(sLogLevelError, fs...)
}

func (l *Logger) Panic(a ...any) {
	l.log(sLogLevelError, a...)
	panic(fmt.Sprint(a...))
}

func (l *Logger) Panicf(format string, a ...any) {
	l.logf(sLogLevelError, format, a...)
	panic(fmt.Sprintf(format, a...))
}

func (l *Logger) PanicF(fs ...log.Field) {
	l.logF(sLogLevelError, fs...)
	panic(fields(fs).String())
}

func (l *Logger) Fatal(a ...any) {
	l.log(sLogLevelError, a...)
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, a ...any) {
	l.logf(sLogLevelError, format, a...)
	os.Exit(1)
}

func (l *Logger) FatalF(fs ...log.Field) {
	l.logF(sLogLevelError, fs...)
	os.Exit(1)
}

func (l *Logger) SkipCaller(calldepth int) log.Logger {
	cloned := l.clone().(*Logger)
	cloned.callDepth = calldepth
	return l
}

func (l *Logger) With(fs ...log.Field) log.Logger {
	return l.clone(l.fieldsToAttrs(fs...)...)
}

func (l *Logger) Clone() log.Logger {
	return l.clone()
}

func (l *Logger) ServeHTTP(http.ResponseWriter, *http.Request) {}

func (l *Logger) log(level slog.Level, a ...any) {
	l.Logger.LogAttrsDepth(l.callDepth, level, fmt.Sprint(a...))
}

func (l *Logger) logf(level slog.Level, format string, a ...any) {
	l.Logger.LogAttrsDepth(l.callDepth, level, fmt.Sprintf(format, a...))
}

func (l *Logger) logF(level slog.Level, fs ...log.Field) {
	var msg string
	var args []slog.Attr
	for _, f := range fs {
		if log.IsMsgField(f) {
			msg = f.Value().(string)
			continue
		}
		if log.IsErrField(f) {
			args = append(args, slog.String(slog.ErrorKey, fmt.Sprintf("%v", f.Value())))
			continue
		}
		args = append(args, slog.Any(f.Key(), f.Value()))
	}
	l.Logger.LogAttrsDepth(l.callDepth, level, msg, args...)
}

func (l *Logger) fieldsToAttrs(fs ...log.Field) []any {
	var attrs []any
	for _, f := range fs {
		key := f.Key()
		value := f.Value()
		attr := slog.Any(key, value)
		attrs = append(attrs, attr)
	}
	return attrs
}

func (l *Logger) clone(a ...any) log.Logger {
	cloned := *l
	cloned.Logger = l.Logger.With(a...)
	l.Logger.With()
	return &cloned
}

type fields []log.Field

func (fs fields) String() string {
	var stringBuilder strings.Builder
	for _, f := range fs {
		stringBuilder.WriteString(fmt.Sprintf("%s: %v, ", f.Key(), f.Value()))
	}
	return stringBuilder.String()
}
