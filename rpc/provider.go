package rpc

import (
	"context"
)

// Provider is rpc service provider
type Provider interface {
	// Start 开始运行
	Start(ctx context.Context) error
	// Stop 停止运行
	Stop(ctx context.Context) error
}
