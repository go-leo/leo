package healthx

import (
	"context"
	"github.com/go-leo/gox/mapx"
	"google.golang.org/grpc/health"
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
	return mapx.FromRanger[map[string]Checker](&checkers)
}

func GetChecker(name string) (Checker, bool) {
	value, ok := checkers.Load(name)
	if !ok {
		return nil, false
	}
	return value.(Checker), true
}

type Checker interface {
	// Name returns the name of the Checker.
	Name() string

	// HealthServer implements `service Health`.
	grpc_health_v1.HealthServer

	// Shutdown sets all serving status to NOT_SERVING, and configures the server to
	// ignore all future status changes.
	//
	// This changes serving status for all services. To set status for a particular
	// services, call SetServingStatus().
	Shutdown()

	// Resume sets all serving status to SERVING, and configures the server to
	// accept all future status changes.
	//
	// This changes serving status for all services. To set status for a particular
	// services, call SetServingStatus().
	Resume()
}

// CheckerProvider provides a Checker.
type CheckerProvider interface {
	HealthChecker(ctx context.Context) Checker
}

type checker struct {
	*health.Server
	name string
}

func (c *checker) Name() string {
	return c.name
}

func NewChecker(name string) Checker {
	return &checker{
		Server: health.NewServer(),
		name:   name,
	}
}
