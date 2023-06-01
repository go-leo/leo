package slog

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"golang.org/x/exp/slog"

	"codeup.aliyun.com/qimao/leo/leo/log"
)

var _ log.Logger = new(logger)

type logger struct {
	o         *options
	level     *slog.LevelVar
	Logger    *slog.Logger
	field     []log.Field
	callDepth int
}

func (l *logger) SetLevel(lvl log.Level) {
	switch lvl.Code() {
	case log.LevelDebugCode:
		l.EnableDebug()
	case log.LevelInfoCode:
		l.EnableInfo()
	case log.LevelWarnCode:
		l.EnableWarn()
	case log.LevelErrorCode:
		l.EnableError()
	case log.LevelPanicCode:
		l.EnablePanic()
	case log.LevelFatalCode:
		l.EnableFatal()
	}
}

func (l *logger) GetLevel() log.Level {
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

func (l *logger) EnableDebug() {
	l.level.Set(sLogLevelDebug)
}

func (l *logger) IsDebugEnabled() bool {
	return l.Logger.Enabled(context.Background(), sLogLevelDebug)
}

func (l *logger) EnableInfo() {
	l.level.Set(sLogLevelInfo)
}

func (l *logger) IsInfoEnabled() bool {
	return l.Logger.Enabled(context.Background(), sLogLevelInfo)
}

func (l *logger) EnableWarn() {
	l.level.Set(sLogLevelWarn)
}

func (l *logger) IsWarnEnabled() bool {
	return l.Logger.Enabled(context.Background(), sLogLevelWarn)
}

func (l *logger) EnableError() {
	l.level.Set(sLogLevelError)
}

func (l *logger) IsErrorEnabled() bool {
	return l.Logger.Enabled(context.Background(), sLogLevelError)
}

func (l *logger) EnablePanic() {
	l.level.Set(sLogLevelPanic)
}

func (l *logger) IsPanicEnabled() bool {
	return l.Logger.Enabled(context.Background(), sLogLevelPanic)
}

func (l *logger) EnableFatal() {
	l.level.Set(sLogLevelFatal)
}

func (l *logger) IsFatalEnabled() bool {
	return l.Logger.Enabled(context.Background(), sLogLevelFatal)
}

func (l *logger) SetDefault() log.Logger {
	log.SetLogger(l.SkipCaller(4))
	return l
}

func (l *logger) SkipCaller(calldepth int) log.Logger {
	cloned := l.clone().(*logger)
	cloned.callDepth = calldepth
	return cloned
}

func (l *logger) With(fields ...log.Field) log.Logger {
	return l.clone(l.fieldsToAttrs(fields...)...)
}

func (l *logger) WithContext(ctx context.Context, creators ...log.FieldCreator) log.Logger {
	var fields []log.Field
	for _, creator := range creators {
		fields = append(fields, creator.Create(ctx)...)
	}
	return l.clone(l.fieldsToAttrs(fields...)...)
}

func (l *logger) Clone() log.Logger {
	return l.clone()
}

func (l *logger) Debug(a ...any) {
	l.log(sLogLevelDebug, a...)
}

func (l *logger) Debugf(format string, a ...any) {
	l.logf(sLogLevelDebug, format, a...)
}

func (l *logger) DebugF(fs ...log.Field) {
	l.logF(sLogLevelDebug, fs...)
}

func (l *logger) Info(a ...any) {
	l.log(sLogLevelInfo, a...)
}

func (l *logger) Infof(format string, a ...any) {
	l.logf(sLogLevelInfo, format, a...)
}

func (l *logger) InfoF(fs ...log.Field) {
	l.logF(sLogLevelInfo, fs...)
}

func (l *logger) Warn(a ...any) {
	l.log(sLogLevelWarn, a...)
}

func (l *logger) Warnf(format string, a ...any) {
	l.logf(sLogLevelWarn, format, a...)
}

func (l *logger) WarnF(fs ...log.Field) {
	l.logF(sLogLevelWarn, fs...)
}

func (l *logger) Error(a ...any) {
	l.log(sLogLevelError, a...)
}

func (l *logger) Errorf(format string, a ...any) {
	l.logf(sLogLevelError, format, a...)
}

func (l *logger) ErrorF(fs ...log.Field) {
	l.logF(sLogLevelError, fs...)
}

func (l *logger) Panic(a ...any) {
	l.log(sLogLevelError, a...)
	panic(fmt.Sprint(a...))
}

func (l *logger) Panicf(format string, a ...any) {
	l.logf(sLogLevelError, format, a...)
	panic(fmt.Sprintf(format, a...))
}

func (l *logger) PanicF(fs ...log.Field) {
	l.logF(sLogLevelError, fs...)
	panic(fields(fs).String())
}

func (l *logger) Fatal(a ...any) {
	l.log(sLogLevelError, a...)
	os.Exit(1)
}

func (l *logger) Fatalf(format string, a ...any) {
	l.logf(sLogLevelError, format, a...)
	os.Exit(1)
}

func (l *logger) FatalF(fs ...log.Field) {
	l.logF(sLogLevelError, fs...)
	os.Exit(1)
}

func (l *logger) log(level slog.Level, a ...any) {
	msg := fmt.Sprint(a...)
	if !l.Logger.Enabled(context.Background(), level) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(l.callDepth, pcs[:])
	r := slog.NewRecord(time.Now(), level, msg, pcs[0])
	_ = l.Logger.Handler().Handle(context.Background(), r)
}

func (l *logger) logf(level slog.Level, format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	if !l.Logger.Enabled(context.Background(), level) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(l.callDepth, pcs[:])
	r := slog.NewRecord(time.Now(), level, msg, pcs[0])
	_ = l.Logger.Handler().Handle(context.Background(), r)
}

func (l *logger) logF(level slog.Level, fs ...log.Field) {
	var msg string
	var attrs []slog.Attr
	for _, f := range fs {
		if log.IsMsgField(f) {
			msg = f.Value().(string)
			continue
		}
		if log.IsErrField(f) {
			attrs = append(attrs, slog.String("error", fmt.Sprintf("%v", f.Value())))
			continue
		}
		attrs = append(attrs, slog.Any(f.Key(), f.Value()))
	}
	var pcs [1]uintptr
	runtime.Callers(l.callDepth, pcs[:])
	r := slog.NewRecord(time.Now(), level, msg, pcs[0])
	r.AddAttrs(attrs...)
	_ = l.Logger.Handler().Handle(context.Background(), r)
}

func (l *logger) fieldsToAttrs(fs ...log.Field) []any {
	var attrs []any
	for _, f := range fs {
		key := f.Key()
		value := f.Value()
		attr := slog.Any(key, value)
		attrs = append(attrs, attr)
	}
	return attrs
}

func (l *logger) clone(a ...any) log.Logger {
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
