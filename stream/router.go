package stream

import "context"

type Router interface {
	// Start 开始运行
	Start(ctx context.Context) error
	// Stop 停止运行
	Stop(ctx context.Context) error
	// Bridges 所有Bridges
	Bridges() []Bridge
}
