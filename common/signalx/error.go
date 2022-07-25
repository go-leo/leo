package signalx

import (
	"fmt"
	"os"
)

type SignalError struct {
	Signal os.Signal
}

func (e SignalError) Error() string {
	return fmt.Sprintf("received [%s] signal", e.Signal)
}

func IsSignalError(err error) bool {
	_, ok := err.(*SignalError)
	return ok
}

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
