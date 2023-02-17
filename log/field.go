package log

type Field interface {
	Key() string
	Value() any
}

// F is field
type F struct {
	// k is key, type is string
	K string
	// V is value, type is any
	V any
}

func (f *F) Key() string {
	return f.K
}

func (f *F) Value() any {
	return f.V
}

type FieldAppender interface {
	// With returns a new Logger with additional fields.
	With(fields ...Field) Logger
}
