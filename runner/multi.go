package runner

import (
	"context"
	"fmt"
	"runtime"

	"github.com/go-leo/gox/syncx/brave"
	"golang.org/x/sync/errgroup"
)

type mutilRunner struct {
	runners []Runner
}

func (r *mutilRunner) Run(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, runner := range r.runners {
		eg.Go(doRun(ctx, runner))
		runtime.Gosched()
	}
	return eg.Wait()
}

func doRun(ctx context.Context, runner Runner) func() error {
	return func() error {
		return brave.DoE(
			func() error { return runner.Run(ctx) },
			func(p any) error { return fmt.Errorf("panic triggered: %+v", p) },
		)
	}
}

// MutilRunner 多个 Runner 合并成一个 mutilRunner, mutilRunner 使用 errgroup 并发的运行多个 Runner 并且阻塞。
// 如果其中一个 Runner 运行失败，则会返回该 error。
// 如果所有 Runner 都运行成功，则不会返回 error。
func MutilRunner(runners ...Runner) Runner {
	r := make([]Runner, len(runners))
	copy(r, runners)
	return &mutilRunner{runners: r}
}
