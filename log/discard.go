package log

import (
	"net/http"
)

var _ Logger = new(Discard)

type Discard struct{}

func (l Discard) EnableDebug() {}

func (l Discard) IsDebugEnabled() bool { return false }

func (l Discard) Debug(args ...any) {}

func (l Discard) Debugf(template string, args ...any) {}

func (l Discard) DebugF(fields ...F) {}

func (l Discard) EnableInfo() {}

func (l Discard) IsInfoEnabled() bool { return false }

func (l Discard) Info(args ...any) {}

func (l Discard) Infof(template string, args ...any) {}

func (l Discard) InfoF(fields ...F) {}

func (l Discard) EnableWarn() {}

func (l Discard) IsWarnEnabled() bool { return false }

func (l Discard) Warn(args ...any) {}

func (l Discard) Warnf(template string, args ...any) {}

func (l Discard) WarnF(fields ...F) {}

func (l Discard) EnableError() {}

func (l Discard) IsErrorEnabled() bool { return false }

func (l Discard) Error(args ...any) {}

func (l Discard) Errorf(template string, args ...any) {}

func (l Discard) ErrorF(fields ...F) {}

func (l Discard) EnablePanic() {}

func (l Discard) IsPanicEnabled() bool { return false }

func (l Discard) Panic(args ...any) {}

func (l Discard) Panicf(template string, args ...any) {}

func (l Discard) PanicF(fields ...F) {}

func (l Discard) EnableFatal() {}

func (l Discard) IsFatalEnabled() bool { return false }

func (l Discard) Fatal(args ...any) {}

func (l Discard) Fatalf(template string, args ...any) {}

func (l Discard) FatalF(fields ...F) {}

func (l Discard) SkipCaller(depth int) Logger { return l }

func (l Discard) With(fields ...F) Logger { return l }

func (l Discard) Clone() Logger { return l }

func (l Discard) ServeHTTP(writer http.ResponseWriter, request *http.Request) {}
