// Code generated by protoc-gen-go-leo. DO NOT EDIT.
// If you want edit it, can move this file to another directory.

package cq

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
	v1 "github.com/go-leo/leo/v3/example/api/library/v1"
)

var _ MergeShelvesHandler = (*mergeShelvesHandler)(nil)

type MergeShelvesHandler cqrs.QueryHandler[MergeShelvesQuery, MergeShelvesResult]

type MergeShelvesQuery struct {
	v1.UnimplementedMergeShelvesQuery
}

type MergeShelvesResult struct {
	v1.UnimplementedMergeShelvesResult
}

func NewMergeShelvesHandler() MergeShelvesHandler {
	return &mergeShelvesHandler{}
}

type mergeShelvesHandler struct {
}

func (h *mergeShelvesHandler) Handle(ctx context.Context, q MergeShelvesQuery) (MergeShelvesResult, error) {
	// TODO implement me
	panic("implement me")
}
