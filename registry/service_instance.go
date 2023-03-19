package registry

import (
	"net/netip"
	"sync"

	"golang.org/x/exp/maps"
)

// ServiceInstance represents an instance of a service in a discovery system.
type ServiceInstance interface {
	// ID is the service instance ID.
	ID() string
	// Name is service instance name.
	Name() string
	// Address is service host:port address
	Address() netip.AddrPort
	// Metadata is other information the service instance carried.
	Metadata() map[string]string
	// Scheme is the scheme of the service instance.
	Scheme()
}

type DefaultServiceInstance struct {
	id       string
	name     string
	address  netip.AddrPort
	metadata map[string]string
	scheme   string
}

func NewDefaultServiceInstance(id string, name string, address netip.AddrPort, metadata map[string]string, scheme string) *DefaultServiceInstance {
	return &DefaultServiceInstance{id: id, name: name, address: address, metadata: metadata, scheme: scheme}
}

func (d *DefaultServiceInstance) ID() string {
	return d.id
}

func (d *DefaultServiceInstance) Name() string {
	return d.name
}

func (d *DefaultServiceInstance) Address() netip.AddrPort {
	return d.address
}

func (d *DefaultServiceInstance) Metadata() map[string]string {
	return d.metadata
}

func (d *DefaultServiceInstance) Scheme() string {
	return d.scheme
}

type SetInstancesListener interface {
	OnSetInstances(serviceName string, instances []ServiceInstance)
}

type OnSetInstances func(serviceName string, instances []ServiceInstance)

func (f OnSetInstances) OnSetInstances(serviceName string, instances []ServiceInstance) {
	f(serviceName, instances)
}

type ServiceInstanceCache struct {
	instancesMap         map[string][]ServiceInstance
	setInstancesListener SetInstancesListener
	sync.RWMutex
}

func NewServiceInstanceCache(listener SetInstancesListener) *ServiceInstanceCache {
	return &ServiceInstanceCache{instancesMap: make(map[string][]ServiceInstance), setInstancesListener: listener}
}

func (c *ServiceInstanceCache) SetInstances(serviceName string, instances []ServiceInstance) {
	c.Lock()
	defer c.Unlock()
	c.instancesMap[serviceName] = instances
	if c.setInstancesListener != nil {
		c.setInstancesListener.OnSetInstances(serviceName, instances)
	}
}

func (c *ServiceInstanceCache) GetInstances(serviceName string) ([]ServiceInstance, bool) {
	c.RLock()
	defer c.RUnlock()
	instances, ok := c.instancesMap[serviceName]
	return instances, ok
}

func (c *ServiceInstanceCache) GetServices() []string {
	c.RLock()
	defer c.RUnlock()
	return maps.Keys(c.instancesMap)
}
