package metadatax

import (
	"context"
)

type _RawMD struct {
	md    Metadata
	added []Metadata
}

type incomingKey struct{}

// NewIncomingContext creates a new context with incoming _Metadata attached.
func NewIncomingContext(ctx context.Context, md Metadata) context.Context {
	return context.WithValue(ctx, incomingKey{}, md)
}

// FromIncomingContext returns the incoming metadata in ctx if it exists.
func FromIncomingContext(ctx context.Context) (Metadata, bool) {
	md, ok := ctx.Value(incomingKey{}).(Metadata)
	if !ok {
		return nil, false
	}
	return Join(md), true
}

type outgoingKey struct{}

// NewOutgoingContext creates a new context with outgoing Metadata attached.
func NewOutgoingContext(ctx context.Context, md Metadata) context.Context {
	return context.WithValue(ctx, outgoingKey{}, _RawMD{md: md})
}

// AppendOutgoingContext appends the Metadata to the context.
func AppendOutgoingContext(ctx context.Context, mds ...Metadata) context.Context {
	old, _ := ctx.Value(outgoingKey{}).(_RawMD)
	added := make([]Metadata, 0, len(old.added)+len(mds))
	added = append(added, old.added...)
	added = append(added, mds...)
	return context.WithValue(ctx, outgoingKey{}, _RawMD{md: old.md, added: added})
}

// FromOutgoingContext returns the outgoing metadata in ctx if it exists.
func FromOutgoingContext(ctx context.Context) (Metadata, bool) {
	rawMD, ok := ctx.Value(outgoingKey{}).(_RawMD)
	if !ok {
		return nil, false
	}
	res := Join(append([]Metadata{rawMD.md}, rawMD.added...)...)
	return res, ok
}
