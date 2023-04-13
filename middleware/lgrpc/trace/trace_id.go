package trace

import (
	"context"
	"encoding/hex"
	"math/rand"
	"time"

	"github.com/go-leo/gox/stringx"
	"go.opentelemetry.io/otel/trace"
)

var randSource = rand.New(rand.NewSource(time.Now().UnixNano()))

func TraceIDFromContext(ctx context.Context) string {
	spanContext := trace.SpanContextFromContext(ctx)
	if spanContext.HasTraceID() {
		return spanContext.TraceID().String()
	}
	return ""
}

func TraceIDFromContextOrGenerate(ctx context.Context) string {
	traceId := TraceIDFromContext(ctx)
	if stringx.IsNotBlank(traceId) {
		return traceId
	}
	var tid [16]byte
	randSource.Read(tid[:])
	return hex.EncodeToString(tid[:])
}

func NewContextWithTraceID(ctx context.Context, key string) context.Context {
	return context.WithValue(ctx, key, TraceIDFromContextOrGenerate(ctx))
}
