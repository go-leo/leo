package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type QueryArgs struct {
}

type QueryRes struct {
}

type Query cqrs.QueryHandler[*QueryArgs, *QueryRes]

func NewQuery() Query {
	return &query{}
}

type query struct {
}

func (h *query) Handle(ctx context.Context, args *QueryArgs) (*QueryRes, error) {
	// TODO implement me
	panic("implement me")
}
