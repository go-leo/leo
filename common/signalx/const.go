package signalx

import (
	"os"
	"syscall"
)

// Deprecated: Do not use. use github.com/go-leo/osx/signalx instead.
func ShutdownSignals() []os.Signal {
	return []os.Signal{
		syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL,
	}
}
