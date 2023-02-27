package log

type Level interface {
	Name() string
	Code() int
}

type level struct {
	name string
	code int
}

func (l level) Name() string {
	return l.name
}

func (l level) Code() int {
	return l.code
}

const (
	LevelDebugCode = -1
	LevelInfoCode  = 0
	LevelWarnCode  = 1
	LevelErrorCode = 2
	LevelPanicCode = 3
	LevelFatalCode = 4
)

const (
	LevelDebugName = "debug"
	LevelInfoName  = "info"
	LevelWarnName  = "warn"
	LevelErrorName = "error"
	LevelPanicName = "panic"
	LevelFatalName = "fatal"
)

var (
	LevelDebug = level{name: LevelDebugName, code: LevelDebugCode}
	LevelInfo  = level{name: LevelInfoName, code: LevelInfoCode}
	LevelWarn  = level{name: LevelWarnName, code: LevelWarnCode}
	LevelError = level{name: LevelErrorName, code: LevelErrorCode}
	LevelPanic = level{name: LevelPanicName, code: LevelPanicCode}
	LevelFatal = level{name: LevelFatalName, code: LevelFatalCode}
)
