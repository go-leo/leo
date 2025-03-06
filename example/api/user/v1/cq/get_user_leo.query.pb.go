// Code generated by protoc-gen-go-leo. DO NOT EDIT.
// If you want edit it, can move this file to another directory.

package cq

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
	v1 "github.com/go-leo/leo/v3/example/api/user/v1"
)

var _ GetUserHandler = (*getUserHandler)(nil)

type GetUserHandler cqrs.QueryHandler[GetUserQuery, GetUserResult]

type GetUserQuery struct {
	v1.UnimplementedGetUserQuery
}

func (GetUserQuery) From(ctx context.Context, req *v1.GetUserRequest) (v1.GetUserQuery, context.Context, error) {
	panic("implement me")
	return GetUserQuery{}, ctx, nil
}

type GetUserResult struct {
	v1.UnimplementedGetUserResult
}

func (r GetUserResult) To(ctx context.Context) (*v1.GetUserResponse, error) {
	panic("implement me")
	return &v1.GetUserResponse{}, nil
}

func NewGetUserHandler() GetUserHandler {
	return &getUserHandler{}
}

type getUserHandler struct {
}

func (h *getUserHandler) Handle(ctx context.Context, q GetUserQuery) (GetUserResult, error) {
	// TODO implement me
	panic("implement me")
}
