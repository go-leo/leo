package slog

import (
	"os"
	"sync"

	"codeup.aliyun.com/qimao/leo/leo/log"
)

var (
	global    log.Logger
	globalLog log.Logger
	logLocker sync.RWMutex
)

func init() {
	SetLogger(New(LevelAdapt(log.LevelDebug), JSON()))
}

func SetLogger(l log.Logger) {
	logLocker.Lock()
	defer logLocker.Unlock()
	global = l
	globalLog = global.SkipCaller(4)
}

func Logger() log.Logger {
	logLocker.RLock()
	defer logLocker.RUnlock()
	return global
}

// Debug logs a message at debug level.
func Debug(a ...interface{}) {
	logLocker.RLock()
	globalLog.Debug(a...)
	logLocker.RUnlock()
}

// Debugf logs a message at debug level.
func Debugf(format string, a ...interface{}) {
	logLocker.RLock()
	globalLog.Debugf(format, a...)
	logLocker.RUnlock()
}

// Info logs a message at info level.
func Info(a ...interface{}) {
	logLocker.RLock()
	globalLog.Info(a...)
	logLocker.RUnlock()

}

// Infof logs a message at info level.
func Infof(format string, a ...interface{}) {
	logLocker.RLock()
	globalLog.Infof(format, a...)
	logLocker.RUnlock()
}

// Warn logs a message at warn level.
func Warn(a ...interface{}) {
	logLocker.RLock()
	globalLog.Warn(a...)
	logLocker.RUnlock()
}

// Warnf logs a message at warnf level.
func Warnf(format string, a ...interface{}) {
	logLocker.RLock()
	globalLog.Warnf(format, a...)
	logLocker.RUnlock()
}

// Error logs a message at error level.
func Error(a ...interface{}) {
	logLocker.RLock()
	globalLog.Error(a...)
	logLocker.RUnlock()
}

// Errorf logs a message at error level.
func Errorf(format string, a ...interface{}) {
	logLocker.RLock()
	globalLog.Errorf(format, a...)
	logLocker.RUnlock()
}

// Fatal logs a message at fatal level.
func Fatal(a ...interface{}) {
	logLocker.RLock()
	globalLog.Fatal(a...)
	logLocker.RUnlock()

	os.Exit(1)
}

// Fatalf logs a message at fatal level.
func Fatalf(format string, a ...interface{}) {
	logLocker.RLock()
	globalLog.Fatalf(format, a...)
	logLocker.RUnlock()
	os.Exit(1)
}
