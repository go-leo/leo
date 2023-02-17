package log

import (
	"context"
	"net/http"
)

var _ Logger = new(Discard)

type Discard struct{}

func (l Discard) Context() context.Context {
	return context.Background()
}

func (l Discard) WithContext(context.Context) Logger {
	return l
}

func (l Discard) SetLevel(Level) {}

func (l Discard) GetLevel() Level {
	return Debug
}

func (l Discard) EnableDebug() {}

func (l Discard) IsDebugEnabled() bool { return false }

func (l Discard) Debug(...any) {}

func (l Discard) Debugf(string, ...any) {}

func (l Discard) DebugF(...Field) {}

func (l Discard) EnableInfo() {}

func (l Discard) IsInfoEnabled() bool { return false }

func (l Discard) Info(...any) {}

func (l Discard) Infof(string, ...any) {}

func (l Discard) InfoF(...Field) {}

func (l Discard) EnableWarn() {}

func (l Discard) IsWarnEnabled() bool { return false }

func (l Discard) Warn(...any) {}

func (l Discard) Warnf(string, ...any) {}

func (l Discard) WarnF(...Field) {}

func (l Discard) EnableError() {}

func (l Discard) IsErrorEnabled() bool { return false }

func (l Discard) Error(...any) {}

func (l Discard) Errorf(string, ...any) {}

func (l Discard) ErrorF(...Field) {}

func (l Discard) EnablePanic() {}

func (l Discard) IsPanicEnabled() bool { return false }

func (l Discard) Panic(...any) {}

func (l Discard) Panicf(string, ...any) {}

func (l Discard) PanicF(...Field) {}

func (l Discard) EnableFatal() {}

func (l Discard) IsFatalEnabled() bool { return false }

func (l Discard) Fatal(...any) {}

func (l Discard) Fatalf(string, ...any) {}

func (l Discard) FatalF(...Field) {}

func (l Discard) SkipCaller(int) Logger { return l }

func (l Discard) With(...Field) Logger { return l }

func (l Discard) Clone() Logger { return l }

func (l Discard) ServeHTTP(http.ResponseWriter, *http.Request) {}
