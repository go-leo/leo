package console

import (
	"context"
)

// Server is command line server app
type Server interface {
	// Start 开始运行
	Start(ctx context.Context) error
}
