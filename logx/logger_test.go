package logx

import (
	"context"
	"github.com/go-kit/log/level"
	"os"
	"testing"
)

func TestNewError(t *testing.T) {
	logger := New(os.Stdout, Logfmt(), Level(level.WarnValue()), Timestamp(), Caller(0), Sync())
	logger.Log("ss", "cc")
	logger.Log("qq", "ww", "level", level.DebugValue())
	logger.Log("qq", "ww", "level", level.ErrorValue())
}

func TestDefaultDebug(t *testing.T) {
	Replace(New(os.Stdout, Logfmt(), Level(level.DebugValue()), Timestamp(), Caller(2), Sync()))

	Debug(context.Background(), "method", "Debug")
	Debugf(context.Background(), "method: %s", "Debugf")
	Debugln(context.Background(), "method", "Debugln")

	Info(context.Background(), "method", "Info")
	Infof(context.Background(), "method: %s", "Infof")
	Infoln(context.Background(), "method", "Infoln")

	Warn(context.Background(), "method", "Warn")
	Warnf(context.Background(), "method: %s", "Warnf")
	Warnln(context.Background(), "method", "Warnln")

	Error(context.Background(), "method", "Error")
	Errorf(context.Background(), "method: %s", "Errorf")
	Errorln(context.Background(), "method", "Errorln")
}

func TestDefaultInfo(t *testing.T) {
	Replace(New(os.Stdout, Logfmt(), Level(level.InfoValue()), Timestamp(), Caller(2), Sync()))

	Debug(context.Background(), "method", "Debug")
	Debugf(context.Background(), "method: %s", "Debugf")
	Debugln(context.Background(), "method", "Debugln")

	Info(context.Background(), "method", "Info")
	Infof(context.Background(), "method: %s", "Infof")
	Infoln(context.Background(), "method", "Infoln")

	Warn(context.Background(), "method", "Warn")
	Warnf(context.Background(), "method: %s", "Warnf")
	Warnln(context.Background(), "method", "Warnln")

	Error(context.Background(), "method", "Error")
	Errorf(context.Background(), "method: %s", "Errorf")
	Errorln(context.Background(), "method", "Errorln")
}

func TestDefaultWarn(t *testing.T) {
	Replace(New(os.Stdout, Logfmt(), Level(level.WarnValue()), Timestamp(), Caller(2), Sync()))

	Debug(context.Background(), "method", "Debug")
	Debugf(context.Background(), "method: %s", "Debugf")
	Debugln(context.Background(), "method", "Debugln")

	Info(context.Background(), "method", "Info")
	Infof(context.Background(), "method: %s", "Infof")
	Infoln(context.Background(), "method", "Infoln")

	Warn(context.Background(), "method", "Warn")
	Warnf(context.Background(), "method: %s", "Warnf")
	Warnln(context.Background(), "method", "Warnln")

	Error(context.Background(), "method", "Error")
	Errorf(context.Background(), "method: %s", "Errorf")
	Errorln(context.Background(), "method", "Errorln")
}

func TestDefaultError(t *testing.T) {
	Replace(New(os.Stdout, Logfmt(), Level(level.ErrorValue()), Timestamp(), Caller(2), Sync()))

	Debug(context.Background(), "method", "Debug")
	Debugf(context.Background(), "method: %s", "Debugf")
	Debugln(context.Background(), "method", "Debugln")

	Info(context.Background(), "method", "Info")
	Infof(context.Background(), "method: %s", "Infof")
	Infoln(context.Background(), "method", "Infoln")

	Warn(context.Background(), "method", "Warn")
	Warnf(context.Background(), "method: %s", "Warnf")
	Warnln(context.Background(), "method", "Warnln")

	Error(context.Background(), "method", "Error")
	Errorf(context.Background(), "method: %s", "Errorf")
	Errorln(context.Background(), "method", "Errorln")
}
