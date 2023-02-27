package log

import (
	"net/http"
)

var _ Logger = new(Discard)

type Discard struct{}

func (d Discard) SetLevel(level Level) {
}

func (d Discard) GetLevel() Level {
	return LevelDebug
}

func (d Discard) EnableDebug() {
}

func (d Discard) IsDebugEnabled() bool {
	return true
}

func (d Discard) EnableInfo() {
}

func (d Discard) IsInfoEnabled() bool {
	return true
}

func (d Discard) EnableWarn() {
}

func (d Discard) IsWarnEnabled() bool {
	return true
}

func (d Discard) EnableError() {
}

func (d Discard) IsErrorEnabled() bool {
	return true
}

func (d Discard) EnablePanic() {
}

func (d Discard) IsPanicEnabled() bool {
	return true
}

func (d Discard) EnableFatal() {
}

func (d Discard) IsFatalEnabled() bool {
	return true
}

func (d Discard) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
}

func (d Discard) SkipCaller(depth int) Logger {
	return d
}

func (d Discard) With(fields ...Field) Logger {
	return d
}

func (d Discard) Clone() Logger {
	return d
}

func (d Discard) Debug(args ...any) {
}

func (d Discard) Debugf(template string, args ...any) {
}

func (d Discard) DebugF(fields ...Field) {
}

func (d Discard) Info(args ...any) {
}

func (d Discard) Infof(template string, args ...any) {
}

func (d Discard) InfoF(fields ...Field) {
}

func (d Discard) Warn(args ...any) {
}

func (d Discard) Warnf(template string, args ...any) {
}

func (d Discard) WarnF(fields ...Field) {
}

func (d Discard) Error(args ...any) {
}

func (d Discard) Errorf(template string, args ...any) {
}

func (d Discard) ErrorF(fields ...Field) {
}

func (d Discard) Panic(args ...any) {
}

func (d Discard) Panicf(template string, args ...any) {
}

func (d Discard) PanicF(fields ...Field) {

}

func (d Discard) Fatal(args ...any) {

}

func (d Discard) Fatalf(template string, args ...any) {

}

func (d Discard) FatalF(fields ...Field) {

}
