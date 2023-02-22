package runner

import (
	"context"
	"fmt"

	"github.com/go-leo/gox/syncx/brave"
)

// AsyncRunner 异步启动者
type AsyncRunner interface {
	AsyncRun(ctx context.Context) <-chan error
}

type asyncRunner struct {
	runner Runner
}

func (r *asyncRunner) AsyncRun(ctx context.Context) <-chan error {
	return brave.GoE(
		func() error { return r.runner.Run(ctx) },
		func(p any) error { return fmt.Errorf("%s", p) },
	)
}

func NewAsyncRunner(runner Runner) AsyncRunner {
	return &asyncRunner{runner: runner}
}
