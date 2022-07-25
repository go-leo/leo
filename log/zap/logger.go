package zap

import (
	"net/http"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/go-leo/leo/log"
)

var _ log.Logger = new(Logger)

type Logger struct {
	level zap.AtomicLevel
	zl    *zap.Logger
	zsl   *zap.SugaredLogger
}

func LevelAdapt(level log.Level) zap.AtomicLevel {
	switch level {
	case log.Debug:
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	case log.Info:
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	case log.Warn:
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	case log.Error:
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	case log.Panic:
		return zap.NewAtomicLevelAt(zap.PanicLevel)
	case log.Fatal:
		return zap.NewAtomicLevelAt(zap.FatalLevel)
	default:
		return zap.NewAtomicLevel()
	}
}

func New(level zap.AtomicLevel, opts ...Option) *Logger {
	o := new(options)
	o.apply(opts...)
	o.init()
	var cores []zapcore.Core
	// 控制台输出
	if o.Console {
		cores = append(cores, newConsoleCore(o.Encoder, level.Level())...)
	}
	// 配置日志文件
	if o.FileOptions != nil {
		cores = append(cores, newFileCore(o.Encoder, level.Level(), o.FileOptions))
	}
	core := zapcore.NewTee(cores...)
	zl := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.Fields(o.Fields...))
	return &Logger{level: level, zl: zl, zsl: zl.Sugar()}
}

func (l *Logger) SkipCaller(depth int) log.Logger {
	zl := l.zl.WithOptions(zap.WithCaller(true), zap.AddCallerSkip(depth))
	zsl := zl.Sugar()
	copy := *l
	copy.zl = zl
	copy.zsl = zsl
	return &copy
}

func (l *Logger) With(fields ...log.F) log.Logger {
	zapFields := make([]zap.Field, 0)
	for _, field := range fields {
		zapFields = append(zapFields, zap.Any(field.K, field.V))
	}
	copy := *l
	copy.zl = l.zl.With(zapFields...)
	copy.zsl = copy.zl.Sugar()
	return &copy
}

func (l *Logger) Clone() log.Logger {
	copy := *l
	zl := *l.zl
	copy.zl = &zl
	copy.zsl = copy.zl.Sugar()
	copy.zl.Sugar()
	return &copy
}

func (l *Logger) EnableDebug() {
	l.level.SetLevel(zapcore.DebugLevel)
}

func (l *Logger) IsDebugEnabled() bool {
	return l.level.Enabled(zapcore.DebugLevel)
}

func (l *Logger) Debug(args ...any) {
	l.zsl.Debug(args...)
}

func (l *Logger) Debugf(template string, args ...any) {
	l.zsl.Debugf(template, args...)
}

func (l *Logger) DebugF(fields ...log.F) {
	fs := toZapFields(fields...)
	l.zl.Debug("", fs...)
}

func (l *Logger) EnableInfo() {
	l.level.SetLevel(zapcore.InfoLevel)
}

func (l *Logger) IsInfoEnabled() bool {
	return l.level.Enabled(zapcore.InfoLevel)
}

func (l *Logger) Info(args ...any) {
	l.zsl.Info(args...)
}

func (l *Logger) Infof(template string, args ...any) {
	l.zsl.Infof(template, args...)
}

func (l *Logger) InfoF(fields ...log.F) {
	fs := toZapFields(fields...)
	l.zl.Info("", fs...)
}

func (l *Logger) EnableWarn() {
	l.level.SetLevel(zapcore.WarnLevel)
}

func (l *Logger) IsWarnEnabled() bool {
	return l.level.Enabled(zapcore.WarnLevel)
}

func (l *Logger) Warn(args ...any) {
	l.zsl.Warn(args...)
}

func (l *Logger) Warnf(template string, args ...any) {
	l.zsl.Warnf(template, args...)
}

func (l *Logger) WarnF(fields ...log.F) {
	fs := toZapFields(fields...)
	l.zl.Warn("", fs...)
}

func (l *Logger) EnableError() {
	l.level.SetLevel(zapcore.ErrorLevel)
}

func (l *Logger) IsErrorEnabled() bool {
	return l.level.Enabled(zap.ErrorLevel)
}

func (l *Logger) Error(args ...any) {
	l.zsl.Error(args...)
}

func (l *Logger) Errorf(template string, args ...any) {
	l.zsl.Errorf(template, args...)
}

func (l *Logger) ErrorF(fields ...log.F) {
	fs := toZapFields(fields...)
	l.zl.Error("", fs...)
}

func (l *Logger) EnablePanic() {
	l.level.Enabled(zapcore.PanicLevel)
}

func (l *Logger) IsPanicEnabled() bool {
	return l.level.Enabled(zapcore.PanicLevel)
}

func (l *Logger) Panic(args ...any) {
	l.zsl.Panic(args...)
}

func (l *Logger) Panicf(template string, args ...any) {
	l.zsl.Panicf(template, args...)
}

func (l *Logger) PanicF(fields ...log.F) {
	fs := toZapFields(fields...)
	l.zl.Panic("", fs...)
}

func (l *Logger) EnableFatal() {
	l.level.SetLevel(zapcore.FatalLevel)
}

func (l *Logger) IsFatalEnabled() bool {
	return l.level.Enabled(zapcore.FatalLevel)
}

func (l *Logger) Fatal(args ...any) {
	l.zsl.Fatal(args...)
}

func (l *Logger) Fatalf(template string, args ...any) {
	l.zsl.Fatalf(template, args...)
}

func (l *Logger) FatalF(fields ...log.F) {
	fs := toZapFields(fields...)
	l.zl.Fatal("", fs...)
}

func (l *Logger) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	l.level.ServeHTTP(resp, req)
}

func toZapFields(fields ...log.F) []zap.Field {
	zapFields := make([]zap.Field, 0, len(fields))
	for _, field := range fields {
		zapFields = append(zapFields, zap.Any(field.K, field.V))
	}
	return zapFields
}

func newConsoleCore(enc zapcore.Encoder, lvl zapcore.Level) []zapcore.Core {
	stdOutCore := zapcore.NewCore(enc, zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= lvl && level < zapcore.ErrorLevel
	}))
	stdErrCore := zapcore.NewCore(enc, zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stderr)), zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.ErrorLevel
	}))
	return []zapcore.Core{stdOutCore, stdErrCore}
}

func newFileCore(enc zapcore.Encoder, lvl zapcore.Level, o *fileOptions) zapcore.Core {
	lj := &lumberjack.Logger{
		Filename:   o.Filename,
		MaxSize:    o.MaxSize,
		MaxAge:     o.MaxAge,
		MaxBackups: o.MaxBackups,
	}
	sync := zapcore.AddSync(lj)
	fileCore := zapcore.NewCore(enc, sync, lvl)
	return fileCore
}
