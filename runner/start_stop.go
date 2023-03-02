package runner

import (
	"context"
	"errors"
	"fmt"
	"runtime"

	"github.com/go-leo/gox/syncx/brave"
	"github.com/go-leo/gox/syncx/chanx"
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

// StartRunner 启动 Starter
func StartRunner(starter Starter) Runner {
	return &startRunner{starter: starter}
}

// StartStopRunner 启动 StartStopper
func StartStopRunner(startStopper StartStopper) Runner {
	return &startStopRunner{startStopper: startStopper}
}

// MasterSlaveStartStopRunner 启动两个有主从关系的 StartStopper
func MasterSlaveStartStopRunner(master, slave StartStopper) Runner {
	return &masterSlaveStartStopRunner{master: master, slave: slave}
}
