package log

type WarnLogger interface {
	// Warn formats using the default formats for its operands and logs a message at Warn level.
	Warn(args ...any)
	// Warnf formats according to a format specifier and log a templated message at Warn level.
	Warnf(template string, args ...any)
	// WarnF logs a message at Warn level. The message includes any fields passed
	// at the log site, as well as any fields accumulated on the logger.
	WarnF(fields ...Field)
}
