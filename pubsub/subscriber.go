package pubsub

import (
	"context"
)

// Subscriber is message queue subscriber
type Subscriber interface {
	// Start 开始运行
	Start(ctx context.Context) error
	// Stop 停止运行
	Stop(ctx context.Context) error
}
