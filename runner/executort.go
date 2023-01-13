package runner

//
//import (
//	"context"
//	"os"
//	"runtime"
//	"sync"
//	"syscall"
//	"time"
//
//	"github.com/go-leo/osx/execx"
//	"github.com/go-leo/osx/signalx"
//	"github.com/go-leo/slicex"
//	"github.com/go-leo/syncx"
//	"golang.org/x/sync/errgroup"
//
//	"github.com/go-leo/leo/v2/log"
//)
//
//type options struct {
//	Logger          log.Logger
//	ShutdownSignals []os.Signal
//	ShutdownHook    func(signal os.Signal)
//	RestartSignals  []os.Signal
//	RestartHook     func(signal os.Signal)
//	StopTimeout     time.Duration
//}
//
//type Option func(e *options)
//
//func (o *options) init() {
//	if o.Logger == nil {
//		o.Logger = log.Discard{}
//	}
//	if o.ShutdownHook == nil {
//		o.ShutdownHook = func(_ os.Signal) {}
//	}
//	if o.RestartHook == nil {
//		o.RestartHook = func(_ os.Signal) {}
//	}
//	if o.StopTimeout == 0 {
//		o.StopTimeout = 5 * time.Second
//	}
//}
//
//func (o *options) apply(opts ...Option) {
//	for _, opt := range opts {
//		opt(o)
//	}
//}
//
//type Executor struct {
//	errGroup  *errgroup.Group
//	waitGroup *sync.WaitGroup
//	cancel    context.CancelFunc
//	callables []Callable
//	runnables []Runnable
//	o         *options
//}
//
//func NewExecutor(opts ...Option) *Executor {
//	o := &options{
//		ShutdownSignals: []os.Signal{syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT},
//		RestartSignals:  []os.Signal{syscall.SIGHUP},
//		StopTimeout:     5 * time.Second,
//	}
//	o.apply(opts...)
//	o.init()
//	return &Executor{o: o}
//}
//
//func (exe *Executor) AddRunnable(target Runnable) {
//	exe.runnables = append(exe.runnables, target)
//}
//
//func (exe *Executor) AddCallable(target Callable) {
//	exe.callables = append(exe.callables, target)
//}
//
//func (exe *Executor) Execute(ctx context.Context) error {
//	exe.waitGroup = new(sync.WaitGroup)
//	ctx, exe.cancel = context.WithCancel(ctx)
//	exe.errGroup, ctx = errgroup.WithContext(ctx)
//	for _, callable := range exe.callables {
//		exe.invoke(ctx, callable)
//	}
//	for _, runnable := range exe.runnables {
//		exe.start(ctx, runnable)
//	}
//	// 等待系统信号，支持关闭和重启信号
//	if len(exe.o.ShutdownSignals)+len(exe.o.RestartSignals) > 0 {
//		exe.listenSignal(ctx)
//	}
//	// 等待退出
//	return exe.wait()
//}
//
//func (exe *Executor) invoke(ctx context.Context, target Callable) {
//	// 并发启动
//	exe.waitGroup.Add(1)
//	exe.errGroup.Go(func() (err error) {
//		f := func() error {
//			defer exe.o.Logger.Infof("%s is exited", target.String())
//			defer exe.waitGroup.Done()
//			exe.o.Logger.Infof("calling %s", target.String())
//			return target.Invoke(ctx)
//		}
//		r := func(p any) {
//			exe.o.Logger.Infof("panic trigger when calling %s", target.String())
//		}
//		return syncx.BraveDoE(f, r)
//	})
//	runtime.Gosched()
//}
//
//func (exe *Executor) start(ctx context.Context, target Runnable) {
//	// 并发启动
//	exe.waitGroup.Add(1)
//	exe.errGroup.Go(func() error {
//		f := func() error {
//			defer exe.waitGroup.Done()
//			exe.o.Logger.Infof("starting %s", target.String())
//			return target.Start(ctx)
//		}
//		r := func(p any) {
//			exe.o.Logger.Infof("panic trigger when starting %s", target.String())
//		}
//		return syncx.BraveDoE(f, r)
//	})
//	runtime.Gosched()
//	// 监听停止信号
//	exe.errGroup.Go(func() error {
//		<-ctx.Done()
//		f := func() error {
//			defer exe.o.Logger.Infof("%s is stopped", target.String())
//			ctx, _ := context.WithTimeout(context.Background(), exe.o.StopTimeout)
//			return target.Stop(ctx)
//		}
//		r := func(p any) {
//			exe.o.Logger.Infof("panic trigger when stopping %s", target.String())
//		}
//		return syncx.BraveDoE(f, r)
//	})
//	runtime.Gosched()
//}
//
//func (exe *Executor) listenSignal(ctx context.Context) {
//	exe.errGroup.Go(func() error {
//		signals := slicex.Concat(exe.o.ShutdownSignals, exe.o.RestartSignals)
//		hooks := []func(os.Signal){exe.o.ShutdownHook, exe.o.RestartHook}
//		errC := make(chan error)
//		go func() {
//			runtime.Gosched()
//			exe.o.Logger.Info("wait signals...")
//			errC <- signalx.NewSignalWaiter(signals, exe.o.StopTimeout, hooks...).Wait().Err()
//			close(errC)
//		}()
//		select {
//		case <-ctx.Done():
//			return ctx.Err()
//		case e := <-errC:
//			return e
//		}
//	})
//	runtime.Gosched()
//}
//
//func (exe *Executor) wait() error {
//	exe.errGroup.Go(func() error {
//		exe.waitGroup.Wait()
//		exe.cancel()
//		return nil
//	})
//	err := exe.errGroup.Wait()
//	if err == nil {
//		return nil
//	}
//	if !signalx.IsSignal(err, exe.o.RestartSignals) {
//		return err
//	}
//	if _, e := execx.StartProcess(); e != nil {
//		exe.o.Logger.Errorf("failed to restart process, %v", e)
//		return err
//	}
//	exe.o.Logger.Infof("restart process success")
//	return err
//}
