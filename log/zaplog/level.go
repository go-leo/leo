package zaplog

import (
	"go.uber.org/zap"

	"codeup.aliyun.com/qimao/leo/leo/log"
)

func LevelAdapt(level log.Level) zap.AtomicLevel {
	switch level.Code() {
	case log.LevelDebugCode:
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	case log.LevelInfoCode:
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	case log.LevelWarnCode:
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	case log.LevelErrorCode:
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	case log.LevelPanicCode:
		return zap.NewAtomicLevelAt(zap.PanicLevel)
	case log.LevelFatalCode:
		return zap.NewAtomicLevelAt(zap.FatalLevel)
	default:
		return zap.NewAtomicLevel()
	}
}
