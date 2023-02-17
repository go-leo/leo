package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/go-leo/leo/v2/log"
)

type options struct {
	Console     bool // Console 是否在终端输出
	FileOptions *fileOptions
	Fields      []zap.Field
	Encoder     zapcore.Encoder
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
		o.Encoder = zapcore.NewJSONEncoder(newEncoderConfig())
	}
	switch {
	case o.FileOptions != nil:
	case o.Console:
	default:
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
			l.Fields = append(l.Fields, zap.Any(field.Key(), field.Value()))
		}
	}
}

func JSON() Option {
	return func(o *options) {
		o.Encoder = zapcore.NewJSONEncoder(newEncoderConfig())
	}
}

func PlainText() Option {
	return func(o *options) {
		o.Encoder = zapcore.NewConsoleEncoder(newEncoderConfig())
	}
}

func newEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:       "msg",
		LevelKey:         "level",
		TimeKey:          "ts",
		NameKey:          "logger",
		CallerKey:        "caller",
		FunctionKey:      zapcore.OmitKey,
		StacktraceKey:    "stacktrace",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      zapcore.LowercaseLevelEncoder,
		EncodeTime:       zapcore.RFC3339TimeEncoder,
		EncodeDuration:   zapcore.SecondsDurationEncoder,
		EncodeCaller:     zapcore.ShortCallerEncoder,
		ConsoleSeparator: "\t",
	}
}
