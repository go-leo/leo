package logx

import (
	"github.com/go-kit/log"
	"os"
	"sync"
)

var (
	defaultLogger = log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	defaultMutex  sync.RWMutex
)

func Replace(logger log.Logger) log.Logger {
	if logger == nil {
		panic("logx: can not replace default logger with nil")
	}
	var l log.Logger
	defaultMutex.Lock()
	l = defaultLogger
	defaultLogger = logger
	defaultMutex.Unlock()
	return l
}

func L() log.Logger {
	var l log.Logger
	defaultMutex.RLock()
	l = defaultLogger
	defaultMutex.RUnlock()
	return l
}
