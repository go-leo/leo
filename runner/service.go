package runner

import (
	"context"

	"codeup.aliyun.com/qimao/leo/leo/registry"
)

type ServiceRegistry interface {
	// ServiceInstance 服务信息
	ServiceInstance() registry.ServiceInstance
	// Registrar 服务注册器
	Registrar() registry.Registrar
	// Register 注册服务
	Register(ctx context.Context) error
	// Deregister 注销服务
	Deregister(ctx context.Context) error
}

type Service interface {
	StartStopper
	ServiceRegistry
}

type serviceRegistryStartStopper struct {
	serviceRegistry ServiceRegistry
}

func (s *serviceRegistryStartStopper) Start(ctx context.Context) error {
	service := s.serviceRegistry.ServiceInstance()
	if service == nil {
		return nil
	}
	registrar := s.serviceRegistry.Registrar()
	if registrar == nil {
		return nil
	}
	return registrar.Register(ctx, service)
}

func (s *serviceRegistryStartStopper) Stop(ctx context.Context) error {
	service := s.serviceRegistry.ServiceInstance()
	if service == nil {
		return nil
	}
	registrar := s.serviceRegistry.Registrar()
	if registrar == nil {
		return nil
	}
	return registrar.Deregister(ctx, service)
}

func ServiceRunner(service Service) Runner {
	return MasterSlaveStartStopRunner(service, service)
}
