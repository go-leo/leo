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
	if app.o.Logger != nil {
		app.o.Logger.Log("msg", `=================================`)
		app.o.Logger.Log("msg", `============  /\_/\  ============`)
		app.o.Logger.Log("msg", `============ (='.'=) ============`)
		app.o.Logger.Log("msg", `=================================`)
	}
	if app.o.ShutdownSignals != nil {
		var causeFunc context.CancelFunc
		ctx, causeFunc = signal.NotifyContext(ctx, app.o.ShutdownSignals...)
		defer causeFunc()
	}
	return runner.MultiRunner(app.o.Runners...).Run(ctx)
}
