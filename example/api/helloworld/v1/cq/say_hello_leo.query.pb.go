// Code generated by protoc-gen-go-leo. DO NOT EDIT.
// If you want edit it, can move this file to another directory.

package cq

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
	v1 "github.com/go-leo/leo/v3/example/api/helloworld/v1"
)

var _ SayHelloHandler = (*sayHelloHandler)(nil)

type SayHelloHandler cqrs.QueryHandler[SayHelloQuery, SayHelloResult]

type SayHelloQuery struct {
	v1.UnimplementedSayHelloQuery
}

type SayHelloResult struct {
	v1.UnimplementedSayHelloResult
}

func NewSayHelloHandler() SayHelloHandler {
	return &sayHelloHandler{}
}

type sayHelloHandler struct {
}

func (h *sayHelloHandler) Handle(ctx context.Context, q SayHelloQuery) (SayHelloResult, error) {
	// TODO implement me
	panic("implement me")
}
