// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package query

import (
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type queryEndpoints struct {
	svc interface {
		Query(ctx context.Context, request *QueryRequest) (*emptypb.Empty, error)
	}
}

func (e *queryEndpoints) Query() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.Query(ctx, request.(*QueryRequest))
	}
}

func NewQueryEndpoints(
	svc interface {
		Query(ctx context.Context, request *QueryRequest) (*emptypb.Empty, error)
	},
) interface {
	Query() endpoint.Endpoint
} {
	return &queryEndpoints{svc: svc}
}
