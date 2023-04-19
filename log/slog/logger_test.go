package slog_test

import (
	"testing"

	"codeup.aliyun.com/qimao/leo/leo/log"
	"codeup.aliyun.com/qimao/leo/leo/log/slog"
)

func TestDebugSlog(t *testing.T) {
	logger := slog.New(slog.LevelAdapt(log.LevelDebug))
	if !logger.IsDebugEnabled() {
		t.Errorf("debug level failed")
		return
	}
	logger.Debug("this is Debug log")
	logger.Debugf("this is %s log", "Debugf")

	if !logger.IsInfoEnabled() {
		t.Errorf("info level failed")
		return
	}
	logger.Info("this is Info log")
	logger.Infof("this is %s log", "Infof")

	if !logger.IsWarnEnabled() {
		t.Errorf("warn level failed")
		return
	}
	logger.Warn("this is Warn log")
	logger.Warnf("this is %s log", "Warnf")

	if !logger.IsErrorEnabled() {
		t.Errorf("error level failed")
		return
	}
	logger.Error("this is Error log")
	logger.Errorf("this is %s log", "Errorf")

	logger.With(log.F{K: "TraceID", V: "1234567"})

	if !logger.IsDebugEnabled() {
		t.Errorf("debug level failed")
		return
	}
	logger.Debug("this is with traceID Debug log")
	logger.Debugf("this is  with traceID %s log", "Debugf")

	if !logger.IsInfoEnabled() {
		t.Errorf("info level failed")
		return
	}
	logger.Info("this is with traceID Info log")
	logger.Infof("this is with traceID %s log", "Infof")

	if !logger.IsWarnEnabled() {
		t.Errorf("warn level failed")
		return
	}
	logger.Warn("this is with traceID Warn log")
	logger.Warnf("this is with traceID %s log", "Warnf")

	if !logger.IsErrorEnabled() {
		t.Errorf("error level failed")
		return
	}
	logger.Error("this is with traceID Error log")
	logger.Errorf("this is with traceID %s log", "Errorf")

	logger.With(log.F{K: "user", V: "jax"}, log.F{K: "platform", V: "Android"})
	if !logger.IsDebugEnabled() {
		t.Errorf("debug level failed")
		return
	}
	logger.Debug("this is with Fields Debug log")
	logger.Debugf("this is with Fields %s log", "Debugf")

	if !logger.IsInfoEnabled() {
		t.Errorf("info level failed")
		return
	}
	logger.Info("this is with Fields Info log")
	logger.Infof("this is with Fields %s log", "Infof")

	if !logger.IsWarnEnabled() {
		t.Errorf("warn level failed")
		return
	}
	logger.Warn("this is with Fields Warn log")
	logger.Warnf("this is with Fields %s log", "Warnf")

	if !logger.IsErrorEnabled() {
		t.Errorf("error level failed")
		return
	}
	logger.Error("this is with Fields Error log")
	logger.Errorf("this is with Fields %s log", "Errorf")
}

func TestInfoSlog(t *testing.T) {
	logger := slog.New(slog.LevelAdapt(log.LevelInfo))
	if logger.IsDebugEnabled() {
		t.Errorf("debug level failed")
		return
	}
	logger.Debug("this is Debug log")
	logger.Debugf("this is %s log", "Debugf")

	if !logger.IsInfoEnabled() {
		t.Errorf("info level failed")
		return
	}
	logger.Info("this is Info log")
	logger.Infof("this is %s log", "Infof")

	if !logger.IsWarnEnabled() {
		t.Errorf("warn level failed")
		return
	}
	logger.Warn("this is Warn log")
	logger.Warnf("this is %s log", "Warnf")

	if !logger.IsErrorEnabled() {
		t.Errorf("error level failed")
		return
	}
	logger.Error("this is Error log")
	logger.Errorf("this is %s log", "Errorf")

	logger.With(log.F{K: "TraceID", V: "1234567"})

	if logger.IsDebugEnabled() {
		t.Errorf("debug level failed")
		return
	}
	logger.Debug("this is with traceID Debug log")
	logger.Debugf("this is  with traceID %s log", "Debugf")

	if !logger.IsInfoEnabled() {
		t.Errorf("info level failed")
		return
	}
	logger.Info("this is with traceID Info log")
	logger.Infof("this is with traceID %s log", "Infof")

	if !logger.IsWarnEnabled() {
		t.Errorf("warn level failed")
		return
	}
	logger.Warn("this is with traceID Warn log")
	logger.Warnf("this is with traceID %s log", "Warnf")

	if !logger.IsErrorEnabled() {
		t.Errorf("error level failed")
		return
	}
	logger.Error("this is with traceID Error log")
	logger.Errorf("this is with traceID %s log", "Errorf")

	logger.With(log.F{K: "user", V: "jax"}, log.F{K: "platform", V: "Android"})
	if logger.IsDebugEnabled() {
		t.Errorf("debug level failed")
		return
	}
	logger.Debug("this is with Fields Debug log")
	logger.Debugf("this is with Fields %s log", "Debugf")

	if !logger.IsInfoEnabled() {
		t.Errorf("info level failed")
		return
	}
	logger.Info("this is with Fields Info log")
	logger.Infof("this is with Fields %s log", "Infof")

	if !logger.IsWarnEnabled() {
		t.Errorf("warn level failed")
		return
	}
	logger.Warn("this is with Fields Warn log")
	logger.Warnf("this is with Fields %s log", "Warnf")

	if !logger.IsErrorEnabled() {
		t.Errorf("error level failed")
		return
	}
	logger.Error("this is with Fields Error log")
	logger.Errorf("this is with Fields %s log", "Errorf")
}

func TestWarnSlog(t *testing.T) {
	logger := slog.New(slog.LevelAdapt(log.LevelWarn))
	if logger.IsDebugEnabled() {
		t.Errorf("debug level failed")
		return
	}
	logger.Debug("this is Debug log")
	logger.Debugf("this is %s log", "Debugf")

	if logger.IsInfoEnabled() {
		t.Errorf("info level failed")
		return
	}
	logger.Info("this is Info log")
	logger.Infof("this is %s log", "Infof")

	if !logger.IsWarnEnabled() {
		t.Errorf("warn level failed")
		return
	}
	logger.Warn("this is Warn log")
	logger.Warnf("this is %s log", "Warnf")

	if !logger.IsErrorEnabled() {
		t.Errorf("error level failed")
		return
	}
	logger.Error("this is Error log")
	logger.Errorf("this is %s log", "Errorf")

	logger.With(log.F{K: "TraceID", V: "1234567"})

	if logger.IsDebugEnabled() {
		t.Errorf("debug level failed")
		return
	}
	logger.Debug("this is with traceID Debug log")
	logger.Debugf("this is  with traceID %s log", "Debugf")

	if logger.IsInfoEnabled() {
		t.Errorf("info level failed")
		return
	}
	logger.Info("this is with traceID Info log")
	logger.Infof("this is with traceID %s log", "Infof")

	if !logger.IsWarnEnabled() {
		t.Errorf("warn level failed")
		return
	}
	logger.Warn("this is with traceID Warn log")
	logger.Warnf("this is with traceID %s log", "Warnf")

	if !logger.IsErrorEnabled() {
		t.Errorf("error level failed")
		return
	}
	logger.Error("this is with traceID Error log")
	logger.Errorf("this is with traceID %s log", "Errorf")

	logger.With(log.F{K: "user", V: "jax"}, log.F{K: "platform", V: "Android"})
	if logger.IsDebugEnabled() {
		t.Errorf("debug level failed")
		return
	}
	logger.Debug("this is with Fields Debug log")
	logger.Debugf("this is with Fields %s log", "Debugf")

	if logger.IsInfoEnabled() {
		t.Errorf("info level failed")
		return
	}
	logger.Info("this is with Fields Info log")
	logger.Infof("this is with Fields %s log", "Infof")

	if !logger.IsWarnEnabled() {
		t.Errorf("warn level failed")
		return
	}
	logger.Warn("this is with Fields Warn log")
	logger.Warnf("this is with Fields %s log", "Warnf")

	if !logger.IsErrorEnabled() {
		t.Errorf("error level failed")
		return
	}
	logger.Error("this is with Fields Error log")
	logger.Errorf("this is with Fields %s log", "Errorf")
}

func TestErrorSlog(t *testing.T) {
	var logger log.Logger
	logger = slog.New(slog.LevelAdapt(log.LevelError))
	if logger.IsDebugEnabled() {
		t.Errorf("debug level failed")
		return
	}
	logger.Debug("this is Debug log")
	logger.Debugf("this is %s log", "Debugf")

	if logger.IsInfoEnabled() {
		t.Errorf("info level failed")
		return
	}
	logger.Info("this is Info log")
	logger.Infof("this is %s log", "Infof")

	if logger.IsWarnEnabled() {
		t.Errorf("warn level failed")
		return
	}
	logger.Warn("this is Warn log")
	logger.Warnf("this is %s log", "Warnf")

	if !logger.IsErrorEnabled() {
		t.Errorf("error level failed")
		return
	}
	logger.Error("this is Error log")
	logger.Errorf("this is %s log", "Errorf")

	logger.With(log.F{K: "TraceID", V: "1234567"})

	if logger.IsDebugEnabled() {
		t.Errorf("debug level failed")
		return
	}
	logger.Debug("this is with traceID Debug log")
	logger.Debugf("this is  with traceID %s log", "Debugf")

	if logger.IsInfoEnabled() {
		t.Errorf("info level failed")
		return
	}
	logger.Info("this is with traceID Info log")
	logger.Infof("this is with traceID %s log", "Infof")

	if logger.IsWarnEnabled() {
		t.Errorf("warn level failed")
		return
	}
	logger.Warn("this is with traceID Warn log")
	logger.Warnf("this is with traceID %s log", "Warnf")

	if !logger.IsErrorEnabled() {
		t.Errorf("error level failed")
		return
	}
	logger.Error("this is with traceID Error log")
	logger.Errorf("this is with traceID %s log", "Errorf")

	logger = logger.With(log.F{K: "user", V: "jax"}, log.F{K: "platform", V: "Android"})
	if logger.IsDebugEnabled() {
		t.Errorf("debug level failed")
		return
	}
	logger.Debug("this is with Fields Debug log")
	logger.Debugf("this is with Fields %s log", "Debugf")

	if logger.IsInfoEnabled() {
		t.Errorf("info level failed")
		return
	}
	logger.Info("this is with Fields Info log")
	logger.Infof("this is with Fields %s log", "Infof")

	if logger.IsWarnEnabled() {
		t.Errorf("warn level failed")
		return
	}
	logger.Warn("this is with Fields Warn log")
	logger.Warnf("this is with Fields %s log", "Warnf")

	if !logger.IsErrorEnabled() {
		t.Errorf("error level failed")
		return
	}
	logger.Error("this is with Fields Error log")
	logger.Errorf("this is with Fields %s log", "Errorf")
}
