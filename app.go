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
	app.o.InfoLogger.Infof("leo app %d starting...", os.Getpid())
	defer app.o.InfoLogger.Infof("leo app %d stopping...", os.Getpid())
	ctx, causeFunc := contextx.WithSignal(ctx, app.o.ShutdownSignals...)
	defer causeFunc(nil)
	return runner.MutilRunner(app.o.Runners...).Run(ctx)
}
