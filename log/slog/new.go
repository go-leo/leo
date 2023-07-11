package slog

import (
	"io"
	"os"
	"strings"

	"golang.org/x/exp/slog"
	"gopkg.in/natefinch/lumberjack.v2"

	"codeup.aliyun.com/qimao/leo/leo/log"
)

func New(level *slog.LevelVar, opts ...Option) log.Logger {
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
	return &logger{
		level:  level,
		Logger: slog.New(handler),
		field:  nil,
		skip:   0,
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
	}
	if a.Key == slog.TimeKey {
	}

	switch a.Key {
	case slog.TimeKey:
	case slog.LevelKey:
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
		default:
			a.Value = slog.StringValue("unknown")
		}
	case slog.SourceKey:
		source := a.Value.Any().(*slog.Source)
		if index := strings.LastIndex(source.File, "/"); index > 0 {
			index = strings.LastIndex(source.File[:index], "/")
			if index > 0 {
				source.File = source.File[index+1:]
			}
		}
		if index := strings.LastIndex(source.Function, "/"); index > 0 {
			index = strings.LastIndex(source.Function[:index], "/")
			if index > 0 {
				source.Function = source.Function[index+1:]
			}
		}

	case slog.MessageKey:

	}
	return a
}
