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
	causeFunc := context.CancelFunc(func() {})
	if app.o.Signals != nil {
		ctx, causeFunc = signal.NotifyContext(ctx, app.o.Signals...)
	}
	defer causeFunc()

	runners := make([]runner.Runner, 0, len(app.o.Runners)+len(app.o.Servers))
	runners = append(runners, app.o.Runners...)

	for _, server := range app.o.Servers {
		runners = append(runners, runner.StartStopperRunner(server))
	}

	return runner.MultiRunner(runners...).Run(ctx)
}
