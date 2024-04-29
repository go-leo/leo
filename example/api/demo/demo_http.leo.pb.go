// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package demo

import (
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	http "github.com/go-kit/kit/transport/http"
	endpointx "github.com/go-leo/kitx/endpointx"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func NewDemoServiceHTTPServer(
	endpoints interface {
		CreateUser() endpoint.Endpoint
		UpdateUser() endpoint.Endpoint
		GetUser() endpoint.Endpoint
		GetUsers() endpoint.Endpoint
		DeleteUser() endpoint.Endpoint
	},
	mdw []endpoint.Middleware,
	opts ...http.ServerOption,
) interface {
	CreateUser(ctx context.Context, request *CreateUserRequest) (*emptypb.Empty, error)
	UpdateUser(ctx context.Context, request *UpdateUserRequest) (*emptypb.Empty, error)
	GetUser(ctx context.Context, request *GetUserRequest) (*GetUserResponse, error)
	GetUsers(ctx context.Context, request *GetUsersRequest) (*GetUsersResponse, error)
	DeleteUser(ctx context.Context, request *DeleteUsersRequest) (*emptypb.Empty, error)
} {
	return &gRPCDemoServiceServer{
		createUser: http.NewServer(endpointx.Chain(endpoints.CreateUser(), mdw...), f, func(_ context.Context, v any) (any, error) { return v, nil }, opts...),
		updateUser: http.NewServer(endpointx.Chain(endpoints.UpdateUser(), mdw...), func(_ context.Context, v any) (any, error) { return v, nil }, func(_ context.Context, v any) (any, error) { return v, nil }, opts...),
		getUser:    http.NewServer(endpointx.Chain(endpoints.GetUser(), mdw...), func(_ context.Context, v any) (any, error) { return v, nil }, func(_ context.Context, v any) (any, error) { return v, nil }, opts...),
		getUsers:   http.NewServer(endpointx.Chain(endpoints.GetUsers(), mdw...), func(_ context.Context, v any) (any, error) { return v, nil }, func(_ context.Context, v any) (any, error) { return v, nil }, opts...),
		deleteUser: http.NewServer(endpointx.Chain(endpoints.DeleteUser(), mdw...), func(_ context.Context, v any) (any, error) { return v, nil }, func(_ context.Context, v any) (any, error) { return v, nil }, opts...),
	}
}