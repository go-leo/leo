package slog

import (
	"golang.org/x/exp/slog"

	"codeup.aliyun.com/qimao/leo/leo/log"
)

var (
	sLogLevelDebug = slog.Level(log.LevelDebugCode)
	sLogLevelInfo  = slog.Level(log.LevelInfoCode)
	sLogLevelWarn  = slog.Level(log.LevelWarnCode)
	sLogLevelError = slog.Level(log.LevelErrorCode)
	sLogLevelPanic = slog.Level(log.LevelPanicCode)
	sLogLevelFatal = slog.Level(log.LevelFatalCode)
)

func LevelAdapt(level log.Level) *slog.LevelVar {
	switch level.Code() {
	case log.LevelDebugCode:
		lvl := &slog.LevelVar{}
		lvl.Set(sLogLevelDebug)
		return lvl
	case log.LevelInfoCode:
		lvl := &slog.LevelVar{}
		lvl.Set(sLogLevelInfo)
		return lvl
	case log.LevelWarnCode:
		lvl := &slog.LevelVar{}
		lvl.Set(sLogLevelWarn)
		return lvl
	case log.LevelErrorCode:
		lvl := &slog.LevelVar{}
		lvl.Set(sLogLevelError)
		return lvl
	case log.LevelPanicCode:
		lvl := &slog.LevelVar{}
		lvl.Set(sLogLevelPanic)
		return lvl
	case log.LevelFatalCode:
		lvl := &slog.LevelVar{}
		lvl.Set(sLogLevelFatal)
		return lvl
	default:
		return &slog.LevelVar{}
	}
}
