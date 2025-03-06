// Code generated by protoc-gen-go-leo. DO NOT EDIT.
// If you want edit it, can move this file to another directory.

package cq

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
	v1 "github.com/go-leo/leo/v3/example/api/route/v1"
)

var _ HttpRequestHandler = (*httpRequestHandler)(nil)

type HttpRequestHandler cqrs.CommandHandler[HttpRequestCommand]

type HttpRequestCommand struct {
	v1.UnimplementedHttpRequestCommand
}

func NewHttpRequestHandler() HttpRequestHandler {
	return &httpRequestHandler{}
}

type httpRequestHandler struct {
}

func (h *httpRequestHandler) Handle(ctx context.Context, cmd HttpRequestCommand) error {
	// TODO implement me
	panic("implement me")
}
