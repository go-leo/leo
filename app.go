package leo

import (
	"context"
	"github.com/go-leo/leo/v3/runner"
	"os/signal"
)

type App struct {
	o *options
}

func NewApp(opts ...Option) *App {
	o := new(options).apply(opts...).init()
	return &App{o: o}
}

// Run 启动app
func (app *App) Run(ctx context.Context) error {
	if app.o.Signals == nil {
		return runner.MultiRunner(app.o.Runners...).Run(ctx)
	}
	var causeFunc context.CancelFunc
	ctx, causeFunc = signal.NotifyContext(ctx, app.o.Signals...)
	defer causeFunc()
	return runner.MultiRunner(app.o.Runners...).Run(ctx)
}
