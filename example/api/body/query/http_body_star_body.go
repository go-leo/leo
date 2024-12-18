package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type HttpBodyStarBodyArgs struct {
}

type HttpBodyStarBodyRes struct {
}

type HttpBodyStarBody cqrs.QueryHandler[*HttpBodyStarBodyArgs, *HttpBodyStarBodyRes]

func NewHttpBodyStarBody() HttpBodyStarBody {
	return &httpBodyStarBody{}
}

type httpBodyStarBody struct {
}

func (h *httpBodyStarBody) Handle(ctx context.Context, args *HttpBodyStarBodyArgs) (*HttpBodyStarBodyRes, error) {
	// TODO implement me
	panic("implement me")
}
