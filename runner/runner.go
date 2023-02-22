package runner

import (
	"context"
)

// Runner 启动者
type Runner interface {
	// Run 启动
	Run(ctx context.Context) error
}
