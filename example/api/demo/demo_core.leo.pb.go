// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package demo

import (
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type demoServiceEndpoints struct {
	svc interface {
		CreateUser(ctx context.Context, request *CreateUserRequest) (*emptypb.Empty, error)
		UpdateUser(ctx context.Context, request *UpdateUserRequest) (*emptypb.Empty, error)
		GetUser(ctx context.Context, request *GetUserRequest) (*GetUserResponse, error)
		GetUsers(ctx context.Context, request *GetUsersRequest) (*GetUsersResponse, error)
		DeleteUser(ctx context.Context, request *DeleteUsersRequest) (*emptypb.Empty, error)
	}
}

func (e *demoServiceEndpoints) CreateUser() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.CreateUser(ctx, request.(*CreateUserRequest))
	}
}

func (e *demoServiceEndpoints) UpdateUser() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.UpdateUser(ctx, request.(*UpdateUserRequest))
	}
}

func (e *demoServiceEndpoints) GetUser() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.GetUser(ctx, request.(*GetUserRequest))
	}
}

func (e *demoServiceEndpoints) GetUsers() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.GetUsers(ctx, request.(*GetUsersRequest))
	}
}

func (e *demoServiceEndpoints) DeleteUser() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.DeleteUser(ctx, request.(*DeleteUsersRequest))
	}
}

func NewdemoServiceEndpoints(
	svc interface {
		CreateUser(ctx context.Context, request *CreateUserRequest) (*emptypb.Empty, error)
		UpdateUser(ctx context.Context, request *UpdateUserRequest) (*emptypb.Empty, error)
		GetUser(ctx context.Context, request *GetUserRequest) (*GetUserResponse, error)
		GetUsers(ctx context.Context, request *GetUsersRequest) (*GetUsersResponse, error)
		DeleteUser(ctx context.Context, request *DeleteUsersRequest) (*emptypb.Empty, error)
	},
) interface {
	CreateUser() endpoint.Endpoint
	UpdateUser() endpoint.Endpoint
	GetUser() endpoint.Endpoint
	GetUsers() endpoint.Endpoint
	DeleteUser() endpoint.Endpoint
} {
	return &demoServiceEndpoints{svc: svc}
}
