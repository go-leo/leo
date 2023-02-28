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
		errC := make(chan error)
		asyncRunner := AsyncRunner(runner, errC)
		_ = asyncRunner.Run(ctx)
		errCs = append(errCs, errC)
	}
	errC := chanx.Combine(errCs...)
	var errs []error
	for e := range errC {
		errs = append(errs, e)
	}
	return errors.Join(errs...)
}

func MutilRunner(runners ...Runner) Runner {
	r := make([]Runner, len(runners))
	copy(r, runners)
	return &mutilRunner{runners: r}
}
