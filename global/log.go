package global

import (
	"os"
	"sync"

	"codeup.aliyun.com/qimao/leo/leo/log"
	"codeup.aliyun.com/qimao/leo/leo/log/slog"
)

var (
	global    log.Logger = slog.New(slog.LevelAdapt(log.LevelDebug), slog.JSON())
	globalLog log.Logger = global.SkipCaller(4)
	logLocker sync.RWMutex
)

// 支持两种模式：
// 1. 通过global.SetLogger()设置全局logger，实现in current process打印log
// 2. 当前包内部定义info,debug等函数用于非上下文场景便捷使用

func Logger() log.Logger {
	logLocker.RLock()
	defer logLocker.RUnlock()
	return global
}

func SetLogger(l log.Logger) {
	logLocker.Lock()
	defer logLocker.Unlock()
	global = l
}

// Debug logs a message at debug level.
func Debug(a ...interface{}) {
	globalLog.Debug(a...)
}

// Debugf logs a message at debug level.
func Debugf(format string, a ...interface{}) {
	globalLog.Debugf(format, a...)

}

// Info logs a message at info level.
func Info(a ...interface{}) {
	// globalLog.SkipCaller(4).Info(a...)
	globalLog.Info(a...)

}

// Infof logs a message at info level.
func Infof(format string, a ...interface{}) {
	globalLog.Infof(format, a...)

}

// Warn logs a message at warn level.
func Warn(a ...interface{}) {
	globalLog.Warn(a...)

}

// Warnf logs a message at warnf level.
func Warnf(format string, a ...interface{}) {
	globalLog.Warnf(format, a...)

}

// Error logs a message at error level.
func Error(a ...interface{}) {
	globalLog.Error(a...)

}

// Errorf logs a message at error level.
func Errorf(format string, a ...interface{}) {
	globalLog.Errorf(format, a...)

}

// Fatal logs a message at fatal level.
func Fatal(a ...interface{}) {
	globalLog.Fatal(a...)

	os.Exit(1)
}

// Fatalf logs a message at fatal level.
func Fatalf(format string, a ...interface{}) {
	globalLog.Fatalf(format, a...)
	os.Exit(1)
}
