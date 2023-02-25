package resource

import (
	"context"

	"codeup.aliyun.com/qimao/leo/leo/registry"
)

// Server is restful server
type Server interface {
	// Start 开始运行
	Start(ctx context.Context) error
	// Stop 停止运行
	Stop(ctx context.Context) error
	// Service 服务信息
	Service() registry.Service
	// Registrar 注册服务
	registry.Registrar
	// Routes 路由
	Routes() []Route
}

type Route interface {
	Method() string
	Path() string
}
