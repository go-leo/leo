package runner

import (
	"context"
	"errors"
	"github.com/go-leo/gox/syncx"
	"golang.org/x/sync/errgroup"
	"runtime"
)

// Starter is Starter
type Starter interface {
	Start(ctx context.Context) error
}

// Stopper is Stopper
type Stopper interface {
	Stop(ctx context.Context) error
}

// StartStopper is StartStopper
type StartStopper interface {
	Starter
	Stopper
}

// startStopperRunner is startStopperRunner
type startStopperRunner struct {
	startStopper StartStopper
}

func (r *startStopperRunner) Run(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error { return r.startStopper.Start(ctx) })
	runtime.Gosched()
	select {
	case err := <-syncx.WaitNotifyE(eg):
		return err
	case <-ctx.Done():
		ctx = context.WithoutCancel(ctx)
		err := r.startStopper.Stop(ctx)
		return errors.Join(ctx.Err(), err)
	}
}

// priorityStartStopperRunner is priorityStartStopperRunner
type priorityStartStopperRunner struct {
	// dominant
	dominant StartStopper
	// subordinate
	subordinate StartStopper
}

func (r *priorityStartStopperRunner) Run(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)

	// start dominant first
	eg.Go(func() error { return r.dominant.Start(ctx) })
	runtime.Gosched()

	// then start subordinate
	eg.Go(func() error {
		runtime.Gosched()
		return r.subordinate.Start(ctx)
	})

	select {
	case err := <-syncx.WaitNotifyE(eg):
		return err
	case <-ctx.Done():
		ctx = context.WithoutCancel(ctx)
		// stop subordinate first
		subStopErr := r.subordinate.Stop(ctx)
		// then stop dominant
		domStopErr := r.dominant.Stop(ctx)
		return errors.Join(ctx.Err(), subStopErr, domStopErr)
	}
}

// StartRunner wrap Starter to Runner
func StartRunner(starter Starter) Runner {
	return RunnerFunc(func(ctx context.Context) error { return starter.Start(ctx) })
}

// ServerRunner wrap Server to Runner
func ServerRunner(server StartStopper) Runner {
	return &startStopperRunner{startStopper: server}
}

// PriorityServerRunner
// start dominant StartStopper first，then start subordinate StartStopper.
// stop subordinate StartStopper first，then stop dominant StartStopper.
func PriorityServerRunner(dominant, subordinate StartStopper) Runner {
	return &priorityStartStopperRunner{dominant: dominant, subordinate: subordinate}
}
