// Code generated by protoc-gen-leo-core. DO NOT EDIT.

package helloworld

import (
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	sd "github.com/go-kit/kit/sd"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	transportx "github.com/go-leo/leo/v3/transportx"
)

type GreeterService interface {
	SayHello(ctx context.Context, request *HelloRequest) (*HelloReply, error)
}

type GreeterEndpoints interface {
	SayHello(ctx context.Context) endpoint.Endpoint
}

type GreeterClientTransports interface {
	SayHello() transportx.ClientTransport
}

type GreeterFactories interface {
	SayHello(middlewares ...endpoint.Middleware) sd.Factory
}

type GreeterEndpointers interface {
	SayHello() sd.Endpointer
}

type greeterServerEndpoints struct {
	svc         GreeterService
	middlewares []endpoint.Middleware
}

func (e *greeterServerEndpoints) SayHello(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.SayHello(ctx, request.(*HelloRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func newGreeterServerEndpoints(svc GreeterService, middlewares ...endpoint.Middleware) GreeterEndpoints {
	return &greeterServerEndpoints{svc: svc, middlewares: middlewares}
}

type greeterClientEndpoints struct {
	transports  GreeterClientTransports
	middlewares []endpoint.Middleware
}

func (e *greeterClientEndpoints) SayHello(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.SayHello().Endpoint(ctx), e.middlewares...)
}

func newGreeterClientEndpoints(transports GreeterClientTransports, middlewares ...endpoint.Middleware) GreeterEndpoints {
	return &greeterClientEndpoints{transports: transports, middlewares: middlewares}
}
