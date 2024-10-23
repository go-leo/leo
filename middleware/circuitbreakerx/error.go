package circuitbreakerx

import "errors"

var (
	// ErrCircuitOpen is returned by the transport when the downstream is
	// unavailable due to a broken circuit.
	ErrCircuitOpen = errors.New("circuit open")
)
