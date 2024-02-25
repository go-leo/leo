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
	ctx, causeFunc := signal.NotifyContext(ctx, app.o.ShutdownSignals...)
	defer causeFunc()
	var runners []Runner
	runners = append(runners, app.o.Runners...)
	return MutilRunner(runners...).Run(ctx)
}
