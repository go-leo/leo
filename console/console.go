package console

import (
	"context"
)

// Console is command line
type Console interface {
	// Start 开始运行
	Start(ctx context.Context) error
}
