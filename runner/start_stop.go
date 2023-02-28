package runner

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-leo/gox/syncx/brave"
)

type Starter interface {
	Start(ctx context.Context) error
}

type Stopper interface {
	Stop(ctx context.Context) error
}

type StartStopper interface {
	Starter
	Stopper
}

// startRunner 只有开始阶段的启动者
type startRunner struct {
	starter Starter
}

func (r *startRunner) Run(ctx context.Context) error {
	return r.starter.Start(ctx)
}

// startStopRunner 有开始也有结束的启动者
type startStopRunner struct {
	startStopper StartStopper
}

func (r *startStopRunner) Run(ctx context.Context) error {
	errC := brave.GoE(
		func() error { return r.startStopper.Start(ctx) },
		func(p any) error { return fmt.Errorf("%s", p) },
	)
	var err error
	select {
	case err = <-errC:
	case <-ctx.Done():
	}
	errC = brave.GoE(
		func() error { return r.startStopper.Stop(ctx) },
		func(p any) error { return fmt.Errorf("%s", p) },
	)
	return errors.Join(ctx.Err(), err, <-errC)
}

func StartRunner(starter Starter) Runner {
	return &startRunner{starter: starter}
}

func StartStopRunner(startStopper StartStopper) Runner {
	return &startStopRunner{startStopper: startStopper}
}
