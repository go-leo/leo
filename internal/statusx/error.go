package statusx

import (
	"errors"
	"fmt"
	"github.com/go-leo/gox/errorx"
	"github.com/go-leo/gox/protox"
	httpstatus "github.com/go-leo/leo/v3/statusx/http"
	"golang.org/x/exp/slices"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

// Error is a error
func (x *Error) Error() string {
	return fmt.Sprintf("rpc error: code = %s desc = %s", codes.Code(x.GetGrpcStatus().GetCode()), x.GetGrpcStatus().GetMessage())
}

// GRPCStatus returns the gRPC Status.
func (x *Error) GRPCStatus() *grpcstatus.Status {
	// append http status to grpc status details
	oldDetails := x.GetGrpcStatus().GetDetails()
	mewDetails := make([]*anypb.Any, 0, len(oldDetails)+1)
	mewDetails = append(mewDetails, oldDetails...)
	mewDetails = append(mewDetails, errorx.Ignore(anypb.New(x.HttpStatus)))
	// return grpc status
	return grpcstatus.FromProto(&rpcstatus.Status{
		Code:    x.GetGrpcStatus().GetCode(),
		Message: x.GetGrpcStatus().GetMessage(),
		Details: mewDetails,
	})
}

// HTTPStatus returns the HTTP status.
func (x *Error) HTTPStatus() *httpstatus.Status {
	return protox.Clone(x.GetHttpStatus())
}

// Proto return the gRPC and HTTP status protocol buffers.
func (x *Error) Proto() (*rpcstatus.Status, *httpstatus.Status) {
	return protox.Clone(x.GetGrpcStatus()), protox.Clone(x.GetHttpStatus())
}

// Is implements future errors.Is functionality.
// It determines if an error (target) is equivalent to the current Error instance (x) by comparing
// their gRPC status and HTTP status.
func (x *Error) Is(target error) bool {
	var targetErr *Error
	if !errors.As(target, &targetErr) {
		return false
	}
	return proto.Equal(x.GetGrpcStatus(), targetErr.GetGrpcStatus()) &&
		proto.Equal(x.GetHttpStatus(), targetErr.GetHttpStatus())
}

// Equals compares the status code and http status code of two errors.
func (x *Error) Equals(target error) bool {
	var targetErr *Error
	if !errors.As(target, &targetErr) {
		return false
	}
	return x.GetGrpcStatus().GetCode() == targetErr.GetGrpcStatus().GetCode() &&
		x.GetHttpStatus().GetCode() == targetErr.GetHttpStatus().GetCode()
}

// WithMessage sets the message of the gRPC status.
func (x *Error) WithMessage(msg string) {
	x.GrpcStatus.Message = msg
}

// WithDetails sets the details of the gRPC status.
func (x *Error) WithDetails(details ...proto.Message) {
	for _, detail := range details {
		anyDetail, err := anypb.New(detail)
		if err != nil {
			continue
		}
		x.GrpcStatus.Details = append(x.GrpcStatus.Details, anyDetail)
	}
}

func (x *Error) Details() []proto.Message {
	details := x.GetGrpcStatus().GetDetails()
	res := make([]proto.Message, 0, len(details))
	for _, anyDetail := range details {
		detail, err := anyDetail.UnmarshalNew()
		if err != nil {
			continue
		}
		res = append(res, detail)
	}
	return res
}

// WithHttpHeader sets the http header info.
func (x *Error) WithHttpHeader(infos ...*httpstatus.Header) {
	for _, info := range infos {
		x.HttpStatus.Headers = append(x.HttpStatus.Headers, info)
	}
}

func (x *Error) HttpHeader() []*httpstatus.Header {
	return slices.Clone(x.GetHttpStatus().GetHeaders())
}

func (x *Error) WithHttpBody(body *anypb.Any) {
	x.HttpStatus.Body = body
}

func (x *Error) HttpBody() proto.Message {
	body, _ := x.GetHttpStatus().GetBody().UnmarshalNew()
	return body
}
