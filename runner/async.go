package runner

import (
	"context"
)

// asyncRunner 异步运行 Runner
type asyncRunner struct {
	runner Runner
	errC   chan error
}

func (r *asyncRunner) Run(ctx context.Context) error {
	go func() {
		if err := r.runner.Run(ctx); err != nil {
			r.errC <- err
		}
		close(r.errC)
	}()
	return nil
}

// AsyncRunner 异步运行 Runner
func AsyncRunner(runner Runner) (Runner, <-chan error) {
	asyncRunner := &asyncRunner{runner: runner, errC: make(chan error, 1)}
	return asyncRunner, asyncRunner.errC
}
