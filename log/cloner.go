package log

type Cloner interface {
	// Clone return new Logger instance.
	Clone() Logger
}
