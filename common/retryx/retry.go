package retryx

import (
	"context"
	"time"

	"github.com/go-leo/backoffx"
)

func Call(ctx context.Context, maxAttempts uint, backoffFunc backoffx.BackoffFunc, method func(attemptTime int) error) error {
	var err error
	max := int(maxAttempts)
	for i := 0; i <= max; i++ {
		// call method
		err = method(i)

		// if method not return error, no need to retry
		if err == nil {
			break
		}

		// If the maximum number of attempts is exceeded, no need to retry
		if i >= max {
			break
		}

		// sleep and wait retry
		time.Sleep(backoffFunc(ctx, uint(i+1)))
	}
	return err
}
