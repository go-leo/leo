package rpc

import (
	"context"

	"codeup.aliyun.com/qimao/leo/leo/registry"
)

// Provider is rpc service provider
type Provider interface {
	// Start 开始运行
	Start(ctx context.Context) error
	// Stop 停止运行
	Stop(ctx context.Context) error
	// Service 服务信息
	Service() registry.Service
	// Registrar 注册服务
	registry.Registrar
}
