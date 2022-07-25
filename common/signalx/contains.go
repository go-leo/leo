package signalx

import "os"

func Contains(signals []os.Signal, signal os.Signal) bool {
	for _, s := range signals {
		if s.String() == signal.String() {
			return true
		}
	}
	return false
}
