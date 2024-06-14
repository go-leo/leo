package lbx

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-leo/gox/backoff"
	"github.com/go-leo/leo/v3/endpointx"
	"github.com/go-leo/leo/v3/statusx"
	"time"
)

func Retry(
	ctx context.Context,
	max int,
	timeout time.Duration,
	backoffFunc backoff.BackoffFunc,
	b lb.Balancer,
) endpoint.Endpoint {
	if max == 0 {
		e, err := b.Endpoint()
		if err != nil {
			return endpointx.Error(err)
		}
		return e
	}
	callback := func(n int, err error) (bool, error) {
		statusErr, ok := statusx.FromError(err)
		if !ok {
			return false, nil
		}

		time.Sleep(backoffFunc(ctx, uint(n)))
		return n < max, nil
	}
	e := lb.RetryWithCallback(timeout, b, callback)
	return func(ctx context.Context, request any) (any, error) {
		resp, err := e(ctx, request)
		if err == nil {
			return resp, nil
		}
		var retryErr lb.RetryError
		if errors.As(err, &retryErr) {
			return nil, statusx.ErrNotFound
		}
		return nil, err
	}
}
