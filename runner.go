package leo

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-leo/gox/syncx/chanx"
	"runtime"

	"github.com/go-leo/gox/syncx/brave"
	"golang.org/x/sync/errgroup"
)

// Runner 启动者
type Runner interface {
	// Run 启动
	Run(ctx context.Context) error
}

type mutilRunner struct {
	runners []Runner
}

func (r *mutilRunner) Run(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, runner := range r.runners {
		eg.Go(doRun(ctx, runner))
		runtime.Gosched()
	}
	return eg.Wait()
}

func doRun(ctx context.Context, runner Runner) func() error {
	return func() error {
		return brave.DoE(
			func() error { return runner.Run(ctx) },
			func(p any) error { return fmt.Errorf("panic triggered: %+v", p) },
		)
	}
}

// MutilRunner 多个 Runner 合并成一个 mutilRunner, mutilRunner 使用 errgroup 并发的运行多个 Runner 并且阻塞。
// 如果其中一个 Runner 运行失败，则会返回该 error。
// 如果所有 Runner 都运行成功，则不会返回 error。
func MutilRunner(runners ...Runner) Runner {
	r := make([]Runner, len(runners))
	copy(r, runners)
	return &mutilRunner{runners: r}
}

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

type Starter interface {
	Start(ctx context.Context) error
}

// startRunner 只有开始阶段的启动者
type startRunner struct {
	starter Starter
}

func (r *startRunner) Run(ctx context.Context) error {
	return r.starter.Start(ctx)
}

// StartRunner 启动 Starter
func StartRunner(starter Starter) Runner {
	return &startRunner{starter: starter}
}

type Stopper interface {
	Stop(ctx context.Context) error
}

type StartStopper interface {
	Starter
	Stopper
}

// startStopRunner 有开始也有结束的启动者
type startStopRunner struct {
	startStopper StartStopper
}

func (r *startStopRunner) Run(ctx context.Context) error {
	return runStartStopper(ctx, r)
}

func runStartStopper(ctx context.Context, r *startStopRunner) error {
	startErrC := brave.GoE(
		func() error { return r.startStopper.Start(ctx) },
		func(p any) error { return fmt.Errorf("%v", p) },
	)
	runtime.Gosched()
	select {
	case startErr := <-startErrC:
		return startErr
	case <-ctx.Done():
		stopErrC := brave.GoE(
			func() error { return r.startStopper.Stop(ctx) },
			func(p any) error { return fmt.Errorf("%v", p) },
		)
		return errors.Join(ctx.Err(), <-stopErrC)
	}
}

// StartStopRunner 启动 StartStopper
func StartStopRunner(startStopper StartStopper) Runner {
	return &startStopRunner{startStopper: startStopper}
}

// masterSlaveStartStopRunner 两个有主从关系的 StartStopper，主先运行，从后运行，从先停止，主后停止。
type masterSlaveStartStopRunner struct {
	master StartStopper
	slave  StartStopper
}

func (r *masterSlaveStartStopRunner) Run(ctx context.Context) error {
	// master先运行
	masterStartErrC := brave.GoE(
		func() error { return r.master.Start(ctx) },
		func(p any) error { return fmt.Errorf("%s", p) },
	)

	// slave后运行
	runtime.Gosched()
	slaveStartErrC := brave.GoE(
		func() error {
			runtime.Gosched()
			return r.slave.Start(ctx)
		},
		func(p any) error { return fmt.Errorf("%s", p) },
	)

	// wait
	var err error
	select {
	case err = <-chanx.Combine(masterStartErrC, slaveStartErrC):
	case <-ctx.Done():
	}

	// slave先停止
	slaveStopErrC := brave.GoE(
		func() error { return r.slave.Stop(ctx) },
		func(p any) error { return fmt.Errorf("%s", p) },
	)

	// master后停止
	runtime.Gosched()
	masterStopErrC := brave.GoE(
		func() error {
			runtime.Gosched()
			return r.master.Stop(ctx)
		},
		func(p any) error { return fmt.Errorf("%s", p) },
	)

	return errors.Join(ctx.Err(), err, <-slaveStopErrC, <-masterStopErrC)
}

// MasterSlaveStartStopRunner 启动两个有主从关系的 StartStopper
func MasterSlaveStartStopRunner(master, slave StartStopper) Runner {
	return &masterSlaveStartStopRunner{master: master, slave: slave}
}
