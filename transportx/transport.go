package transportx

import "github.com/go-kit/kit/endpoint"

// Transport is a transport that can be used to invoke a remote endpoint.
type Transport interface {
	// Endpoint returns a usable endpoint that invokes the remote endpoint.
	Endpoint() endpoint.Endpoint
}
