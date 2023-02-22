package controller

import (
	"context"
)

// Server is view controller server
type Server interface {
	// Start 开始运行
	Start(ctx context.Context) error
	// Stop 停止运行
	Stop(ctx context.Context) error
}
