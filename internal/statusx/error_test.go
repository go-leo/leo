package statusx

import (
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	httpstatus "google.golang.org/genproto/googleapis/rpc/http"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net/http"
	"testing"
)

var testError = &Error{
	Cause: &Cause{
		Cause: &Cause_Message{
			Message: wrapperspb.String("object is nil"),
		},
	},
	Detail: &Detail{
		RequestInfo: &errdetails.RequestInfo{
			RequestId:   "58384717947134",
			ServingData: "hello",
		},
		BadRequest: &errdetails.BadRequest{},
	},
	HttpStatus: &httpstatus.HttpResponse{
		Status:  400,
		Reason:  "Bad Request",
		Headers: []*httpstatus.HttpHeader{},
		Body:    []byte("OK"),
	},
	GrpcStatus: &rpcstatus.Status{
		Code:    int32(codes.InvalidArgument),
		Message: "object is nil",
		Details: nil,
	},
}

func TestEncode(t *testing.T) {
	status, header, body := testError.Encode()
	assert.Equal(t, 400, status)
	assert.Equal(t, http.Header{
		kStatusCauseKey:  []string{`{"message":"object is nil"}`},
		kStatusDetailKey: []string{`{"badRequest":{},"requestInfo":{"requestId":"58384717947134","servingData":"hello"}}`},
		kStatusGrpcKey:   []string{`{"code":3,"message":"object is nil"}`},
	}, header)
	assert.EqualValues(t, []byte("OK"), body)
}

func TestDecode(t *testing.T) {
	x := new(Error)
	status := 400
	header := http.Header{
		kStatusCauseKey:  []string{`{"message":"object is nil"}`},
		kStatusDetailKey: []string{`{"badRequest":{},"requestInfo":{"requestId":"58384717947134","servingData":"hello"}}`},
		kStatusGrpcKey:   []string{`{"code":3,"message":"object is nil"}`},
	}
	body := []byte("OK")
	x.Decode(status, header, body)
	assert.True(t, proto.Equal(testError, x))
}
