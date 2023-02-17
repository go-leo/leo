package log

type DebugLogger interface {
	// Debug formats using the default formats for its operands and logs a message at Debug level.
	Debug(args ...any)
	// Debugf formats according to a format specifier and log a templated message at Debug level.
	Debugf(template string, args ...any)
	// DebugF logs a message at Debug level. The message includes any fields passed
	// at the log site, as well as any fields accumulated on the logger.
	DebugF(fields ...Field)
}
