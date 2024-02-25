package log

import "net/http"

// Logger is a interface to log messages.
type Logger interface {
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

	// SkipCaller returns a Logger that will offset the call stack by the specified number of frames
	// when logging call site information.
	SkipCaller(depth int) Logger

	// With returns a new Logger with additional fields.
	With(fields ...Field) Logger

	// Clone return new Logger instance.
	Clone() Logger

	// Debug formats using the default formats for its operands and logs a message at Debug level.
	Debug(args ...any)
	// Debugf formats according to a format specifier and log a templated message at Debug level.
	Debugf(template string, args ...any)
	// DebugF logs a message at Debug level. The message includes any fields passed
	// at the log site, as well as any fields accumulated on the logger.
	DebugF(fields ...Field)

	// Info formats using the default formats for its operands and logs a message at Info level.
	Info(args ...any)
	// Infof formats according to a format specifier and log a templated message at Info level.
	Infof(template string, args ...any)
	// InfoF logs a message at Info level. The message includes any fields passed
	// at the log site, as well as any fields accumulated on the logger.
	InfoF(fields ...Field)

	// Warn formats using the default formats for its operands and logs a message at Warn level.
	Warn(args ...any)
	// Warnf formats according to a format specifier and log a templated message at Warn level.
	Warnf(template string, args ...any)
	// WarnF logs a message at Warn level. The message includes any fields passed
	// at the log site, as well as any fields accumulated on the logger.
	WarnF(fields ...Field)

	// Error formats using the default formats for its operands and logs a message at Error level.
	Error(args ...any)
	// Errorf formats according to a format specifier and log a templated message at Error level.
	Errorf(template string, args ...any)
	// ErrorF logs a message at Error level. The message includes any fields passed
	// at the log site, as well as any fields accumulated on the logger.
	ErrorF(fields ...Field)

	// Panic formats using the default formats for its operands and logs a message at Panic level, then panic.
	Panic(args ...any)
	// Panicf formats according to a format specifier and log a templated message at Panic level, then panic.
	Panicf(template string, args ...any)
	// PanicF logs a message at Panic level, then panic. The message includes any fields passed
	// at the log site, as well as any fields accumulated on the logger.
	PanicF(fields ...Field)

	// Fatal formats using the default formats for its operands and logs a message at Fatal level, then calls os.Exit(1).
	Fatal(args ...any)
	// Fatalf formats according to a format specifier and log a templated message at Fatal level, then calls os.Exit(1).
	Fatalf(template string, args ...any)
	// FatalF logs a message at Fatal level, then calls os.Exit(1). The message includes any fields passed
	// at the log site, as well as any fields accumulated on the logger.
	FatalF(fields ...Field)
}
