package circuitbreakerx

import (
	"github.com/go-leo/status"
	"google.golang.org/grpc/codes"
)

var (
	// ErrCircuitOpen is returned by the transport when the downstream is
	// unavailable due to a broken circuit.
	ErrCircuitOpen = status.New(codes.Unavailable,
		status.Message("circuitbreakerx: circuit breaker is open"),
		status.Identifier("github.com/go-leo/leo/v3/circuitbreakerx.ErrCircuitOpen"),
	)

	errLoadBreaker = "circuitbreakerx: failed to load %s breaker"
)
