package httpheader

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"
)

type key struct{}

func NewContext(ctx context.Context, h http.Header) context.Context {
	return context.WithValue(ctx, key{}, h)
}

func FromContext(ctx context.Context) (h http.Header, ok bool) {
	h, ok = ctx.Value(key{}).(http.Header)
	return
}

const outerKey = "HTTP-Header"

func FromIncoming(ctx context.Context) (string, bool) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		vals := md.Get(outerKey)
		if len(vals) > 0 {
			return vals[0], true
		}
	}
	return "", false
}

func ToOutgoing(ctx context.Context, v string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, outerKey, v)
}
