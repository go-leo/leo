package leo

import (
	"context"
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
	app.o.Logger.Infof("leo app %d starting...", os.Getpid())
	defer app.o.Logger.Infof("leo app %d stopping...", os.Getpid())
	ctx, causeFunc := contextx.WithSignal(ctx, app.o.ShutdownSignals...)
	defer causeFunc(nil)

	var runners []runner.Runner
	if app.o.ActuatorServer != nil {
		runners = append(runners, app.o.ActuatorServer)
	}
	runners = append(runners, app.o.Runners...)
	return runner.MutilRunner(runners...).Run(ctx)
}
