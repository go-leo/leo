package nacos

import (
	"context"
	"errors"
	"net/url"
)

// ErrServiceNotFound not found service.
var ErrServiceNotFound = errors.New("service not found")

// Discovery is service Discovery
type Discovery interface {
	// GetInstances gets all service instances associated with a particular service name.
	GetInstances(ctx context.Context, serviceName string) ([]ServiceInstance, error)
	// GetServices return all known service names.
	GetServices(ctx context.Context) ([]string, error)
	// Watch whe service instances associated with a particular service name.
	Watch(ctx context.Context, serviceName string) (<-chan []ServiceInstance, error)
	// Stop watch service instances
	Stop(ctx context.Context, serviceName string, watcher <-chan []ServiceInstance) error
}

type DiscoveryFactory interface {
	Create(uri *url.URL) (Discovery, error)
}

//
// var _ Discovery = new(BaseDiscovery)
//
// type BaseDiscovery struct {
// 	instanceCache *ServiceInstanceCache
// 	watchers      map[string][]chan []ServiceInstance
// 	sync.RWMutex
// }
//
// func (d *BaseDiscovery) GetInstances(ctx context.Context, serviceName string) ([]ServiceInstance, error) {
// 	instances, ok := d.instanceCache.GetInstances(serviceName)
// 	if !ok {
// 		return nil, ErrServiceNotFound
// 	}
// 	return append([]ServiceInstance{}, instances...), nil
// }
//
// func (d *BaseDiscovery) GetServices(ctx context.Context) ([]string, error) {
// 	return d.instanceCache.GetServices(), nil
// }
//
// func (d *BaseDiscovery) Watch(ctx context.Context, serviceName string) (<-chan []ServiceInstance, error) {
// 	d.Lock()
// 	defer d.Unlock()
// 	_, ok := d.instances[serviceName]
// 	if !ok {
// 		return nil, ErrServiceNotFound
// 	}
// 	ch := make(chan []ServiceInstance)
// 	d.watchers[serviceName] = append(d.watchers[serviceName], ch)
// 	return ch, nil
// }
//
// func (d *BaseDiscovery) Stop(ctx context.Context, serviceName string, watcher <-chan []ServiceInstance) error {
// 	d.Lock()
// 	defer d.Unlock()
// 	_, ok := d.instances[serviceName]
// 	if !ok {
// 		return ErrServiceNotFound
// 	}
// 	watchers, ok := d.watchers[serviceName]
// 	if !ok {
// 		return nil
// 	}
// 	var newWatchers []chan []ServiceInstance
// 	for _, w := range watchers {
// 		if w != watcher {
// 			newWatchers = append(newWatchers, w)
// 			continue
// 		}
// 		close(w)
// 	}
// 	d.watchers[serviceName] = newWatchers
// 	return nil
// }
//
// func (d *BaseDiscovery) onSetInstances(serviceName string, instances []ServiceInstance) {
// 	d.Lock()
// 	defer d.Unlock()
// 	watchers, ok := d.watchers[serviceName]
// 	if !ok {
// 		return
// 	}
// 	for _, watcher := range watchers {
// 		select {
// 		case watcher <- instances:
// 		default:
// 			// 如果放不进去，跳过
// 			continue
// 		}
// 	}
// }
//
// func NewBaseDiscovery() *BaseDiscovery {
// 	discovery := &BaseDiscovery{
// 		watchers: make(map[string][]chan []ServiceInstance),
// 	}
// 	cache := NewServiceInstanceCache(discovery.onSetInstances)
// 	return discovery
// }
