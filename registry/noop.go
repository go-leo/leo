package registry

import "context"

type Noop struct{}

func (n Noop) Register(ctx context.Context, service *ServiceInfo) error {
	return nil
}

func (n Noop) Deregister(ctx context.Context, service *ServiceInfo) error {
	return nil
}

func (n Noop) Scheme() string {
	return ""
}

func (n Noop) GetService(ctx context.Context, service *ServiceInfo) ([]*ServiceInfo, error) {
	return make([]*ServiceInfo, 0), nil
}

func (n Noop) Watch(ctx context.Context, service *ServiceInfo) (<-chan []*ServiceInfo, error) {
	return nil, nil
}

func (n Noop) StopWatch(ctx context.Context, service *ServiceInfo) error {
	return nil
}
