package trace

import (
	"context"
	"encoding/hex"
	"math/rand"
	"time"

	"go.opentelemetry.io/otel/trace"
)

var randSource = rand.New(rand.NewSource(time.Now().UnixNano()))

func FromContext(ctx context.Context) (string, bool) {
	spanContext := trace.SpanContextFromContext(ctx)
	if spanContext.HasTraceID() {
		return spanContext.TraceID().String(), true
	}
	return "", false
}

func FromContextOrGenerate(ctx context.Context) string {
	traceId, ok := FromContext(ctx)
	if ok {
		return traceId
	}
	var tid [16]byte
	randSource.Read(tid[:])
	return hex.EncodeToString(tid[:])
}
