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

var (
	debugLevel = level{name: "debug", code: -1}
	infoLevel  = level{name: "info", code: 0}
	warnLevel  = level{name: "warn", code: 1}
	errorLevel = level{name: "error", code: 2}
	panicLevel = level{name: "panic", code: 3}
	fatalLevel = level{name: "fatal", code: 4}
)

func Debug() Level {
	return debugLevel
}

func Info() Level {
	return infoLevel
}

func Warn() Level {
	return warnLevel
}
func Error() Level {
	return errorLevel
}
func Panic() Level {
	return panicLevel
}
func Fatal() Level {
	return fatalLevel
}
