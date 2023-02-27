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

func (f F) Key() string {
	return f.K
}

func (f F) Value() any {
	return f.V
}

const (
	msgKey = "msg"
	errKey = "err"
)

func MsgField(msg string) Field {
	return &F{K: msgKey, V: msg}
}

func IsMsgField(f Field) bool {
	if f.Key() == msgKey {
		return true
	}
	return false
}

func ErrField(err error) Field {
	return &F{K: errKey, V: err}
}

func IsErrField(f Field) bool {
	if f.Key() == errKey {
		return true
	}
	return false
}
