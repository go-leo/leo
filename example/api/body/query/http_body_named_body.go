package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type HttpBodyNamedBodyArgs struct {
}

type HttpBodyNamedBodyRes struct {
}

type HttpBodyNamedBody cqrs.QueryHandler[*HttpBodyNamedBodyArgs, *HttpBodyNamedBodyRes]

func NewHttpBodyNamedBody() HttpBodyNamedBody {
	return &httpBodyNamedBody{}
}

type httpBodyNamedBody struct {
}

func (h *httpBodyNamedBody) Handle(ctx context.Context, args *HttpBodyNamedBodyArgs) (*HttpBodyNamedBodyRes, error) {
	// TODO implement me
	panic("implement me")
}
