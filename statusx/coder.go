package statusx

import (
	"context"
	"net/http"
)

type Encoder interface {
	Encode(ctx context.Context, err *Error) (int, http.Header, []byte)
}

type Decoder interface {
	Decode(ctx context.Context, status int, header http.Header, body []byte) *Error
}
