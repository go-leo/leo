package log

type PanicLogger interface {
	// Panic formats using the default formats for its operands and logs a message at Panic level, then panic.
	Panic(args ...any)
	// Panicf formats according to a format specifier and log a templated message at Panic level, then panic.
	Panicf(template string, args ...any)
	// PanicF logs a message at Panic level, then panic. The message includes any fields passed
	// at the log site, as well as any fields accumulated on the logger.
	PanicF(fields ...Field)
}
