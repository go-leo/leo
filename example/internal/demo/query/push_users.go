package query

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/go-leo/leo/v3/cqrs"
)

type PushUsersArgs struct {
	Data []byte
}

type PushUsersRes struct {
	Data []byte
}

type PushUsers cqrs.QueryHandler[*PushUsersArgs, *PushUsersRes]

func NewPushUsers() PushUsers {
	return &pushUsers{}
}

type pushUsers struct {
}

func (h *pushUsers) Handle(ctx context.Context, args *PushUsersArgs) (*PushUsersRes, error) {
	fmt.Println("push users", args)
	b := make([]byte, 128)
	_, _ = rand.Read(b)
	return &PushUsersRes{
		Data: b,
	}, nil
}
