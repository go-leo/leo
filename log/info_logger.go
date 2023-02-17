package log

type InfoLogger interface {
	// Info formats using the default formats for its operands and logs a message at Info level.
	Info(args ...any)
	// Infof formats according to a format specifier and log a templated message at Info level.
	Infof(template string, args ...any)
	// InfoF logs a message at Info level. The message includes any fields passed
	// at the log site, as well as any fields accumulated on the logger.
	InfoF(fields ...Field)
}
