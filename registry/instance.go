package registry

import (
	"fmt"

	"github.com/go-leo/gox/mathx/randx"
	"github.com/go-leo/gox/stringx"
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

type ServiceInstanceBuilder struct {
	id       string
	name     string
	ip       string
	port     int
	metadata map[string]string
	scheme   string
}

func (b *ServiceInstanceBuilder) ID(id string) *ServiceInstanceBuilder {
	b.id = id
	return b
}

func (b *ServiceInstanceBuilder) Name(name string) *ServiceInstanceBuilder {
	b.name = name
	return b
}

func (b *ServiceInstanceBuilder) IP(ip string) *ServiceInstanceBuilder {
	b.ip = ip
	return b
}

func (b *ServiceInstanceBuilder) Port(port int) *ServiceInstanceBuilder {
	b.port = port
	return b
}

func (b *ServiceInstanceBuilder) Metadata(metadata map[string]string) *ServiceInstanceBuilder {
	b.metadata = metadata
	return b
}

func (b *ServiceInstanceBuilder) Scheme(scheme string) *ServiceInstanceBuilder {
	b.scheme = scheme
	return b
}

func (b *ServiceInstanceBuilder) Build() ServiceInstance {
	if stringx.IsBlank(b.id) {
		b.id = fmt.Sprintf("%s.%s.%d.%s", b.name, b.ip, b.port, randx.WordString(8))
	}
	if b.metadata == nil {
		b.metadata = make(map[string]string)
	}
	return &serviceInstance{
		id:       b.id,
		name:     b.name,
		ip:       b.ip,
		port:     b.port,
		metadata: b.metadata,
		scheme:   b.scheme,
	}
}

func Builder() *ServiceInstanceBuilder {
	return new(ServiceInstanceBuilder)
}
