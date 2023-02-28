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

	for _, commander := range app.o.Commanders {
		runners = append(runners, runner.StartRunner(commander))
	}

	for _, server := range app.o.Controllers {
		runners = append(runners, runner.StartStopRunner(server))
	}

	for _, server := range app.o.Resources {
		runners = append(runners, runner.StartStopRunner(server))
	}

	for _, router := range app.o.Routers {
		runners = append(runners, runner.StartStopRunner(router))
	}

	for _, router := range app.o.Routers {
		runners = append(runners, runner.StartStopRunner(router))
	}

	for _, provider := range app.o.Providers {
		runners = append(runners, runner.StartStopRunner(provider))
	}

	for _, scheduler := range app.o.Schedulers {
		runners = append(runners, runner.StartStopRunner(scheduler))
	}

	runners = append(runners, app.o.Runners...)
	err := runner.MutilRunner(runners...).Run(ctx)
	ctxErr := contextx.Error(ctx)
	return errors.Join(ctxErr, err)
}
