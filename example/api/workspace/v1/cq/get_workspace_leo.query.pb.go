// Code generated by protoc-gen-go-leo. DO NOT EDIT.
// If you want edit it, can move this file to another directory.

package cq

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
	v1 "github.com/go-leo/leo/v3/example/api/workspace/v1"
)

var _ GetWorkspaceHandler = (*getWorkspaceHandler)(nil)

type GetWorkspaceHandler cqrs.QueryHandler[GetWorkspaceQuery, GetWorkspaceResult]

type GetWorkspaceQuery struct {
	v1.UnimplementedGetWorkspaceQuery
}

type GetWorkspaceResult struct {
	v1.UnimplementedGetWorkspaceResult
}

func NewGetWorkspaceHandler() GetWorkspaceHandler {
	return &getWorkspaceHandler{}
}

type getWorkspaceHandler struct {
}

func (h *getWorkspaceHandler) Handle(ctx context.Context, q GetWorkspaceQuery) (GetWorkspaceResult, error) {
	// TODO implement me
	panic("implement me")
}
