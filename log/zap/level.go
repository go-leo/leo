package zap

import (
	"go.uber.org/zap"

	"github.com/go-leo/leo/v2/log"
)

func LevelAdapt(level log.Level) zap.AtomicLevel {
	switch level {
	case log.Debug:
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	case log.Info:
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	case log.Warn:
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	case log.Error:
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	case log.Panic:
		return zap.NewAtomicLevelAt(zap.PanicLevel)
	case log.Fatal:
		return zap.NewAtomicLevelAt(zap.FatalLevel)
	default:
		return zap.NewAtomicLevel()
	}
}
