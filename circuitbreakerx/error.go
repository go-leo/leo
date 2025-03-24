package circuitbreakerx

import (
	"github.com/go-leo/leo/v3/statusx"
)

var (
	// ErrCircuitOpen is returned by the transport when the downstream is
	// unavailable due to a broken circuit.
	ErrCircuitOpen = statusx.Unavailable(
		statusx.Message("circuitbreakerx: circuit breaker is open"),
		statusx.Identifier("github.com/go-leo/leo/v3/circuitbreakerx.ErrCircuitOpen"),
	)

	errLoadBreaker = "circuitbreakerx: failed to load %s breaker"
)
