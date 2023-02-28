package runner

import (
	"context"

	"codeup.aliyun.com/qimao/leo/leo/registry"
)

type registrarStartStopper struct {
	Registrar registry.Registrar
	Service   registry.Service
}

func (r *registrarStartStopper) Start(ctx context.Context) error {
	return r.Registrar.Register(ctx, r.Service)
}

func (r *registrarStartStopper) Stop(ctx context.Context) error {
	return r.Registrar.Deregister(ctx, r.Service)
}

func RegistrarRunner(registrar registry.Registrar, service registry.Service) Runner {
	return StartStopRunner(&registrarStartStopper{Registrar: registrar, Service: service})
}
