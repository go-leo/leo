package client

import "context"

type Interface interface {
	Invoke(ctx context.Context, method string, path string, in any, out any) error
}

type InterfaceMapper interface {
	Interfaces() (map[string]Interface, error)
}

type InterfaceBalancer interface {
	Pick(map[string]Interface) (Interface, error)
}
