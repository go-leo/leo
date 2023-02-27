package slog

import (
	"io"

	"golang.org/x/exp/slog"

	"codeup.aliyun.com/qimao/leo/leo/log"
)

type options struct {
	Console     bool // Console 是否在终端输出
	FileOptions *fileOptions
	Fields      []slog.Attr
	Encoder     func(handlerOptions *slog.HandlerOptions, w io.Writer) slog.Handler
}

type fileOptions struct {
	Filename   string // Filename 文件名
	MaxSize    int    // megabytes
	MaxAge     int    // days
	MaxBackups int    // count
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *options) init() {
	if o.Encoder == nil {
		o.Encoder = func(handlerOptions *slog.HandlerOptions, w io.Writer) slog.Handler {
			return handlerOptions.NewJSONHandler(w)
		}
	}
	if o.FileOptions == nil && !o.Console {
		o.Console = true
	}
}

type Option func(*options)

func File(filename string, maxSize, maxAge, maxBackups int) Option {
	return func(o *options) {
		fileOptions := &fileOptions{
			Filename:   filename,
			MaxSize:    maxSize,
			MaxAge:     maxAge,
			MaxBackups: maxBackups,
		}
		o.FileOptions = fileOptions
	}
}

func Console(enabled bool) Option {
	return func(o *options) {
		o.Console = enabled
	}
}

func Fields(fields ...log.Field) Option {
	return func(l *options) {
		for _, field := range fields {
			l.Fields = append(l.Fields, slog.Any(field.Key(), field.Value()))
		}
	}
}

func JSON() Option {
	return func(o *options) {
		o.Encoder = func(handlerOptions *slog.HandlerOptions, w io.Writer) slog.Handler {
			return handlerOptions.NewJSONHandler(w)
		}
	}
}

func PlainText() Option {
	return func(o *options) {
		o.Encoder = func(handlerOptions *slog.HandlerOptions, w io.Writer) slog.Handler {
			return handlerOptions.NewTextHandler(w)
		}
	}
}
