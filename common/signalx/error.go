package signalx

import (
	"fmt"
	"os"
)

// Deprecated: Do not use. use github.com/go-leo/osx/signalx instead.
type SignalError struct {
	Signal os.Signal
}

// Deprecated: Do not use. use github.com/go-leo/osx/signalx instead.
func (e SignalError) Error() string {
	return fmt.Sprintf("received [%s] signal", e.Signal)
}

// Deprecated: Do not use. use github.com/go-leo/osx/signalx instead.
func IsSignalError(err error) bool {
	_, ok := err.(*SignalError)
	return ok
}

// Deprecated: Do not use. use github.com/go-leo/osx/signalx instead.
func IsSignal(err error, signals []os.Signal) bool {
	signalErr, ok := err.(*SignalError)
	if ok && Contains(signals, signalErr.Signal) {
		return true
	}
	signalError, ok := err.(SignalError)
	if ok && Contains(signals, signalError.Signal) {
		return true
	}
	return false
}
