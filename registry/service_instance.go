package registry

// ServiceInstance represents an instance of a service in a discovery system.
type ServiceInstance interface {
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
	// Address is host:port
	Address() string
}
