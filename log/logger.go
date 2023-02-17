package log

// Logger is a interface to log messages.
type Logger interface {
	Leveler
	CallerSkipper
	FieldAppender
	Cloner
	DebugLogger
	InfoLogger
	WarnLogger
	ErrorLogger
	PanicLogger
	FatalLogger
}
