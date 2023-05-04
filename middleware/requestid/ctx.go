package requestid

import (
	"context"
	"encoding/hex"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/metadata"

	"github.com/go-leo/stringx"

	"github.com/go-leo/leo/runner/net/http/header"
)

var randPool = sync.Pool{New: func() any { return rand.New(rand.NewSource(time.Now().UnixNano())) }}

const outerKey = "x-leo-request-id"

type innerKey struct{}

func FromContext(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(innerKey{}).(string)
	return val, ok
}

func NewContext(ctx context.Context, v string) context.Context {
	return context.WithValue(ctx, innerKey{}, v)
}

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

func FromHeader(ctx context.Context) (string, bool) {
	h, ok := header.FromContext(ctx)
	if !ok {
		return "", false
	}
	val := h.Get(outerKey)
	if stringx.IsBlank(val) {
		return "", false
	}
	return val, true
}

func ToHeader(ctx context.Context, v string) context.Context {
	h, ok := header.FromContext(ctx)
	if !ok {
		h = make(http.Header)
		ctx = header.NewContext(ctx, h)
	}
	h.Set(outerKey, v)
	return ctx
}

func FromTrace(ctx context.Context) (string, bool) {
	spanContext := trace.SpanContextFromContext(ctx)
	if spanContext.HasTraceID() {
		return spanContext.TraceID().String(), true
	}
	return "", false
}

func FromAnyWhere(ctx context.Context) (requestID string, generated bool) {
	// 1. from context
	requestID, _ = FromContext(ctx)
	if stringx.IsNotBlank(requestID) {
		return requestID, true
	}
	// 2. from grpc incoming
	requestID, _ = FromIncoming(ctx)
	if stringx.IsNotBlank(requestID) {
		return requestID, true
	}
	// 3. from http header
	requestID, _ = FromHeader(ctx)
	if stringx.IsNotBlank(requestID) {
		return requestID, true
	}
	// 3. from trace system traceID
	requestID, _ = FromTrace(ctx)
	if stringx.IsNotBlank(requestID) {
		return requestID, true
	}
	// 4. generate
	return generate(), false
}

func generate() (requestID string) {
	var tid [16]byte
	randSource := randPool.Get().(*rand.Rand)
	randSource.Read(tid[:])
	randPool.Put(randSource)
	requestID = hex.EncodeToString(tid[:])
	return
}
