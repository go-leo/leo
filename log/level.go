package log

import "net/http"

type Level string

const (
	Debug Level = "debug"
	Info  Level = "info"
	Warn  Level = "warn"
	Error Level = "error"
	Panic Level = "panic"
	Fatal Level = "fatal"
)

type Leveler interface {
	// SetLevel enable this Logger for the specified Level.
	SetLevel(level Level)
	// GetLevel return the Level of the Logger
	GetLevel() Level
	// EnableDebug enable this Logger for the Debug Level.
	EnableDebug()
	// IsDebugEnabled return True if this Logger is enabled for the Debug Level, false otherwise.
	IsDebugEnabled() bool
	// EnableInfo enable this Logger for the Info Level.
	EnableInfo()
	// IsInfoEnabled return True if this Logger is enabled for the Info Level, false otherwise.
	IsInfoEnabled() bool
	// EnableWarn enable this Logger for the Warn Level.
	EnableWarn()
	// IsWarnEnabled return True if this Logger is enabled for the Warn Level, false otherwise.
	IsWarnEnabled() bool
	// EnableError enable this Logger for the Error Level.
	EnableError()
	// IsErrorEnabled return True if this Logger is enabled for the Error Level, false otherwise.
	IsErrorEnabled() bool
	// EnablePanic enable this Logger for the Panic Level.
	EnablePanic()
	// IsPanicEnabled return True if this Logger is enabled for the Panic Level, false otherwise.
	IsPanicEnabled() bool
	// EnableFatal enable this Logger for the Fatal Level.
	EnableFatal()
	// IsFatalEnabled return True if this Logger is enabled for the Fatal Level, false otherwise.
	IsFatalEnabled() bool
	// Handler is a http endpoint that can report on or change the current logging Level.
	// The GET request returns a JSON description of the current logging Level like
	// The PUT request changes the logging Level. It is perfectly safe to change the
	// logging Level while a program is running.
	http.Handler
}
