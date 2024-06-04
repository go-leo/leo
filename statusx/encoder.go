package statusx

import (
	"context"
	httpstatus "github.com/go-leo/leo/v3/statusx/http"
	grpcstatus "google.golang.org/grpc/status"
)

type Encoder interface {
	Encode(ctx context.Context, grpcStatus *grpcstatus.Status) (*httpstatus.Status, []*httpstatus.Header, []byte)
}

type Decoder interface {
	Decode(ctx context.Context, httpProto *httpstatus.Status, header []*httpstatus.Header, body []byte) *grpcstatus.Status
}
