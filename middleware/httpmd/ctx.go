package httpmd

import (
	"context"
)

type key struct{}

func NewContext(ctx context.Context, md *Metadata) context.Context {
	return context.WithValue(ctx, key{}, md)
}

func FromContext(ctx context.Context) (md *Metadata, ok bool) {
	md, ok = ctx.Value(key{}).(*Metadata)
	return
}
