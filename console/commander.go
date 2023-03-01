package console

import (
	"context"
)

// Commander is command line app
type Commander interface {
	// Start 开始运行
	Start(ctx context.Context) error

	CommandLine() string
}
