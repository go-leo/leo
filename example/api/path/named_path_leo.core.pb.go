// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package path

import (
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type namedPathEndpoints struct {
	svc interface {
		NamedPathString(ctx context.Context, request *NamedPathRequest) (*emptypb.Empty, error)
		NamedPathOptString(ctx context.Context, request *NamedPathRequest) (*emptypb.Empty, error)
		NamedPathWrapString(ctx context.Context, request *NamedPathRequest) (*emptypb.Empty, error)
		EmbedNamedPathString(ctx context.Context, request *EmbedNamedPathRequest) (*emptypb.Empty, error)
		EmbedNamedPathOptString(ctx context.Context, request *EmbedNamedPathRequest) (*emptypb.Empty, error)
		EmbedNamedPathWrapString(ctx context.Context, request *EmbedNamedPathRequest) (*emptypb.Empty, error)
	}
}

func (e *namedPathEndpoints) NamedPathString() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.NamedPathString(ctx, request.(*NamedPathRequest))
	}
}

func (e *namedPathEndpoints) NamedPathOptString() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.NamedPathOptString(ctx, request.(*NamedPathRequest))
	}
}

func (e *namedPathEndpoints) NamedPathWrapString() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.NamedPathWrapString(ctx, request.(*NamedPathRequest))
	}
}

func (e *namedPathEndpoints) EmbedNamedPathString() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.EmbedNamedPathString(ctx, request.(*EmbedNamedPathRequest))
	}
}

func (e *namedPathEndpoints) EmbedNamedPathOptString() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.EmbedNamedPathOptString(ctx, request.(*EmbedNamedPathRequest))
	}
}

func (e *namedPathEndpoints) EmbedNamedPathWrapString() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.EmbedNamedPathWrapString(ctx, request.(*EmbedNamedPathRequest))
	}
}

func NewNamedPathEndpoints(
	svc interface {
		NamedPathString(ctx context.Context, request *NamedPathRequest) (*emptypb.Empty, error)
		NamedPathOptString(ctx context.Context, request *NamedPathRequest) (*emptypb.Empty, error)
		NamedPathWrapString(ctx context.Context, request *NamedPathRequest) (*emptypb.Empty, error)
		EmbedNamedPathString(ctx context.Context, request *EmbedNamedPathRequest) (*emptypb.Empty, error)
		EmbedNamedPathOptString(ctx context.Context, request *EmbedNamedPathRequest) (*emptypb.Empty, error)
		EmbedNamedPathWrapString(ctx context.Context, request *EmbedNamedPathRequest) (*emptypb.Empty, error)
	},
) interface {
	NamedPathString() endpoint.Endpoint
	NamedPathOptString() endpoint.Endpoint
	NamedPathWrapString() endpoint.Endpoint
	EmbedNamedPathString() endpoint.Endpoint
	EmbedNamedPathOptString() endpoint.Endpoint
	EmbedNamedPathWrapString() endpoint.Endpoint
} {
	return &namedPathEndpoints{svc: svc}
}
