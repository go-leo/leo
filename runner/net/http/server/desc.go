package server

import (
	"context"
)

type ServiceDesc struct {
	ServiceName string
	HandlerType any
	Methods     []*MethodDesc
	Metadata    any
}

type MethodDesc struct {
	HTTPMethod         string
	MethodName         string
	Path               string
	RequestConstructor func() any
	Handler            func(cli any, ctx context.Context, in any) (any, error)
}
