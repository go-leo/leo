package slog

import "golang.org/x/exp/slog"

type Logger struct {
	*slog.Logger
}

func NewLogger(logger *slog.Logger) *Logger {
	return &Logger{Logger: logger}
}
