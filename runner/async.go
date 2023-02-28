package runner

import (
	"context"
	"fmt"

	"github.com/go-leo/gox/syncx/brave"
	"github.com/go-leo/gox/syncx/chanx"
)

// asyncRunner 异步启动者
type asyncRunner struct {
	runner Runner
	errC   chan<- error
}

func (r *asyncRunner) Run(ctx context.Context) error {
	errC := brave.GoE(
		func() error { return r.runner.Run(ctx) },
		func(p any) error { return fmt.Errorf("%s", p) },
	)
	chanx.AsyncPipe(errC, r.errC)
	return nil
}

func AsyncRunner(runner Runner) (Runner, <-chan error) {
	errC := make(chan error)
	return &asyncRunner{runner: runner, errC: errC}, errC
}
