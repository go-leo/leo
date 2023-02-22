package leo

import (
	"context"
	"fmt"
	"github.com/go-kit/log"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	o *options
}

func NewApp(opts ...Option) *App {
	o := &options{
		Runners:         []Runner{},
		ShutdownSignals: []os.Signal{syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT},
		Logger:          log.NewLogfmtLogger(os.Stdout),
	}
	return &App{o: o.apply(opts...).init()}
}

// Run 启动app
func (app *App) Run(ctx context.Context) error {
	fmt.Println("=================================")
	fmt.Println(`============  /\_/\  ============`)
	fmt.Println(`============ (='.'=) ============`)
	fmt.Println("=================================")
	ctx, causeFunc := signal.NotifyContext(ctx, app.o.ShutdownSignals...)
	defer causeFunc()
	var runners []Runner
	runners = append(runners, app.o.Runners...)
	return MutilRunner(runners...).Run(ctx)
}
