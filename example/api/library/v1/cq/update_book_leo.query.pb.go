// Code generated by protoc-gen-go-leo. DO NOT EDIT.
// If you want edit it, can move this file to another directory.

package cq

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
	v1 "github.com/go-leo/leo/v3/example/api/library/v1"
)

var _ UpdateBookHandler = (*updateBookHandler)(nil)

type UpdateBookHandler cqrs.QueryHandler[UpdateBookQuery, UpdateBookResult]

type UpdateBookQuery struct {
	v1.UnimplementedUpdateBookQuery
}

type UpdateBookResult struct {
	v1.UnimplementedUpdateBookResult
}

func NewUpdateBookHandler() UpdateBookHandler {
	return &updateBookHandler{}
}

type updateBookHandler struct {
}

func (h *updateBookHandler) Handle(ctx context.Context, q UpdateBookQuery) (UpdateBookResult, error) {
	// TODO implement me
	panic("implement me")
}
