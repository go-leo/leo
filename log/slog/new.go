package slog

import (
	"io"
	"os"

	"golang.org/x/exp/slog"
	"gopkg.in/natefinch/lumberjack.v2"

	"codeup.aliyun.com/qimao/leo/leo/log"
)

func New(level *slog.LevelVar, opts ...Option) *Logger {
	o := new(options)
	o.apply(opts...)
	o.init()

	var writers []io.Writer
	// 控制台输出
	if o.Console {
		writers = append(writers, os.Stdout)
	}
	// 配置日志文件
	if o.FileOptions != nil {
		writers = append(writers, fileWriter(o.FileOptions))
	}

	w := io.MultiWriter(writers...)

	handlerOptions := &slog.HandlerOptions{
		AddSource:   true,
		Level:       level,
		ReplaceAttr: replaceAttr,
	}

	handler := o.Encoder(handlerOptions, w)
	return &Logger{
		level:     level,
		Logger:    slog.New(handler),
		field:     nil,
		callDepth: 2,
	}
}

func fileWriter(fileOptions *fileOptions) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   fileOptions.Filename,
		MaxSize:    fileOptions.MaxSize,
		MaxAge:     fileOptions.MaxAge,
		MaxBackups: fileOptions.MaxBackups,
	}
}

func replaceAttr(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.LevelKey {
		level := a.Value.Any().(slog.Level)
		switch level {
		case sLogLevelDebug:
			a.Value = slog.StringValue(log.LevelDebugName)
		case sLogLevelInfo:
			a.Value = slog.StringValue(log.LevelInfoName)
		case sLogLevelWarn:
			a.Value = slog.StringValue(log.LevelWarnName)
		case sLogLevelError:
			a.Value = slog.StringValue(log.LevelErrorName)
		case sLogLevelPanic:
			a.Value = slog.StringValue(log.LevelPanicName)
		case sLogLevelFatal:
			a.Value = slog.StringValue(log.LevelFatalName)
		}
	}
	return a
}
