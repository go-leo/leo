package contextx

import (
	"context"
	"fmt"
)

type contextError struct {
	err      error
	causeErr error
}

func (c contextError) Error() string {
	return fmt.Sprintf("%v, because %v", c.err, c.causeErr)
}

func Error(ctx context.Context) error {
	err := ctx.Err()
	if err == nil {
		return nil
	}
	causeErr := context.Cause(ctx)
	if causeErr == nil {
		return err
	}
	return contextError{err: err, causeErr: causeErr}
}
