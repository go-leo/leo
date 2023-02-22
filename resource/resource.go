package resource

import (
	"context"
)

// Resource is restful server
type Resource interface {
	// Start 开始运行
	Start(ctx context.Context) error
	// Stop 停止运行
	Stop(ctx context.Context) error
}
