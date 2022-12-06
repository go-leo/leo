package pubsub

import (
	"github.com/ThreeDotsLabs/watermill"

	"github.com/go-leo/stringx"

	"github.com/go-leo/leo/v2/log"
)

var _ watermill.LoggerAdapter = new(logger)

type logger struct {
	l log.Logger
}

func NewLogger(l log.Logger) *logger {
	return &logger{l: l}
}

func (l *logger) Error(msg string, err error, fields watermill.LogFields) {
	l.l.ErrorF(l.toLogF(msg, err, fields)...)
}

func (l *logger) Info(msg string, fields watermill.LogFields) {
	l.l.InfoF(l.toLogF(msg, nil, fields)...)
}

func (l *logger) Debug(msg string, fields watermill.LogFields) {
	l.l.DebugF(l.toLogF(msg, nil, fields)...)
}

func (l *logger) Trace(msg string, fields watermill.LogFields) {
	l.l.DebugF(l.toLogF(msg, nil, fields)...)
}

func (l *logger) With(fields watermill.LogFields) watermill.LoggerAdapter {
	fs := l.toLogF("", nil, fields)
	return &logger{l: l.l.With(fs...)}
}

func (l *logger) toLogF(msg string, err error, fields watermill.LogFields) []log.F {
	var fs []log.F
	if stringx.IsNotBlank(msg) {
		fs = append(fs, log.F{K: "msg", V: msg})
	}
	if err != nil {
		fs = append(fs, log.F{K: "err", V: err.Error()})
	}
	for k, v := range fields {
		fs = append(fs, log.F{K: k, V: v})
	}
	return fs
}
