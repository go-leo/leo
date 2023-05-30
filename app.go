package leo

import (
	"context"
	"os"
	"os/signal"
	"syscall"
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
	ctx, causeFunc := signal.NotifyContext(ctx, app.o.ShutdownSignals...)
	defer causeFunc()
	return MutilRunner(app.o.Runners...).Run(ctx)
}
