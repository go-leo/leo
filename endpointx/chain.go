package endpointx

import "github.com/go-kit/kit/endpoint"

// Chain decorates the given endpoint.Endpoint with all endpoint.Middleware.
func Chain(component endpoint.Endpoint, middlewares ...endpoint.Middleware) endpoint.Endpoint {
	if len(middlewares) == 0 {
		return component
	}
	middleware := endpoint.Chain(middlewares[0], middlewares[0:]...)
	return middleware(component)
}
