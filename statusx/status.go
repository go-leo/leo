package statusx

import (
	"errors"
	"github.com/go-leo/gox/protox"
	httpstatus "github.com/go-leo/leo/v3/statusx/http"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/protoadapt"
	"google.golang.org/protobuf/types/known/anypb"
)

var _ Status = (*status)(nil)

type status struct {
	s *grpcstatus.Status
}

// Clone returns a cloned copy of the status object by creating a new status with a copied grpcstatus.Status.
func (s *status) Clone() Status {
	return &status{s: grpcstatus.FromProto(s.s.Proto())}
}

// Is compares the current status's gRPC and HTTP status codes with another target status.
// returns true if both codes are equal.
func (s *status) Is(target Status) bool {
	return s.GrpcProto().GetCode() == target.GrpcProto().GetCode() && s.HttpProto().GetStatus() == s.HttpProto().GetStatus()
}

// Equals checks if the current status is equal to the target status by comparing their gRPC protocol buffer
// representations using proto.Equal.
func (s *status) Equals(target Status) bool {
	return proto.Equal(s.GrpcProto(), target.GrpcProto())
}

// String returns a string representation of the status using the grpcstatus.Status's String() method.
func (s *status) String() string {
	return s.s.String()
}

// GrpcProto Return pointers to the gRPC status protocol buffers respectively.
func (s *status) GrpcProto() *rpcstatus.Status {
	messages := s.Details()
	details := make([]*anypb.Any, 0, len(messages))
	for _, message := range messages {
		if _, ok := message.(*httpstatus.Status); ok {
			continue
		}
		detail, err := anypb.New(message)
		if err != nil {
			panic(err)
		}
		details = append(details, detail)
	}
	return &rpcstatus.Status{
		Code:    int32(s.s.Code()),
		Message: s.s.Message(),
		Details: details,
	}
}

// HttpProto return pointers to the HTTP status protocol buffers respectively.
func (s *status) HttpProto() *httpstatus.Status {
	for _, message := range s.Details() {
		if httpStatus, ok := message.(*httpstatus.Status); ok {
			return proto.Clone(httpStatus).(*httpstatus.Status)
		}
	}
	return nil
}

// Message retrieves the message associated with the status.
func (s *status) Message() string {
	return s.s.Message()
}

// WithMessage creates a new status with an updated message while preserving other information.
func (s *status) WithMessage(msg string, args ...any) Status {
	details := s.Details()
	v1Details := make([]protoadapt.MessageV1, 0, len(details))
	for _, detail := range details {
		v1Details = append(v1Details, protoadapt.MessageV1Of(detail))
	}
	st := grpcstatus.Newf(s.s.Code(), msg, args...)
	st, err := st.WithDetails(v1Details...)
	if err != nil {
		panic(err)
	}
	return &status{s: st}
}

// Err returns the error represented by the status.
func (s *status) Err() error {
	return s.s.Err()
}

// WithErr updates the status with a new error, either as a proto.Message or a simple error converted to errdetails.ErrorInfo.
func (s *status) WithErr(err error) Status {
	if errPb, ok := err.(proto.Message); ok {
		return s.WithDetails(errPb)
	}
	return s.WithDetails(&errdetails.ErrorInfo{Reason: err.Error()})
}

// WithDetails adds additional details to the status as protocol buffer messages.
func (s *status) WithDetails(details ...proto.Message) Status {
	if s.s.Code() == codes.OK {
		return s
	}
	v1Details := make([]protoadapt.MessageV1, 0, len(details))
	for _, detail := range details {
		if _, ok := detail.(*httpstatus.Status); ok {
			panic(errors.New("http status can not be added to grpc status"))
		}
		v1Details = append(v1Details, protoadapt.MessageV1Of(detail))
	}
	st, err := s.s.WithDetails(v1Details...)
	if err != nil {
		panic(err)
	}
	return &status{s: st}
}

// Details return additional details from the status
func (s *status) Details() []proto.Message {
	details := s.s.Details()
	detailPbs := make([]proto.Message, 0, len(details))
	for _, detail := range details {
		detailPbs = append(detailPbs, detail.(proto.Message))
	}
	return detailPbs
}

func (s *status) WithErrorInfo(infos ...*errdetails.ErrorInfo) Status {
	return s.WithDetails(protox.MessageSlice(infos)...)
}

func (s *status) ErrorInfo() []*errdetails.ErrorInfo {
	return protox.ProtoSlice[[]*errdetails.ErrorInfo](s.Details())
}

func (s *status) WithDebugInfo(infos ...*errdetails.DebugInfo) Status {
	return s.WithDetails(protox.MessageSlice(infos)...)
}

func (s *status) DebugInfo() []*errdetails.DebugInfo {
	return protox.ProtoSlice[[]*errdetails.DebugInfo](s.Details())
}

func (s *status) WithLocalizedMessage(infos ...*errdetails.LocalizedMessage) Status {
	return s.WithDetails(protox.MessageSlice(infos)...)
}

func (s *status) LocalizedMessage() []*errdetails.LocalizedMessage {
	return protox.ProtoSlice[[]*errdetails.LocalizedMessage](s.Details())
}

func (s *status) WithBadRequest(infos ...*errdetails.BadRequest) Status {
	return s.WithDetails(protox.MessageSlice(infos)...)
}

func (s *status) BadRequest() []*errdetails.BadRequest {
	return protox.ProtoSlice[[]*errdetails.BadRequest](s.Details())
}

func (s *status) WithPreconditionFailure(infos ...*errdetails.PreconditionFailure) Status {
	return s.WithDetails(protox.MessageSlice(infos)...)
}

func (s *status) PreconditionFailure() []*errdetails.PreconditionFailure {
	return protox.ProtoSlice[[]*errdetails.PreconditionFailure](s.Details())
}

func (s *status) WithRetryInfo(infos ...*errdetails.RetryInfo) Status {
	return s.WithDetails(protox.MessageSlice(infos)...)
}

func (s *status) RetryInfo() []*errdetails.RetryInfo {
	return protox.ProtoSlice[[]*errdetails.RetryInfo](s.Details())
}

func (s *status) WithQuotaFailure(infos ...*errdetails.QuotaFailure) Status {
	return s.WithDetails(protox.MessageSlice(infos)...)
}

func (s *status) QuotaFailure() []*errdetails.QuotaFailure {
	return protox.ProtoSlice[[]*errdetails.QuotaFailure](s.Details())
}

func (s *status) WithResourceInfo(infos ...*errdetails.ResourceInfo) Status {
	return s.WithDetails(protox.MessageSlice(infos)...)
}

func (s *status) ResourceInfo() []*errdetails.ResourceInfo {
	return protox.ProtoSlice[[]*errdetails.ResourceInfo](s.Details())
}

func (s *status) WithHelp(infos ...*errdetails.Help) Status {
	return s.WithDetails(protox.MessageSlice(infos)...)
}

func (s *status) Help() []*errdetails.Help {
	return protox.ProtoSlice[[]*errdetails.Help](s.Details())
}

func (s *status) WithRequestInfo(infos ...*errdetails.RequestInfo) Status {
	return s.WithDetails(protox.MessageSlice(infos)...)
}

func (s *status) RequestInfo() []*errdetails.RequestInfo {
	return protox.ProtoSlice[[]*errdetails.RequestInfo](s.Details())
}

func newStatus(grpcStatus *grpcstatus.Status, httpStatus *httpstatus.Status) Status {
	details, _ := grpcStatus.WithDetails(httpStatus)
	return &status{s: details}
}
