package runner

import (
	"context"
)

// Runner 运行器
type Runner interface {
	// Run 启动
	Run(ctx context.Context) error
}

type RunnerFunc func(ctx context.Context) error

func (f RunnerFunc) Run(ctx context.Context) error {
	return f(ctx)
}
