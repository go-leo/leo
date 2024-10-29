package httpx

import (
	"context"
	"github.com/go-leo/leo/v3/statusx"
	"net/http"
)

func IsErrorResponse(r *http.Response) bool {
	return r.Header.Get(kStatusCoderKey) != ""
}

func ErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	statusErr := statusx.From(err)
	w.Header().Set(kStatusCoderKey, kLeoCoderValue)
	statusErr.Write(w)
}

func ErrorDecoder(ctx context.Context, r *http.Response) error {
	return statusx.From(r)
}
