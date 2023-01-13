package runner

import (
	"context"
	"fmt"
)

// Callable 可被调用的
type Callable interface {
	// Stringer 描述
	fmt.Stringer
	// Invoke 执行
	Invoke(ctx context.Context) error
}
