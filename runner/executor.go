package runner

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/go-leo/osx/execx"
	"github.com/go-leo/slicex"
	"github.com/go-leo/syncx"
	"github.com/go-leo/syncx/chanx"
	"github.com/hashicorp/go-multierror"

	"github.com/go-leo/leo/v2/log"
)

type options struct {
	Logger          log.Logger
	ShutdownSignals []os.Signal
	RestartSignals  []os.Signal
	StopTimeout     time.Duration
}

func (o *options) init() {
	if o.Logger == nil {
		o.Logger = log.Discard{}
	}
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

type Option func(o *options)

func Logger(l log.Logger) Option {
	return func(o *options) {
		o.Logger = l
	}
}

func ShutdownSignals(l []os.Signal) Option {
	return func(o *options) {
		o.ShutdownSignals = l
	}
}

func RestartSignals(l []os.Signal) Option {
	return func(o *options) {
		o.RestartSignals = l
	}
}

func StopTimeout(l time.Duration) Option {
	return func(o *options) {
		o.StopTimeout = l
	}
}

type Executor struct {
	callables []Callable
	runnables []Runnable
	o         *options
}

func NewExecutor(opts ...Option) *Executor {
	o := &options{
		Logger:          log.Discard{},
		ShutdownSignals: []os.Signal{syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT},
		RestartSignals:  []os.Signal{syscall.SIGHUP},
		StopTimeout:     5 * time.Second,
	}
	o.apply(opts...)
	o.init()
	return &Executor{o: o}
}

func (exe *Executor) AddRunnable(target Runnable) {
	exe.runnables = append(exe.runnables, target)
}

func (exe *Executor) AddCallable(target Callable) {
	exe.callables = append(exe.callables, target)
}

func (exe *Executor) Execute(ctx context.Context) error {
	exe.o.Logger.Infof("process %d starting...", os.Getpid())
	defer exe.o.Logger.Infof("process %d stopping...", os.Getpid())
	var runErrChans []<-chan error
	for _, callable := range exe.callables {
		runErrChans = append(runErrChans, exe.invoke(ctx, callable))
	}
	for _, runnable := range exe.runnables {
		runErrChans = append(runErrChans, exe.start(ctx, runnable))
	}
	runErrC := chanx.CombineChannels(runErrChans...)

	signalC := make(chan os.Signal)
	signals := slicex.Concat(exe.o.ShutdownSignals, exe.o.RestartSignals)
	if slicex.IsNotEmpty(signals) {
		signal.Notify(signalC, signals...)
	}

	var incomingSignal os.Signal
	var runErr error
	exe.o.Logger.Info("notify signals...")
	select {
	case incomingSignal = <-signalC:
		exe.o.Logger.Infof("receive signals %s", incomingSignal)
	case runErr = <-runErrC:
	}

	var stopErrChans []<-chan error
	for _, runnable := range exe.runnables {
		stopErrChans = append(stopErrChans, exe.stop(ctx, runnable))
	}
	stopErrC := chanx.CombineChannels(stopErrChans...)

	var err error
	if runErr != nil {
		err = multierror.Append(err, runErr)
	}
	for e := range runErrC {
		err = multierror.Append(err, e)
	}
	for e := range stopErrC {
		err = multierror.Append(err, e)
	}

	// 如果是运行错误，就直接退出
	if runErr != nil {
		return err
	}

	// 没有监听重启信号
	if slicex.IsNotEmpty(exe.o.RestartSignals) {
		return err
	}

	// 没有收到重启信号，需要重启
	f := func(o os.Signal) bool { return o.String() == incomingSignal.String() }
	if slicex.NotContainsFunc(exe.o.RestartSignals, f) {
		return err
	}

	// 重启
	exe.o.Logger.Infof("restart process success")
	if _, e := execx.StartProcess(); e != nil {
		exe.o.Logger.Errorf("failed to restart process, %v", e)
		err = multierror.Append(err, e)
	}
	return err
}

func (exe *Executor) invoke(ctx context.Context, target Callable) <-chan error {
	// 并发启动
	f := func() error {
		defer exe.o.Logger.Infof("%s is exited", target.String())
		exe.o.Logger.Infof("calling %s", target.String())
		err := target.Invoke(ctx)
		if err != nil {
			err = fmt.Errorf("failed to invoke %s, %w", target.String(), err)
		}
		return err
	}
	r := func(p any) {
		exe.o.Logger.Infof("panic trigger when calling %s", target.String())
	}
	errC := syncx.BraveGoE(f, r)
	runtime.Gosched()
	return errC
}

func (exe *Executor) start(ctx context.Context, target Runnable) <-chan error {
	// 并发启动
	f := func() error {
		exe.o.Logger.Infof("starting %s", target.String())
		err := target.Start(ctx)
		if err != nil {
			err = fmt.Errorf("failed to start %s, %w", target.String(), err)
		}
		return err
	}
	r := func(p any) {
		exe.o.Logger.Infof("panic trigger when starting %s", target.String())
	}
	errC := syncx.BraveGoE(f, r)
	runtime.Gosched()
	return errC
}

func (exe *Executor) stop(ctx context.Context, target Runnable) <-chan error {
	f := func() error {
		defer exe.o.Logger.Infof("%s is stopped", target.String())
		ctx, _ := context.WithTimeout(context.Background(), exe.o.StopTimeout)
		err := target.Stop(ctx)
		if err != nil {
			err = fmt.Errorf("failed to stop %s, %w", target.String(), err)
		}
		return err
	}
	r := func(p any) {
		exe.o.Logger.Infof("panic trigger when stopping %s", target.String())
	}
	errC := syncx.BraveGoE(f, r)
	runtime.Gosched()
	return errC
}
