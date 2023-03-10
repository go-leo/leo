package registry

// ServiceInstance represents an instance of a service in a discovery system.
type ServiceInstance interface {
	// InstanceID is the unique instance ID as registered.
	InstanceID() string
	// ID is the service ID as registered.
	ID() string
	// Name is service name.
	Name() string
	// Kind is service kind.
	Kind() string
	// Host is the hostname of the registered service instance.
	Host() string
	// Port is the port of the registered service instance.
	Port() int
	// Metadata is other information the service carried.
	Metadata() map[string]string
	// Version is version of the service
	Version() string
	// Address is host:port
	Address() string
}

type serviceInstance struct {
	// InstanceID is the unique instance ID as registered.
	InstanceID string
	// ID is the service ID as registered.
	ID string
	// Name is service name.
	Name string
	// Kind is service kind.
	Kind string
	// Host is the hostname of the registered service instance.
	Host string
	// Port is the port of the registered service instance.
	Port int
	// Metadata is other information the service carried.
	Metadata map[string]string
	// Version is version of the service
	Version string
	// Address is host:port
	Address string
}
