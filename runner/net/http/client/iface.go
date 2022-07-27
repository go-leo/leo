package client

import "context"

type Interface interface {
	Invoke(ctx context.Context, method string, path string, in any, out any) error
}
