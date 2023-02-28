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
	// Register 注册服务
	Register(ctx context.Context, registrar registry.Registrar) error
	// Deregister 注销服务
	Deregister(ctx context.Context, registrar registry.Registrar) error
	// Routes 路由
	Routes() []Route
}

type Route interface {
	Method() string
	Path() string
}
