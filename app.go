package leo

import (
	"context"
	"github.com/go-leo/gox/contextx"
	"github.com/go-leo/leo/v3/runner"
	"os"
	"syscall"
)

type Option func(o *App)

// Runner 添加runner
func Runner(runners ...runner.Runner) Option {
	return func(o *App) {
		o.runners = append(o.runners, runners...)
	}
}

// Signal 监听信号
func Signal(signals ...os.Signal) Option {
	return func(o *App) {
		o.signals = signals
	}
}

type App struct {
	runners []runner.Runner
	signals []os.Signal
}

func (app *App) apply(opts ...Option) *App {
	for _, opt := range opts {
		opt(app)
	}
	return app
}

func (app *App) complete() *App {
	if app.signals == nil {
		// 默认监听中断信号
		app.signals = []os.Signal{os.Interrupt, syscall.SIGTERM}
	}
	return app
}

// Run 启动app
func (app *App) Run(ctx context.Context) error {
	ctx, causeFunc := contextx.WithSignalCause(ctx, app.signals...)
	defer causeFunc(nil)
	return runner.MultiRunner(app.runners...).Run(ctx)
}

func NewApp(opts ...Option) *App {
	return new(App).apply(opts...).complete()
}
