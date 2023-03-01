package health

import "context"

// Checker actual health check logic.
type Checker interface {
	Check(ctx context.Context) Health
	Name() string
}
