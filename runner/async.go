package runner

import (
	"context"
	"fmt"
	"github.com/go-leo/gox/syncx/brave"
	"github.com/go-leo/gox/syncx/chanx"
)

// asyncRunner 异步运行 Runner
type asyncRunner struct {
	runner Runner
	errC   chan<- error
}

func (r *asyncRunner) Run(ctx context.Context) error {
	errC := brave.GoE(
		func() error { return r.runner.Run(ctx) },
		func(p any) error { return fmt.Errorf("panic triggered: %+v", p) },
	)
	chanx.AsyncPipe(errC, r.errC)
	return nil
}

// AsyncRunner 异步运行 Runner
func AsyncRunner(runner Runner) (Runner, <-chan error) {
	errC := make(chan error)
	return &asyncRunner{runner: runner, errC: errC}, errC
}
