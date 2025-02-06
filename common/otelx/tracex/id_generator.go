package tracex

import (
	"context"
	"crypto/md5"
	"encoding/binary"
	"math/rand"
	"sync"
	"unsafe"

	crand "crypto/rand"

	"github.com/go-leo/gox/stringx"
	"github.com/go-leo/leo/middleware/requestid"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type ctxOrRandomIDGenerator struct {
	sync.Mutex
	randSource *rand.Rand
}

var _ sdktrace.IDGenerator = &ctxOrRandomIDGenerator{}

// NewSpanID returns a non-zero span ID from a randomly-chosen sequence.
func (gen *ctxOrRandomIDGenerator) NewSpanID(ctx context.Context, _ trace.TraceID) trace.SpanID {
	gen.Lock()
	defer gen.Unlock()
	sid := trace.SpanID{}
	_, _ = gen.randSource.Read(sid[:])
	return sid
}

// NewIDs returns a non-zero trace ID and a non-zero span ID from ctx or a
// randomly-chosen sequence.
func (gen *ctxOrRandomIDGenerator) NewIDs(ctx context.Context) (trace.TraceID, trace.SpanID) {
	gen.Lock()
	defer gen.Unlock()
	tid := trace.TraceID{}
	// 1. from context
	requestID, _ := requestid.FromContext(ctx)
	if stringx.IsNotBlank(requestID) {
		tid = genTid(requestID)
	}
	// 2. from grpc incoming
	requestID, _ = requestid.FromIncoming(ctx)
	if stringx.IsNotBlank(requestID) {
		tid = genTid(requestID)
	}
	// 3. from http header
	requestID, _ = requestid.FromHeader(ctx)
	if stringx.IsNotBlank(requestID) {
		tid = genTid(requestID)
	}
	// 4. generate
	_, _ = gen.randSource.Read(tid[:])
	sid := trace.SpanID{}
	_, _ = gen.randSource.Read(sid[:])
	return tid, sid
}

func genTid(requestID string) trace.TraceID {
	return md5.Sum(*(*[]byte)(unsafe.Pointer(&requestID)))
}

func NewIDGenerator() sdktrace.IDGenerator {
	gen := &ctxOrRandomIDGenerator{}
	var rngSeed int64
	_ = binary.Read(crand.Reader, binary.LittleEndian, &rngSeed)
	gen.randSource = rand.New(rand.NewSource(rngSeed))
	return gen
}
