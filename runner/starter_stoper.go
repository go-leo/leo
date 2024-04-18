package runner

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-leo/gox/syncx/brave"
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
	// 如果启动失败报错，error 发送到 startErrC
	// 如果启动成功，startErrC 将会被关闭
	startErrC := brave.GoE(
		func() error { return r.startStopper.Start(ctx) },
		func(p any) error { return fmt.Errorf("panic triggered: %+v", p) },
	)
	runtime.Gosched()
	select {
	case startErr := <-startErrC:
		return startErr
	case <-ctx.Done():
		stopErrC := brave.GoE(
			func() error { return r.startStopper.Stop(ctx) },
			func(p any) error { return fmt.Errorf("panic triggered: %+v", p) },
		)
		return errors.Join(ctx.Err(), <-stopErrC)
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
	dominantStartErrC, subordinateStartErrC := r.start(ctx)
	// 等待信号
	select {
	case err := <-dominantStartErrC:
		// 如果主要的执行失败，次要的要停止
		if err != nil {
			return errors.Join(err, <-r.stopSubordinate(ctx))
		}
		// 主要的执行完毕，等待次要
		return r.waitSubordinate(ctx, subordinateStartErrC)
	case err := <-subordinateStartErrC:
		// 如果次要的执行失败，主要的也要停止
		if err != nil {
			return errors.Join(err, <-r.stopDominant(ctx))
		}
		// 次要的执行完毕，等待次要主要
		return r.waitDominant(ctx, dominantStartErrC)
	case <-ctx.Done():
		subordinateStopErrC, dominantStopErrC := r.stop(ctx)
		return errors.Join(ctx.Err(), <-subordinateStopErrC, <-subordinateStartErrC, <-dominantStopErrC, <-dominantStartErrC)
	}
}

func (r *priorityStartStopRunner) waitDominant(ctx context.Context, errC <-chan error) error {
	select {
	case err := <-errC:
		return err
	case <-ctx.Done():
		return errors.Join(ctx.Err(), <-r.stopDominant(ctx))
	}
}

func (r *priorityStartStopRunner) waitSubordinate(ctx context.Context, errC <-chan error) error {
	select {
	case err := <-errC:
		return err
	case <-ctx.Done():
		return errors.Join(ctx.Err(), <-r.stopSubordinate(ctx))
	}
}

func (r *priorityStartStopRunner) start(ctx context.Context) (<-chan error, <-chan error) {
	// 主要的先运行
	dominantStartErrC := r.startDominant(ctx)
	// 次要的后运行
	runtime.Gosched()
	subordinateStartErrC := r.startSubordinate(ctx)
	return dominantStartErrC, subordinateStartErrC
}

func (r *priorityStartStopRunner) stop(ctx context.Context) (<-chan error, <-chan error) {
	// 次要的先停止
	subordinateStopErrC := r.stopSubordinate(ctx)
	// 主要的后停止
	runtime.Gosched()
	dominantStopErrC := r.stopDominant(ctx)
	return subordinateStopErrC, dominantStopErrC
}

func (r *priorityStartStopRunner) startDominant(ctx context.Context) <-chan error {
	return brave.GoE(
		func() error { return r.dominant.Start(ctx) },
		func(p any) error { return fmt.Errorf("panic triggered: %+v", p) },
	)
}

func (r *priorityStartStopRunner) startSubordinate(ctx context.Context) <-chan error {
	return brave.GoE(
		func() error {
			runtime.Gosched()
			return r.subordinate.Start(ctx)
		},
		func(p any) error { return fmt.Errorf("panic triggered: %+v", p) },
	)
}

func (r *priorityStartStopRunner) stopSubordinate(ctx context.Context) <-chan error {
	return brave.GoE(
		func() error { return r.subordinate.Stop(ctx) },
		func(p any) error { return fmt.Errorf("panic triggered: %+v", p) },
	)
}

func (r *priorityStartStopRunner) stopDominant(ctx context.Context) <-chan error {
	return brave.GoE(
		func() error {
			runtime.Gosched()
			return r.dominant.Stop(ctx)
		},
		func(p any) error { return fmt.Errorf("panic triggered: %+v", p) },
	)
}

// StartStopRunner 启动 StartStopper
func StartStopRunner(startStopper StartStopper) Runner {
	return &startStopRunner{startStopper: startStopper}
}

// PriorityStartStopRunner 启动两个有优先级关系的StartStopper
func PriorityStartStopRunner(dominant, subordinate StartStopper) Runner {
	return &priorityStartStopRunner{dominant: dominant, subordinate: subordinate}
}
