package runner

import (
	"context"
	"errors"
	"github.com/go-leo/gox/syncx"
	"golang.org/x/sync/errgroup"
	"runtime"
)

type Starter interface {
	Start(ctx context.Context) error
}

// Stopper 结束
type Stopper interface {
	Stop(ctx context.Context) error
}

// StartStopper 有开始也有结束
type StartStopper interface {
	Starter
	Stopper
}

// startRunner 启动 Starter
type startRunner struct {
	starter Starter
}

// Run starter
func (r *startRunner) Run(ctx context.Context) error {
	return r.starter.Start(ctx)
}

// StartRunner 启动 Starter
func StartRunner(starter Starter) Runner {
	return &startRunner{starter: starter}
}

// startStopRunner 有开始也有结束的启动者
type startStopRunner struct {
	startStopper StartStopper
}

func (r *startStopRunner) Run(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error { return r.startStopper.Start(ctx) })
	runtime.Gosched()

	select {
	case err := <-syncx.WaitNotifyE(eg):
		return err
	case <-ctx.Done():
		err := r.startStopper.Stop(ctx)
		return errors.Join(ctx.Err(), err)
	}
}

// priorityStartStopRunner 主次关系的 StartStopper，
// 主要的 StartStopper 先运行，次要的 StartStopper 后运行，
// 次要的 StartStopper 先停止，主要的 StartStopper 后停止.
type priorityStartStopRunner struct {
	// dominant 主要的 StartStopper
	dominant StartStopper
	// subordinate 次要的 StartStopper
	subordinate StartStopper
}

func (r *priorityStartStopRunner) Run(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)
	// 主要的先运行
	eg.Go(func() error { return r.dominant.Start(ctx) })
	runtime.Gosched()

	// 次要的后运行
	eg.Go(func() error {
		runtime.Gosched()
		return r.subordinate.Start(ctx)
	})

	select {
	case err := <-syncx.WaitNotifyE(eg):
		return err
	case <-ctx.Done():
		// 次要的先停止
		subStopErr := r.subordinate.Stop(ctx)
		// 主要的后停止
		domStopErr := r.dominant.Stop(ctx)
		return errors.Join(ctx.Err(), subStopErr, domStopErr)
	}
}

// StartStopRunner 启动 StartStopper
func StartStopRunner(startStopper StartStopper) Runner {
	return &startStopRunner{startStopper: startStopper}
}

// PriorityStartStopRunner 启动两个有优先级关系的StartStopper
func PriorityStartStopRunner(dominant, subordinate StartStopper) Runner {
	return &priorityStartStopRunner{dominant: dominant, subordinate: subordinate}
}
