package cron

import (
	"fmt"

	"github.com/robfig/cron/v3"

	"github.com/go-leo/leo/log"
)

type logger struct {
	l log.Logger
}

func NewLogger(l log.Logger) cron.Logger {
	return &logger{l: l}
}

func (l *logger) Info(msg string, keysAndValues ...any) {
	fs := []log.F{{K: "msg", V: msg}}
	for i := 0; i < len(keysAndValues); i += 2 {
		k := fmt.Sprint(keysAndValues[i])
		v := fmt.Sprint(keysAndValues[i+1])
		fs = append(fs, log.F{K: k, V: v})
	}
	l.l.DebugF(fs...)
}

func (l *logger) Error(err error, msg string, keysAndValues ...any) {
	fs := []log.F{{K: "msg", V: msg}, {K: "err", V: err}}
	for i := 0; i < len(keysAndValues); i += 2 {
		k := fmt.Sprint(keysAndValues[i])
		v := fmt.Sprint(keysAndValues[i+1])
		fs = append(fs, log.F{K: k, V: v})
	}
	l.l.InfoF(fs...)
}
