package statusx

import (
	"github.com/go-leo/gox/protox"
	httpstatus "github.com/go-leo/leo/v3/statusx/http"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/protoadapt"
)

var _ Status = (*status)(nil)

type status struct {
	s *grpcstatus.Status
}

func (s *status) Clone() Status {
	return &status{s: grpcstatus.FromProto(s.s.Proto())}
}

func (s *status) String() string {
	return s.s.String()
}

func (s *status) GrpcProto() *rpcstatus.Status {
	return s.s.Proto()
}

func (s *status) HttpProto() *httpstatus.Status {
	for _, message := range s.Details() {
		if httpStatus, ok := message.(*httpstatus.Status); ok {
			return proto.Clone(httpStatus).(*httpstatus.Status)
		}
	}
	return nil
}

func (s *status) Message() string {
	return s.s.Message()
}

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

func (s *status) Err() error {
	return s.s.Err()
}

func (s *status) WithErr(err error) Status {
	if errPb, ok := err.(proto.Message); ok {
		return s.WithDetails(errPb)
	}
	return s.WithDetails(&errdetails.ErrorInfo{Reason: err.Error()})
}

func (s *status) WithDetails(details ...proto.Message) Status {
	if s.s.Code() == codes.OK {
		return s
	}
	v1Details := make([]protoadapt.MessageV1, 0, len(details))
	for _, detail := range details {
		v1Details = append(v1Details, protoadapt.MessageV1Of(detail))
	}
	st, err := s.s.WithDetails(v1Details...)
	if err != nil {
		panic(err)
	}
	return &status{s: st}
}

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
