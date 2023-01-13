package runner

import (
	"context"
	"fmt"
)

// Runnable 可运行的
type Runnable interface {
	// Stringer 描述
	fmt.Stringer
	// Start 开始运行，<-ctx.Done()可保证Start方法退出
	Start(ctx context.Context) error
	// Stop 停止运行，保证Start方法可以退出
	Stop(ctx context.Context) error
}
