// Code generated by protoc-gen-go-leo. DO NOT EDIT.
// If you want edit it, can move this file to another directory.

package cq

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
	v1 "github.com/go-leo/leo/v3/example/api/workspace/v1"
)

var _ ListWorkspacesHandler = (*listWorkspacesHandler)(nil)

type ListWorkspacesHandler cqrs.QueryHandler[ListWorkspacesQuery, ListWorkspacesResult]

type ListWorkspacesQuery struct {
	v1.UnimplementedListWorkspacesQuery
}

func (ListWorkspacesQuery) From(ctx context.Context, req *v1.ListWorkspacesRequest) (v1.ListWorkspacesQuery, context.Context, error) {
	panic("implement me")
	return ListWorkspacesQuery{}, ctx, nil
}

type ListWorkspacesResult struct {
	v1.UnimplementedListWorkspacesResult
}

func (r ListWorkspacesResult) To(ctx context.Context) (*v1.ListWorkspacesResponse, error) {
	panic("implement me")
	return &v1.ListWorkspacesResponse{}, nil
}

func NewListWorkspacesHandler() ListWorkspacesHandler {
	return &listWorkspacesHandler{}
}

type listWorkspacesHandler struct {
}

func (h *listWorkspacesHandler) Handle(ctx context.Context, q ListWorkspacesQuery) (ListWorkspacesResult, error) {
	// TODO implement me
	panic("implement me")
}
