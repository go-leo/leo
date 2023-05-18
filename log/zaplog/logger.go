package zaplog

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"codeup.aliyun.com/qimao/leo/leo/log"
)

var _ log.Logger = new(logger)

type logger struct {
	level zap.AtomicLevel
	zl    *zap.Logger
	zsl   *zap.SugaredLogger
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
	case zapcore.DebugLevel:
		return log.LevelDebug
	case zapcore.InfoLevel:
		return log.LevelInfo
	case zapcore.WarnLevel:
		return log.LevelWarn
	case zapcore.ErrorLevel:
		return log.LevelError
	case zapcore.PanicLevel:
		return log.LevelPanic
	case zapcore.FatalLevel:
		return log.LevelFatal
	}
	return nil
}

func (l *logger) EnableDebug() {
	l.level.SetLevel(zapcore.DebugLevel)
}

func (l *logger) IsDebugEnabled() bool {
	return l.level.Enabled(zapcore.DebugLevel)
}

func (l *logger) EnableInfo() {
	l.level.SetLevel(zapcore.InfoLevel)
}

func (l *logger) IsInfoEnabled() bool {
	return l.level.Enabled(zapcore.InfoLevel)
}

func (l *logger) EnableWarn() {
	l.level.SetLevel(zapcore.WarnLevel)
}

func (l *logger) IsWarnEnabled() bool {
	return l.level.Enabled(zapcore.WarnLevel)
}

func (l *logger) EnableError() {
	l.level.SetLevel(zapcore.ErrorLevel)
}

func (l *logger) IsErrorEnabled() bool {
	return l.level.Enabled(zap.ErrorLevel)
}

func (l *logger) EnablePanic() {
	l.level.Enabled(zapcore.PanicLevel)
}

func (l *logger) IsPanicEnabled() bool {
	return l.level.Enabled(zapcore.PanicLevel)
}

func (l *logger) EnableFatal() {
	l.level.SetLevel(zapcore.FatalLevel)
}

func (l *logger) IsFatalEnabled() bool {
	return l.level.Enabled(zapcore.FatalLevel)
}

func (l *logger) SkipCaller(depth int) log.Logger {
	return l.clone(zap.WithCaller(true), zap.AddCallerSkip(depth))
}

func (l *logger) With(fields ...log.Field) log.Logger {
	return l.clone(zap.Fields(toZapFields(fields...)...))
}

func (l *logger) WithContext(ctx context.Context, creators ...log.FieldCreator) log.Logger {
	var fields []log.Field
	for _, creator := range creators {
		fields = append(fields, creator.Create(ctx)...)
	}
	return l.clone(zap.Fields(toZapFields(fields...)...))
}

func (l *logger) Clone() log.Logger {
	return l.clone()
}

func (l *logger) Debug(args ...any) {
	l.zsl.Debug(args...)
}

func (l *logger) Debugf(template string, args ...any) {
	l.zsl.Debugf(template, args...)
}

func (l *logger) DebugF(fields ...log.Field) {
	fs := toZapFields(fields...)
	l.zl.Debug("", fs...)
}

func (l *logger) Info(args ...any) {
	l.zsl.Info(args...)
}

func (l *logger) Infof(template string, args ...any) {
	l.zsl.Infof(template, args...)
}

func (l *logger) InfoF(fields ...log.Field) {
	fs := toZapFields(fields...)
	l.zl.Info("", fs...)
}

func (l *logger) Warn(args ...any) {
	l.zsl.Warn(args...)
}

func (l *logger) Warnf(template string, args ...any) {
	l.zsl.Warnf(template, args...)
}

func (l *logger) WarnF(fields ...log.Field) {
	fs := toZapFields(fields...)
	l.zl.Warn("", fs...)
}

func (l *logger) Error(args ...any) {
	l.zsl.Error(args...)
}

func (l *logger) Errorf(template string, args ...any) {
	l.zsl.Errorf(template, args...)
}

func (l *logger) ErrorF(fields ...log.Field) {
	fs := toZapFields(fields...)
	l.zl.Error("", fs...)
}

func (l *logger) Panic(args ...any) {
	l.zsl.Panic(args...)
}

func (l *logger) Panicf(template string, args ...any) {
	l.zsl.Panicf(template, args...)
}

func (l *logger) PanicF(fields ...log.Field) {
	fs := toZapFields(fields...)
	l.zl.Panic("", fs...)
}

func (l *logger) Fatal(args ...any) {
	l.zsl.Fatal(args...)
}

func (l *logger) Fatalf(template string, args ...any) {
	l.zsl.Fatalf(template, args...)
}

func (l *logger) FatalF(fields ...log.Field) {
	fs := toZapFields(fields...)
	l.zl.Fatal("", fs...)
}

func (l *logger) clone(opts ...zap.Option) log.Logger {
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
