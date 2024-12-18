package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type HttpBodyResponseArgs struct {
}

type HttpBodyResponseRes struct {
}

type HttpBodyResponse cqrs.QueryHandler[*HttpBodyResponseArgs, *HttpBodyResponseRes]

func NewHttpBodyResponse() HttpBodyResponse {
	return &httpBodyResponse{}
}

type httpBodyResponse struct {
}

func (h *httpBodyResponse) Handle(ctx context.Context, args *HttpBodyResponseArgs) (*HttpBodyResponseRes, error) {
	// TODO implement me
	panic("implement me")
}
