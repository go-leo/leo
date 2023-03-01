package view

import (
	"context"
)

// Controller is view controller server
type Controller interface {
	// Start 开始运行
	Start(ctx context.Context) error
	// Stop 停止运行
	Stop(ctx context.Context) error
	// Views 视图
	Views() []View
}

type View interface {
	Method() string
	Path() string
	View() string
}
