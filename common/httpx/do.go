package httpx

import (
	"context"
	"errors"
	"fmt"
)

// Deprecated: Do not use. use github.com/go-leo/netx instead.
type DoCommand struct{}

// Deprecated: Do not use. use github.com/go-leo/netx instead.
func (cmd *DoCommand) Execute(ctx context.Context) (context.Context, error) {
	cli, ok := ClientFromContext(ctx)
	if !ok {
		return ctx, errors.New("http client is nil")
	}
	req, ok := RequestFromContext(ctx)
	if !ok {
		return ctx, errors.New("http request is nil")
	}
	resp, err := cli.Do(req)
	if err != nil {
		return ctx, fmt.Errorf("failed to send http request, %w", err)
	}
	return NewContextWithResponse(ctx, resp), nil
}
