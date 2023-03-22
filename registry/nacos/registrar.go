package nacos

import (
	"context"

	"codeup.aliyun.com/qimao/leo/leo/registry"
)

type Registrar struct {
}

func (r *Registrar) Register(ctx context.Context, service registry.ServiceInstance) error {
	// TODO implement me
	panic("implement me")
}

func (r *Registrar) Deregister(ctx context.Context, service registry.ServiceInstance) error {
	// TODO implement me
	panic("implement me")
}
