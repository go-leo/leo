package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type SayHelloArgs struct {
}

type SayHelloRes struct {
}

type SayHello cqrs.QueryHandler[*SayHelloArgs, *SayHelloRes]

func NewSayHello() SayHello {
	return &sayHello{}
}

type sayHello struct {
}

func (h *sayHello) Handle(ctx context.Context, args *SayHelloArgs) (*SayHelloRes, error) {
	// TODO implement me
	panic("implement me")
}
