package lbx

import "github.com/go-kit/kit/endpoint"

type BalancerFunc func() (endpoint.Endpoint, error)

func (f BalancerFunc) Endpoint() (endpoint.Endpoint, error) {
	return f()
}
