package healthx

import (
	"context"
	"google.golang.org/grpc/health/grpc_health_v1"
	"sync"
)

var (
	checkers sync.Map
)

func RegisterChecker(checker Checker) {
	checkers.Store(checker.Name(), checker)
}

func GetCheckers() map[string]Checker {
	m := make(map[string]Checker)
	checkers.Range(func(key, value any) bool {
		m[key.(string)] = value.(Checker)
		return true
	})
	return m
}

// Checker actual health check logic.
type Checker interface {
	Check(ctx context.Context) Status
	Name() string
}

// Status express state of a component or system.
type Status interface {
	// Code return the code of the Status.
	Code() int
	// Name return the name of the Status.
	Name() string
}

// CheckerProvider provides a Checker.
type CheckerProvider interface {
	HealthChecker() Checker
}

// status is a Status implementation.
type status grpc_health_v1.HealthCheckResponse_ServingStatus

// Code return the code of the Status.
func (s status) Code() int {
	return int(s)
}

// Name return the name of the Status.
func (s status) Name() string {
	return grpc_health_v1.HealthCheckResponse_ServingStatus_name[int32(s)]
}

var (
	Unknown = status(grpc_health_v1.HealthCheckResponse_UNKNOWN)

	Serving = status(grpc_health_v1.HealthCheckResponse_SERVING)

	NotServing = status(grpc_health_v1.HealthCheckResponse_NOT_SERVING)

	ServiceUnknown = status(grpc_health_v1.HealthCheckResponse_SERVICE_UNKNOWN)
)
