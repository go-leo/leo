package zap

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func New(level zap.AtomicLevel, opts ...Option) *Logger {
	o := new(options)
	o.apply(opts...)
	o.init()
	var cores []zapcore.Core
	// 控制台输出
	if o.Console {
		cores = append(cores, newConsoleCore(o.Encoder, level)...)
	}
	// 配置日志文件
	if o.FileOptions != nil {
		cores = append(cores, newFileCore(o.Encoder, level, o.FileOptions))
	}
	core := zapcore.NewTee(cores...)
	zl := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.Fields(o.Fields...))
	return &Logger{level: level, zl: zl, zsl: zl.Sugar()}
}

func newConsoleCore(enc zapcore.Encoder, level zap.AtomicLevel) []zapcore.Core {
	stdOutCore := zapcore.NewCore(
		enc,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
		zap.LevelEnablerFunc(func(lvl zapcore.Level) bool { return level.Enabled(lvl) && lvl < zapcore.ErrorLevel }),
	)
	stdErrCore := zapcore.NewCore(
		enc,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stderr)),
		zap.LevelEnablerFunc(func(lvl zapcore.Level) bool { return level.Enabled(lvl) && lvl >= zapcore.ErrorLevel }),
	)
	return []zapcore.Core{stdOutCore, stdErrCore}
}

func newFileCore(enc zapcore.Encoder, level zap.AtomicLevel, o *fileOptions) zapcore.Core {
	lj := &lumberjack.Logger{
		Filename:   o.Filename,
		MaxSize:    o.MaxSize,
		MaxAge:     o.MaxAge,
		MaxBackups: o.MaxBackups,
	}
	sync := zapcore.AddSync(lj)
	fileCore := zapcore.NewCore(enc, sync, level)
	return fileCore
}
