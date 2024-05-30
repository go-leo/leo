package statusx

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/go-leo/gox/slicex"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	apistatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

var _ ErrorAPI = (*Error)(nil)

func (x *Error) Error() string {
	return fmt.Sprintf("rpc error: code = %s desc = %s", x.code(), x.message())
}

func (x *Error) Froze() ErrorAPI {
	return proto.Clone(x).(*Error)
}

func (x *Error) Wrap(err error) ErrorAPI {
	if errPb, ok := err.(proto.Message); ok {
		return x.WithDetails(errPb)
	}
	return x.WithDetails(&errdetails.ErrorInfo{Reason: err.Error()})
}

func (x *Error) Unwrap() error {
	for _, detail := range x.Details() {
		if err, ok := detail.(error); ok {
			return err
		}
	}
	for _, detail := range x.Details() {
		if errorInfo, ok := detail.(*errdetails.ErrorInfo); ok {
			return errors.New(errorInfo.GetReason())
		}
	}
	return nil
}

func (x *Error) Is(target error) bool {
	var tse *Error
	ok := errors.As(target, &tse)
	if !ok {
		return false
	}
	return proto.Equal(x.GetGrpcStatus(), tse.GetHttpStatus()) && proto.Equal(x.GetHttpStatus(), tse.GetHttpStatus())
}

func (x *Error) HTTPStatus() int {
	return int(x.GetHttpStatus().GetStatus())
}

func (x *Error) GRPCStatus() *grpcstatus.Status {
	grpcStatus := x.GetGrpcStatus()
	for _, detail := range grpcStatus.GetDetails() {
		if detail.MessageIs((*HttpStatus)(nil)) {
			return grpcstatus.FromProto(grpcStatus)
		}
	}
	httpStatus, err := anypb.New(x.GetHttpStatus())
	if err != nil {
		panic(err)
	}
	grpcStatus.Details = append(grpcStatus.Details, httpStatus)
	return grpcstatus.FromProto(grpcStatus)
}

func (x *Error) WithDetails(messages ...proto.Message) ErrorAPI {
	//if x.code() == codes.OK {
	//	return x
	//}
	details := make([]*anypb.Any, 0, len(x.GetGrpcStatus().GetDetails())+len(messages))
	for _, detail := range x.GetGrpcStatus().GetDetails() {
		details = append(details, &anypb.Any{TypeUrl: detail.GetTypeUrl(), Value: bytes.Clone(detail.GetValue())})
	}
	for _, message := range messages {
		detail, err := anypb.New(message)
		if err != nil {
			panic(err)
		}
		details = append(details, detail)
	}
	return &Error{
		GrpcStatus: &apistatus.Status{
			Code:    x.GetGrpcStatus().GetCode(),
			Message: x.GetGrpcStatus().GetMessage(),
			Details: details,
		},
		HttpStatus: &HttpStatus{
			Status: x.GetHttpStatus().GetStatus(),
			Reason: x.GetHttpStatus().GetReason(),
		},
	}
}

func (x *Error) Details() []proto.Message {
	if x == nil || x.GetHttpStatus() == nil {
		return nil
	}
	messages := make([]proto.Message, 0, len(x.GetGrpcStatus().GetDetails()))
	for _, detail := range x.GetGrpcStatus().GetDetails() {
		message, err := detail.UnmarshalNew()
		if err != nil {
			panic(err)
		}
		messages = append(messages, message)
	}
	return messages
}

func (x *Error) WithErrorInfo(infos ...*errdetails.ErrorInfo) ErrorAPI {
	messages := slicex.Map[[]*errdetails.ErrorInfo, []proto.Message](infos, func(i int, info *errdetails.ErrorInfo) proto.Message { return info })
	return x.WithDetails(messages...)
}

func (x *Error) ErrorInfo() []*errdetails.ErrorInfo {
	var infos []*errdetails.ErrorInfo
	for _, message := range x.Details() {
		if info, ok := message.(*errdetails.ErrorInfo); ok {
			infos = append(infos, info)
		}
	}
	return infos
}

func (x *Error) WithRetryInfo(infos ...*errdetails.RetryInfo) ErrorAPI {
	messages := slicex.Map[[]*errdetails.RetryInfo, []proto.Message](infos, func(i int, info *errdetails.RetryInfo) proto.Message { return info })
	return x.WithDetails(messages...)
}

func (x *Error) RetryInfo() []*errdetails.RetryInfo {
	var infos []*errdetails.RetryInfo
	for _, message := range x.Details() {
		if info, ok := message.(*errdetails.RetryInfo); ok {
			infos = append(infos, info)
		}
	}
	return infos
}

func (x *Error) WithDebugInfo(infos ...*errdetails.DebugInfo) ErrorAPI {
	messages := slicex.Map[[]*errdetails.DebugInfo, []proto.Message](infos, func(i int, info *errdetails.DebugInfo) proto.Message { return info })
	return x.WithDetails(messages...)
}

func (x *Error) DebugInfo() []*errdetails.DebugInfo {
	var infos []*errdetails.DebugInfo
	for _, message := range x.Details() {
		if info, ok := message.(*errdetails.DebugInfo); ok {
			infos = append(infos, info)
		}
	}
	return infos
}

func (x *Error) WithQuotaFailure(infos ...*errdetails.QuotaFailure) ErrorAPI {
	messages := slicex.Map[[]*errdetails.QuotaFailure, []proto.Message](infos, func(i int, info *errdetails.QuotaFailure) proto.Message { return info })
	return x.WithDetails(messages...)
}

func (x *Error) QuotaFailure() []*errdetails.QuotaFailure {
	var infos []*errdetails.QuotaFailure
	for _, message := range x.Details() {
		if info, ok := message.(*errdetails.QuotaFailure); ok {
			infos = append(infos, info)
		}
	}
	return infos
}

func (x *Error) WithPreconditionFailure(infos ...*errdetails.PreconditionFailure) ErrorAPI {
	messages := slicex.Map[[]*errdetails.PreconditionFailure, []proto.Message](infos, func(i int, info *errdetails.PreconditionFailure) proto.Message { return info })
	return x.WithDetails(messages...)
}

func (x *Error) PreconditionFailure() []*errdetails.PreconditionFailure {
	var infos []*errdetails.PreconditionFailure
	for _, message := range x.Details() {
		if info, ok := message.(*errdetails.PreconditionFailure); ok {
			infos = append(infos, info)
		}
	}
	return infos
}

func (x *Error) WithBadRequest(infos ...*errdetails.BadRequest) ErrorAPI {
	messages := slicex.Map[[]*errdetails.BadRequest, []proto.Message](infos, func(i int, info *errdetails.BadRequest) proto.Message { return info })
	return x.WithDetails(messages...)
}

func (x *Error) BadRequest() []*errdetails.BadRequest {
	var infos []*errdetails.BadRequest
	for _, message := range x.Details() {
		if info, ok := message.(*errdetails.BadRequest); ok {
			infos = append(infos, info)
		}
	}
	return infos
}

func (x *Error) WithRequestInfo(infos ...*errdetails.RequestInfo) ErrorAPI {
	messages := slicex.Map[[]*errdetails.RequestInfo, []proto.Message](infos, func(i int, info *errdetails.RequestInfo) proto.Message { return info })
	return x.WithDetails(messages...)
}

func (x *Error) RequestInfo() []*errdetails.RequestInfo {
	var infos []*errdetails.RequestInfo
	for _, message := range x.Details() {
		if info, ok := message.(*errdetails.RequestInfo); ok {
			infos = append(infos, info)
		}
	}
	return infos
}

func (x *Error) WithResourceInfo(infos ...*errdetails.ResourceInfo) ErrorAPI {
	messages := slicex.Map[[]*errdetails.ResourceInfo, []proto.Message](infos, func(i int, info *errdetails.ResourceInfo) proto.Message { return info })
	return x.WithDetails(messages...)
}

func (x *Error) ResourceInfo() []*errdetails.ResourceInfo {
	var infos []*errdetails.ResourceInfo
	for _, message := range x.Details() {
		if info, ok := message.(*errdetails.ResourceInfo); ok {
			infos = append(infos, info)
		}
	}
	return infos
}

func (x *Error) WithHelp(infos ...*errdetails.Help) ErrorAPI {
	messages := slicex.Map[[]*errdetails.Help, []proto.Message](infos, func(i int, info *errdetails.Help) proto.Message { return info })
	return x.WithDetails(messages...)
}

func (x *Error) Help() []*errdetails.Help {
	var infos []*errdetails.Help
	for _, message := range x.Details() {
		if info, ok := message.(*errdetails.Help); ok {
			infos = append(infos, info)
		}
	}
	return infos
}

func (x *Error) WithLocalizedMessage(infos ...*errdetails.LocalizedMessage) ErrorAPI {
	messages := slicex.Map[[]*errdetails.LocalizedMessage, []proto.Message](infos, func(i int, info *errdetails.LocalizedMessage) proto.Message { return info })
	return x.WithDetails(messages...)
}

func (x *Error) LocalizedMessage() []*errdetails.LocalizedMessage {
	var infos []*errdetails.LocalizedMessage
	for _, message := range x.Details() {
		if info, ok := message.(*errdetails.LocalizedMessage); ok {
			infos = append(infos, info)
		}
	}
	return infos
}

// Code returns the status code contained in s.
func (x *Error) code() codes.Code {
	if x == nil || x.GetGrpcStatus() == nil {
		return codes.OK
	}
	return codes.Code(x.GetGrpcStatus().GetCode())
}

// Message returns the message contained in s.
func (x *Error) message() string {
	if x == nil || x.GetGrpcStatus() == nil {
		return ""
	}
	return x.GetGrpcStatus().GetMessage()
}
