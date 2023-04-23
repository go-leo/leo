package log

import (
	"errors"
	"fmt"
	"strings"
)

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
	LevelDebug = &level{name: LevelDebugName, code: LevelDebugCode}
	LevelInfo  = &level{name: LevelInfoName, code: LevelInfoCode}
	LevelWarn  = &level{name: LevelWarnName, code: LevelWarnCode}
	LevelError = &level{name: LevelErrorName, code: LevelErrorCode}
	LevelPanic = &level{name: LevelPanicName, code: LevelPanicCode}
	LevelFatal = &level{name: LevelFatalName, code: LevelFatalCode}
)

type Level interface {
	Name() string
	Code() int
}

type level struct {
	name string
	code int
}

func (l *level) Name() string {
	return l.name
}

func (l *level) Code() int {
	return l.code
}

// MarshalText marshals the Level to text.
func (l *level) MarshalText() ([]byte, error) {
	return []byte(l.name), nil
}

// UnmarshalText unmarshals text to a level.
func (l *level) UnmarshalText(text []byte) error {
	if l == nil {
		return errors.New("can't unmarshal a nil Level")
	}
	if !l.unmarshalText(text) {
		return fmt.Errorf("unrecognized level: %q", text)
	}
	return nil
}

func (l *level) unmarshalText(text []byte) bool {
	switch strings.ToLower(string(text)) {
	case "debug":
		if l == nil {
			*l = level{}
		}
		l.code = LevelDebugCode
		l.name = LevelDebugName
	case "info", "": // make the zero value useful
		if l == nil {
			*l = level{}
		}
		l.code = LevelInfoCode
		l.name = LevelInfoName
	case "warn":
		if l == nil {
			*l = level{}
		}
		l.code = LevelWarnCode
		l.name = LevelWarnName
	case "error":
		if l == nil {
			*l = level{}
		}
		l.code = LevelErrorCode
		l.name = LevelErrorName
	case "panic":
		if l == nil {
			*l = level{}
		}
		l.code = LevelPanicCode
		l.name = LevelPanicName
	case "fatal":
		if l == nil {
			*l = level{}
		}
		l.code = LevelFatalCode
		l.name = LevelFatalName
	default:
		return false
	}
	return true
}

// ParseLevel parses a level based on the lower-case or all-caps ASCII
// representation of the log level. If the provided ASCII representation is
// invalid an error is returned.
//
// This is particularly useful when dealing with text input to configure log
// levels.
func ParseLevel(text string) (Level, error) {
	var level level
	err := level.UnmarshalText([]byte(text))
	return &level, err
}
