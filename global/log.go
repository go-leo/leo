package global

import (
	"sync"

	"codeup.aliyun.com/qimao/leo/leo/log"
	"codeup.aliyun.com/qimao/leo/leo/log/slog"
)

var (
	logger    log.Logger = slog.New(slog.LevelAdapt(log.LevelDebug), slog.Console(true), slog.PlainText())
	logLocker sync.RWMutex
)

func Logger() log.Logger {
	logLocker.RLock()
	defer logLocker.RUnlock()
	return logger
}

func SetLogger(l log.Logger) func() {
	logLocker.Lock()
	defer logLocker.Unlock()
	prev := logger
	logger = l
	return func() { SetLogger(prev) }
}
