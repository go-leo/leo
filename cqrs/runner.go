package cqrs

import (
	"context"
	"github.com/go-leo/leo/v3/runner"
)

var (
	_ runner.Runner = (*CommandRunner)(nil)
	_ runner.Runner = (*QueryRunner)(nil)
)

type CommandRunner[C any] struct {
	CommandHandler CommandHandler[C]
	Command        C
}

func (r *CommandRunner[C]) Run(ctx context.Context) error {
	return r.CommandHandler.Handle(ctx, r.Command)
}

type QueryRunner[Q any, R any] struct {
	QueryHandler QueryHandler[Q, R]
	Query        Q
	Result       R
}

func (r *QueryRunner[Q, R]) Run(ctx context.Context) error {
	res, err := r.QueryHandler.Handle(ctx, r.Query)
	if err != nil {
		return err
	}
	r.Result = res
	return nil
}
