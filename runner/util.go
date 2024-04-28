package runner

import (
	"context"
	"fmt"
)

func panicHandler(p any) error {
	return fmt.Errorf("panic triggered: %+v", p)
}

func start(ctx context.Context, starter Starter) func() error {
	return func() error { return starter.Start(ctx) }
}

func stop(ctx context.Context, stopper Stopper) func() error {
	return func() error { return stopper.Stop(ctx) }
}
