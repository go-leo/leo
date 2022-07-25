package passthrough

import (
	"context"
	"net/url"
	"sync"

	"github.com/go-leo/leo/registry"
)

var _ registry.Registrar = new(Registrar)

var _ registry.Discovery = new(Discovery)

const Scheme = "passthrough"

type Registrar struct{}

func NewRegistrar() *Registrar {
	return &Registrar{}
}

func (r *Registrar) Scheme() string {
	return Scheme
}

func (r *Registrar) Register(ctx context.Context, service *registry.ServiceInfo) error {
	return nil
}

func (r *Registrar) Deregister(ctx context.Context, service *registry.ServiceInfo) error {
	return nil
}

type Discovery struct {
	wg *sync.WaitGroup
}

func NewDiscovery() *Discovery {
	d := &Discovery{}
	return d
}

func (d *Discovery) Scheme() string {
	return Scheme
}

func (d *Discovery) GetService(ctx context.Context, service *registry.ServiceInfo) ([]*registry.ServiceInfo, error) {
	return []*registry.ServiceInfo{service.Clone()}, nil
}

func (d *Discovery) Watch(ctx context.Context, service *registry.ServiceInfo) (<-chan []*registry.ServiceInfo, error) {
	c := make(chan []*registry.ServiceInfo)
	d.wg = &sync.WaitGroup{}
	d.wg.Add(1)
	go func() {
		c <- []*registry.ServiceInfo{service.Clone()}
		d.wg.Wait()
	}()
	return c, nil
}

func (d *Discovery) StopWatch(ctx context.Context, service *registry.ServiceInfo) error {
	if d.wg != nil {
		d.wg.Done()
	}
	return nil
}

type RegistrarFactory struct{}

func (factory *RegistrarFactory) Create() (registry.Registrar, error) {
	return NewRegistrar(), nil
}

type DiscoveryFactory struct {
	URI *url.URL
}

func (factory *DiscoveryFactory) Create() (registry.Discovery, error) {
	return NewDiscovery(), nil
}
