// Code generated by protoc-gen-go-leo. DO NOT EDIT.
// If you want edit it, can move this file to another directory.

package cq

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
	v1 "github.com/go-leo/leo/v3/example/api/demo/v1"
)

var _ GetUserAvatarHandler = (*getUserAvatarHandler)(nil)

type GetUserAvatarHandler cqrs.QueryHandler[GetUserAvatarQuery, GetUserAvatarResult]

type GetUserAvatarQuery struct {
	v1.UnimplementedGetUserAvatarQuery
}

type GetUserAvatarResult struct {
	v1.UnimplementedGetUserAvatarResult
}

func NewGetUserAvatarHandler() GetUserAvatarHandler {
	return &getUserAvatarHandler{}
}

type getUserAvatarHandler struct {
}

func (h *getUserAvatarHandler) Handle(ctx context.Context, q GetUserAvatarQuery) (GetUserAvatarResult, error) {
	// TODO implement me
	panic("implement me")
}
