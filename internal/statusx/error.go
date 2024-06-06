package statusx

import (
	"errors"
	"fmt"
	"github.com/go-leo/gox/errorx"
	"github.com/go-leo/gox/protox"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

// Error wraps a pointer of a status proto. It implements error and Status,
// and a nil *Error should never be returned by this package.
func (x *Error) Error() string {
	return fmt.Sprintf("rpc error: code = %s desc = %s", codes.Code(x.GetGrpcStatus().GetCode()), x.GetGrpcStatus().GetMessage())
}

// GRPCStatus returns the Status represented by se.
func (x *Error) GRPCStatus() *grpcstatus.Status {
	grpcProto := protox.Clone(x.GrpcStatus)
	grpcProto.Details = append(grpcProto.Details, errorx.Ignore(anypb.New(x.HttpStatus)))
	return grpcstatus.FromProto(grpcProto)
}

// Is implements future errors.Is functionality.
func (x *Error) Is(target error) bool {
	var tse *Error
	if !errors.As(target, &tse) {
		return false
	}
	return proto.Equal(x.GetGrpcStatus(), tse.GetGrpcStatus()) && proto.Equal(x.GetHttpStatus(), tse.GetHttpStatus())
}
