// Code generated by protoc-gen-go-leo. DO NOT EDIT.
// If you want edit it, can move this file to another directory.

package cq

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
	v1 "github.com/go-leo/leo/v3/example/api/library/v1"
)

var _ GetBookHandler = (*getBookHandler)(nil)

type GetBookHandler cqrs.QueryHandler[GetBookQuery, GetBookResult]

type GetBookQuery struct {
	v1.UnimplementedGetBookQuery
}

func (GetBookQuery) From(ctx context.Context, req *v1.GetBookRequest) (v1.GetBookQuery, context.Context, error) {
	panic("implement me")
	return GetBookQuery{}, ctx, nil
}

type GetBookResult struct {
	v1.UnimplementedGetBookResult
}

func (r GetBookResult) To(ctx context.Context) (*v1.Book, error) {
	panic("implement me")
	return &v1.Book{}, nil
}

func NewGetBookHandler() GetBookHandler {
	return &getBookHandler{}
}

type getBookHandler struct {
}

func (h *getBookHandler) Handle(ctx context.Context, q GetBookQuery) (GetBookResult, error) {
	// TODO implement me
	panic("implement me")
}
