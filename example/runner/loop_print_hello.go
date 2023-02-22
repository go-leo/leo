package runner

import (
	"context"
	"fmt"
)

type LoopPrintHello struct {
}

func (p LoopPrintHello) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			fmt.Println("hello leo")
		}
	}
	return nil
}
