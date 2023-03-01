package leo

import (
	"context"
	"errors"
	"os"
	"syscall"

	"github.com/go-leo/gox/contextx"

	"codeup.aliyun.com/qimao/leo/leo/runner"
)

type App struct {
	o *options
}

func NewApp(opts ...Option) *App {
	o := &options{
		ShutdownSignals: []os.Signal{syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT},
	}
	o.apply(opts...)
	o.init()
	return &App{o: o}
}

// Run 启动app
func (app *App) Run(ctx context.Context) error {
	app.o.InfoLogger.Infof("leo app %d starting...", os.Getpid())
	defer app.o.InfoLogger.Infof("leo app %d stopping...", os.Getpid())
	ctx, causeFunc := contextx.WithSignal(ctx, app.o.ShutdownSignals...)
	defer causeFunc(nil)

	var runners []runner.Runner

	if app.o.ConsoleCommander != nil {
		runners = append(runners, runner.StartRunner(app.o.ConsoleCommander))
	}

	if app.o.ViewController != nil {
		runners = append(runners, runner.StartStopRunner(app.o.ViewController))
	}

	if app.o.ResourceServer != nil {
		runners = append(runners, runner.StartStopRunner(app.o.ResourceServer))
	}

	if app.o.SteamRouter != nil {
		runners = append(runners, runner.StartStopRunner(app.o.SteamRouter))
	}

	if app.o.RPCProvider != nil {
		runners = append(runners, runner.StartStopRunner(app.o.RPCProvider))
	}

	if app.o.Scheduler != nil {
		runners = append(runners, runner.StartStopRunner(app.o.Scheduler))
	}

	if app.o.ActuatorServer != nil {
		runners = append(runners, runner.StartStopRunner(app.o.ActuatorServer))
	}

	mutilRunner := runner.MutilRunner(append(runners, app.o.Runners...)...)
	return errors.Join(mutilRunner.Run(ctx), contextx.Error(ctx))
}
