// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package demo

import (
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	http "google.golang.org/genproto/googleapis/rpc/http"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type demoServiceEndpoints struct {
	svc interface {
		CreateUser(ctx context.Context, request *CreateUserRequest) (*emptypb.Empty, error)
		UpdateUser(ctx context.Context, request *UpdateUserRequest) (*emptypb.Empty, error)
		GetUser(ctx context.Context, request *GetUserRequest) (*GetUserResponse, error)
		GetUsers(ctx context.Context, request *GetUsersRequest) (*GetUsersResponse, error)
		DeleteUser(ctx context.Context, request *DeleteUsersRequest) (*emptypb.Empty, error)
		UpdateUserName(ctx context.Context, request *UpdateUserNameRequest) (*emptypb.Empty, error)
		UploadUsers(ctx context.Context, request *httpbody.HttpBody) (*httpbody.HttpBody, error)
		UploadUserAvatar(ctx context.Context, request *UploadUserAvatarRequest) (*UploadUserAvatarResponse, error)
		PushUsers(ctx context.Context, request *http.HttpRequest) (*http.HttpResponse, error)
		PushUserAvatar(ctx context.Context, request *PushUserAvatarRequest) (*PushUserAvatarResponse, error)
		ModifyUser(ctx context.Context, request *ModifyUserRequest) (*emptypb.Empty, error)
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

func (e *demoServiceEndpoints) UpdateUserName() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.UpdateUserName(ctx, request.(*UpdateUserNameRequest))
	}
}

func (e *demoServiceEndpoints) UploadUsers() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.UploadUsers(ctx, request.(*httpbody.HttpBody))
	}
}

func (e *demoServiceEndpoints) UploadUserAvatar() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.UploadUserAvatar(ctx, request.(*UploadUserAvatarRequest))
	}
}

func (e *demoServiceEndpoints) PushUsers() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.PushUsers(ctx, request.(*http.HttpRequest))
	}
}

func (e *demoServiceEndpoints) PushUserAvatar() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.PushUserAvatar(ctx, request.(*PushUserAvatarRequest))
	}
}

func (e *demoServiceEndpoints) ModifyUser() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.ModifyUser(ctx, request.(*ModifyUserRequest))
	}
}

func NewdemoServiceEndpoints(
	svc interface {
		CreateUser(ctx context.Context, request *CreateUserRequest) (*emptypb.Empty, error)
		UpdateUser(ctx context.Context, request *UpdateUserRequest) (*emptypb.Empty, error)
		GetUser(ctx context.Context, request *GetUserRequest) (*GetUserResponse, error)
		GetUsers(ctx context.Context, request *GetUsersRequest) (*GetUsersResponse, error)
		DeleteUser(ctx context.Context, request *DeleteUsersRequest) (*emptypb.Empty, error)
		UpdateUserName(ctx context.Context, request *UpdateUserNameRequest) (*emptypb.Empty, error)
		UploadUsers(ctx context.Context, request *httpbody.HttpBody) (*httpbody.HttpBody, error)
		UploadUserAvatar(ctx context.Context, request *UploadUserAvatarRequest) (*UploadUserAvatarResponse, error)
		PushUsers(ctx context.Context, request *http.HttpRequest) (*http.HttpResponse, error)
		PushUserAvatar(ctx context.Context, request *PushUserAvatarRequest) (*PushUserAvatarResponse, error)
		ModifyUser(ctx context.Context, request *ModifyUserRequest) (*emptypb.Empty, error)
	},
) interface {
	CreateUser() endpoint.Endpoint
	UpdateUser() endpoint.Endpoint
	GetUser() endpoint.Endpoint
	GetUsers() endpoint.Endpoint
	DeleteUser() endpoint.Endpoint
	UpdateUserName() endpoint.Endpoint
	UploadUsers() endpoint.Endpoint
	UploadUserAvatar() endpoint.Endpoint
	PushUsers() endpoint.Endpoint
	PushUserAvatar() endpoint.Endpoint
	ModifyUser() endpoint.Endpoint
} {
	return &demoServiceEndpoints{svc: svc}
}
