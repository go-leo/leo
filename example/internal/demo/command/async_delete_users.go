package command

import (
	"context"
	"fmt"
	"github.com/go-leo/cqrs"
	"time"
)

type AsyncDeleteUsersArgs struct {
	Names []string
}

type AsyncDeleteUsers cqrs.CommandHandler[*AsyncDeleteUsersArgs]

func NewAsyncDeleteUsers() AsyncDeleteUsers {
	return &asyncDeleteUsers{}
}

type asyncDeleteUsers struct {
}

func (h *asyncDeleteUsers) Handle(ctx context.Context, args *AsyncDeleteUsersArgs) error {
	time.Sleep(10 * time.Second)
	fmt.Println("delete users", args)
	return nil
}
