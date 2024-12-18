package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type HttpBodyNamedResponseArgs struct {
}

type HttpBodyNamedResponseRes struct {
}

type HttpBodyNamedResponse cqrs.QueryHandler[*HttpBodyNamedResponseArgs, *HttpBodyNamedResponseRes]

func NewHttpBodyNamedResponse() HttpBodyNamedResponse {
	return &httpBodyNamedResponse{}
}

type httpBodyNamedResponse struct {
}

func (h *httpBodyNamedResponse) Handle(ctx context.Context, args *HttpBodyNamedResponseArgs) (*HttpBodyNamedResponseRes, error) {
	// TODO implement me
	panic("implement me")
}
