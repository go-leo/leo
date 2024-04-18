package runner

import (
	"context"
	"fmt"
	"github.com/go-leo/gox/syncx/brave"
	"golang.org/x/sync/errgroup"
	"runtime"
)

// multiRunner 多个 Runner 合并成一个 multiRunner,
// multiRunner 使用 errgroup 并发的运行多个 Runner 并且阻塞。
type multiRunner struct {
	runners []Runner
}

// Run 启动
func (r *multiRunner) Run(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, runner := range r.runners {
		eg.Go(doRun(ctx, runner))
		runtime.Gosched()
	}
	return eg.Wait()
}

// doRun 运行 Runner
func doRun(ctx context.Context, runner Runner) func() error {
	return func() error {
		return brave.DoE(
			func() error { return runner.Run(ctx) },
			func(p any) error { return fmt.Errorf("panic triggered: %+v", p) },
		)
	}
}

// MultiRunner 多个 Runner 合并成一个 multiRunner,
// multiRunner 使用 errgroup 并发的运行多个 Runner 并且阻塞。
// 如果其中一个 Runner 运行失败，则会返回该 error。
// 如果所有 Runner 都运行成功，则不会返回 error。
func MultiRunner(runners ...Runner) Runner {
	allRunners := make([]Runner, 0, len(runners))
	for _, r := range runners {
		if mr, ok := r.(*multiRunner); ok {
			allRunners = append(allRunners, mr.runners...)
			continue
		}
		allRunners = append(allRunners, r)
	}
	return &multiRunner{runners: allRunners}
}
