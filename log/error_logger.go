package log

type ErrorLogger interface {
	// Error formats using the default formats for its operands and logs a message at Error level.
	Error(args ...any)
	// Errorf formats according to a format specifier and log a templated message at Error level.
	Errorf(template string, args ...any)
	// ErrorF logs a message at Error level. The message includes any fields passed
	// at the log site, as well as any fields accumulated on the logger.
	ErrorF(fields ...Field)
}
