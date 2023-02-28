package runner

import (
	"context"
	"fmt"
	"time"
)

type LoopPrintHello struct {
}

func (p LoopPrintHello) Run(ctx context.Context) error {
	tickers := time.NewTicker(time.Second)
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-tickers.C:
			fmt.Println("hello leo loop")
		}
	}
}
