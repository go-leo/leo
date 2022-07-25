package global

import (
	"sync"

	"github.com/go-leo/leo/log"
	"github.com/go-leo/leo/log/zap"
)

var (
	logger    log.Logger = zap.New(zap.LevelAdapt(log.Debug), zap.Console(true), zap.JSON())
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

func initLogger() error {
	logConf := Configuration().Logger
	l := logConf.NewLogger()
	SetLogger(l)
	return nil
}
