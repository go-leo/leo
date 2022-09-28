package signalx

import "os"

// Deprecated: Do not use. use github.com/go-leo/osx/signalx instead.
func Contains(signals []os.Signal, signal os.Signal) bool {
	for _, s := range signals {
		if s.String() == signal.String() {
			return true
		}
	}
	return false
}
