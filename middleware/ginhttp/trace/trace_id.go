package trace

import (
	"context"
	"encoding/hex"
	"math/rand"
	"time"

	"go.opentelemetry.io/otel/trace"
)

var randSource = rand.New(rand.NewSource(time.Now().UnixNano()))

func TraceIDFromContext(ctx context.Context) (string, bool) {
	spanContext := trace.SpanContextFromContext(ctx)
	if spanContext.HasTraceID() {
		return spanContext.TraceID().String(), true
	}
	return "", false
}

func TraceIDFromContextOrGenerate(ctx context.Context) string {
	traceId, ok := TraceIDFromContext(ctx)
	if ok {
		return traceId
	}
	var tid [16]byte
	randSource.Read(tid[:])
	return hex.EncodeToString(tid[:])
}
