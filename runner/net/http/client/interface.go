package client

import "context"

// Interface defines the functions clients need to perform RPCs.
// It is implemented by *Client and *lb.BalancedClient, and is only intended to
// be referenced by generated code.
type Interface interface {
	Invoke(ctx context.Context, method string, path string, in any, out any) error
}
