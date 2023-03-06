package runner

import (
	"context"
	"errors"

	"github.com/go-leo/gox/syncx/chanx"
)

// mutilRunner 多启动者，多个合成一个
type mutilRunner struct {
	runners []Runner
}

func (r *mutilRunner) Run(ctx context.Context) error {
	var errCs []<-chan error
	for _, runner := range r.runners {
		asyncRunner, errC := AsyncRunner(runner)
		_ = asyncRunner.Run(ctx)
		errCs = append(errCs, errC)
	}
	errC := chanx.Combine(errCs...)
	errs := chanx.ReceiveUtilClosed(errC)
	return errors.Join(errs...)
}

func MutilRunner(runners ...Runner) Runner {
	r := make([]Runner, len(runners))
	copy(r, runners)
	return &mutilRunner{runners: r}
}
