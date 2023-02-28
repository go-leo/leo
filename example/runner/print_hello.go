package runner

import (
	"context"
	"fmt"
)

type PrintHello struct {
}

func (p PrintHello) Run(ctx context.Context) error {
	fmt.Println("hello leo once")
	return nil
}
