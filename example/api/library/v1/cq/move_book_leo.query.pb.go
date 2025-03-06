// Code generated by protoc-gen-go-leo. DO NOT EDIT.
// If you want edit it, can move this file to another directory.

package cq

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
	v1 "github.com/go-leo/leo/v3/example/api/library/v1"
)

var _ MoveBookHandler = (*moveBookHandler)(nil)

type MoveBookHandler cqrs.QueryHandler[MoveBookQuery, MoveBookResult]

type MoveBookQuery struct {
	v1.UnimplementedMoveBookQuery
}

func (MoveBookQuery) From(ctx context.Context, req *v1.MoveBookRequest) (v1.MoveBookQuery, context.Context, error) {
	panic("implement me")
	return MoveBookQuery{}, ctx, nil
}

type MoveBookResult struct {
	v1.UnimplementedMoveBookResult
}

func (r MoveBookResult) To(ctx context.Context) (*v1.Book, error) {
	panic("implement me")
	return &v1.Book{}, nil
}

func NewMoveBookHandler() MoveBookHandler {
	return &moveBookHandler{}
}

type moveBookHandler struct {
}

func (h *moveBookHandler) Handle(ctx context.Context, q MoveBookQuery) (MoveBookResult, error) {
	// TODO implement me
	panic("implement me")
}
