package log

type FatalLogger interface {
	// Fatal formats using the default formats for its operands and logs a message at Fatal level, then calls os.Exit(1).
	Fatal(args ...any)
	// Fatalf formats according to a format specifier and log a templated message at Fatal level, then calls os.Exit(1).
	Fatalf(template string, args ...any)
	// FatalF logs a message at Fatal level, then calls os.Exit(1). The message includes any fields passed
	// at the log site, as well as any fields accumulated on the logger.
	FatalF(fields ...Field)
}
