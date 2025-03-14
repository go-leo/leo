package timeoutx

import (
	"context"
	"github.com/go-leo/leo/v3/logx"
	"net/http"
	"time"
)

const metadataKey = "X-Leo-Timeout"

type Key struct{}

func OutgoingInjector(ctx context.Context, request *http.Request) context.Context {
	if deadline, ok := ctx.Deadline(); ok {
		request.Header.Set(metadataKey, encodeDuration(time.Until(deadline)))
	}
	return ctx
}

func IncomingInjector(ctx context.Context, request *http.Request) context.Context {
	if value := request.Header.Get(metadataKey); value != "" {
		timeout, err := decodeTimeout(value)
		if err != nil {
			logx.Error(ctx, "error", err)
		}
		ctx, cancelFunc := context.WithTimeout(ctx, timeout)
		return context.WithValue(ctx, Key{}, cancelFunc)
	}
	ctx, cancelFunc := context.WithCancel(ctx)
	return context.WithValue(ctx, Key{}, cancelFunc)
}

func CancelInvoker(ctx context.Context, code int, r *http.Request) {
	cancelFunc, ok := ctx.Value(Key{}).(context.CancelFunc)
	if ok {
		cancelFunc()
	}
}
