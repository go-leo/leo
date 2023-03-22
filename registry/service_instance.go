package registry

import (
	"sync"
)

// ServiceInstance represents an instance of a service in a discovery system.
type ServiceInstance interface {
	// ID is the service instance ID.
	ID() string
	// Name is service instance name.
	Name() string
	// IP is ip
	IP() string
	// Port is port
	Port() int
	// Metadata is other information the service instance carried.
	Metadata() map[string]string
	// Scheme is the scheme of the service instance.
	Scheme() string
}

var _ ServiceInstance = new(serviceInstance)

type serviceInstance struct {
	id       string
	name     string
	ip       string
	port     int
	metadata map[string]string
	scheme   string
}

func NewServiceInstance(
	id string, name string, ip string, port int,
	metadata map[string]string, scheme string,
) ServiceInstance {
	return &serviceInstance{id: id, name: name, ip: ip, port: port, metadata: metadata, scheme: scheme}
}

func (d *serviceInstance) ID() string {
	return d.id
}

func (d *serviceInstance) Name() string {
	return d.name
}

func (d *serviceInstance) IP() string {
	return d.ip
}

func (d *serviceInstance) Port() int {
	return d.port
}

func (d *serviceInstance) Metadata() map[string]string {
	return d.metadata
}

func (d *serviceInstance) Scheme() string {
	return d.scheme
}

type Cache interface {
	SetInstances(serviceName string, instances []ServiceInstance)
	GetInstances(serviceName string) ([]ServiceInstance, bool)
	SetServices(services []string)
	GetServices() []string
}

type MemeryCache struct {
	instancesMap map[string][]ServiceInstance
	services     []string
	sync.RWMutex
}

func NewMemeryCache() *MemeryCache {
	return &MemeryCache{instancesMap: make(map[string][]ServiceInstance)}
}

func (c *MemeryCache) SetInstances(serviceName string, instances []ServiceInstance) {
	c.Lock()
	defer c.Unlock()
	c.instancesMap[serviceName] = instances
}

func (c *MemeryCache) GetInstances(serviceName string) ([]ServiceInstance, bool) {
	c.RLock()
	defer c.RUnlock()
	instances, ok := c.instancesMap[serviceName]
	return instances, ok
}

func (c *MemeryCache) SetServices(services []string) {
	c.Lock()
	defer c.Unlock()
	c.services = services
}

func (c *MemeryCache) GetServices() []string {
	c.RLock()
	defer c.RUnlock()
	return c.services
}
