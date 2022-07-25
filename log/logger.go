package log

import "net/http"

// F is field
type F struct {
	// k is key, type is string
	K string
	// V is value, type is any
	V any
}

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
	// EnableDebug enable this Logger for the Debug level.
	EnableDebug()
	// IsDebugEnabled return True if this Logger is enabled for the Debug level, false otherwise.
	IsDebugEnabled() bool
	// EnableInfo enable this Logger for the Info level.
	EnableInfo()
	// IsInfoEnabled return True if this Logger is enabled for the Info level, false otherwise.
	IsInfoEnabled() bool
	// EnableWarn enable this Logger for the Warn level.
	EnableWarn()
	// IsWarnEnabled return True if this Logger is enabled for the Warn level, false otherwise.
	IsWarnEnabled() bool
	// EnableError enable this Logger for the Error level.
	EnableError()
	// IsErrorEnabled return True if this Logger is enabled for the Error level, false otherwise.
	IsErrorEnabled() bool
	// EnablePanic enable this Logger for the Panic level.
	EnablePanic()
	// IsPanicEnabled return True if this Logger is enabled for the Panic level, false otherwise.
	IsPanicEnabled() bool
	// EnableFatal enable this Logger for the Fatal level.
	EnableFatal()
	// IsFatalEnabled return True if this Logger is enabled for the Fatal level, false otherwise.
	IsFatalEnabled() bool
	// Handler is a http endpoint that can report on or change the current logging level.
	// The GET request returns a JSON description of the current logging level like
	// The PUT request changes the logging level. It is perfectly safe to change the
	// logging level while a program is running.
	http.Handler
}

// Logger is a interface to log messages.
type Logger interface {
	Leveler
	DebugLogger
	InfoLogger
	WarnLogger
	ErrorLogger
	PanicLogger
	FatalLogger
	CallerSkipable
	FieldAddable
	Cloneable
}

type DebugLogger interface {
	// Debug formats using the default formats for its operands and logs a message at Debug level.
	Debug(args ...any)
	// Debugf formats according to a format specifier and log a templated message at Debug level.
	Debugf(template string, args ...any)
	// DebugF logs a message at Debug level. The message includes any fields passed
	// at the log site, as well as any fields accumulated on the logger.
	DebugF(fields ...F)
}

type InfoLogger interface {
	// Info formats using the default formats for its operands and logs a message at Info level.
	Info(args ...any)
	// Infof formats according to a format specifier and log a templated message at Info level.
	Infof(template string, args ...any)
	// InfoF logs a message at Info level. The message includes any fields passed
	// at the log site, as well as any fields accumulated on the logger.
	InfoF(fields ...F)
}

type WarnLogger interface {
	// Warn formats using the default formats for its operands and logs a message at Warn level.
	Warn(args ...any)
	// Warnf formats according to a format specifier and log a templated message at Warn level.
	Warnf(template string, args ...any)
	// WarnF logs a message at Warn level. The message includes any fields passed
	// at the log site, as well as any fields accumulated on the logger.
	WarnF(fields ...F)
}

type ErrorLogger interface {
	// Error formats using the default formats for its operands and logs a message at Error level.
	Error(args ...any)
	// Errorf formats according to a format specifier and log a templated message at Error level.
	Errorf(template string, args ...any)
	// ErrorF logs a message at Error level. The message includes any fields passed
	// at the log site, as well as any fields accumulated on the logger.
	ErrorF(fields ...F)
}

type PanicLogger interface {
	// Panic formats using the default formats for its operands and logs a message at Panic level, then panic.
	Panic(args ...any)
	// Panicf formats according to a format specifier and log a templated message at Panic level, then panic.
	Panicf(template string, args ...any)
	// PanicF logs a message at Panic level, then panic. The message includes any fields passed
	// at the log site, as well as any fields accumulated on the logger.
	PanicF(fields ...F)
}

type FatalLogger interface {
	// Fatal formats using the default formats for its operands and logs a message at Fatal level, then calls os.Exit(1).
	Fatal(args ...any)
	// Fatalf formats according to a format specifier and log a templated message at Fatal level, then calls os.Exit(1).
	Fatalf(template string, args ...any)
	// FatalF logs a message at Fatal level, then calls os.Exit(1). The message includes any fields passed
	// at the log site, as well as any fields accumulated on the logger.
	FatalF(fields ...F)
}

type CallerSkipable interface {
	// SkipCaller returns a Logger that will offset the call stack by the specified number of frames
	// when logging call site information.
	SkipCaller(depth int) Logger
}

type FieldAddable interface {
	// With returns a new Logger with additional fields.
	With(fields ...F) Logger
}

type Cloneable interface {
	// Clone return new Logger instance.
	Clone() Logger
}
