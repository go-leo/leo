package statusx

import (
	"errors"
	"fmt"
	"github.com/go-leo/leo/v3/proto/leo/status"
	"github.com/go-leo/leo/v3/statusx/internal/statuspb"
	"github.com/go-leo/leo/v3/statusx/internal/util"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	httpstatus "google.golang.org/genproto/googleapis/rpc/http"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net/http"
	"time"
)

type Option func(st *sampleStatus)

func newStatus(code codes.Code, opts ...Option) *sampleStatus {
	statusCode := util.ToHttpStatusCode(code)
	st := &sampleStatus{
		err: &statuspb.Error{
			GrpcStatus: &rpcstatus.Status{
				Code:    int32(code),
				Message: "",
				Details: nil,
			},
			DetailInfo: &statuspb.DetailInfo{
				Identifier: &statuspb.Identifier{
					Value: fmt.Sprintf("%d-%d", code, statusCode),
				},
				ErrorInfo:           nil,
				RetryInfo:           nil,
				DebugInfo:           nil,
				QuotaFailure:        nil,
				PreconditionFailure: nil,
				BadRequest:          nil,
				RequestInfo:         nil,
				ResourceInfo:        nil,
				Help:                nil,
				LocalizedMessage:    nil,
				Header:              nil,
				Extra:               nil,
			},
		},
	}
	for _, opt := range opts {
		opt(st)
	}
	return st
}

// Identifier sets the identifier of the Status.
// This distinguish between two Status objects as being the same when
// both code and status are identical.
func Identifier(id string) Option {
	return func(st *sampleStatus) {
		st.err.DetailInfo.Identifier = &statuspb.Identifier{
			Value: id,
		}
	}
}

// Message sets the message of the Status.
func Message(format string, a ...any) Option {
	return func(st *sampleStatus) {
		st.err.GrpcStatus.Message = fmt.Sprintf(format, a...)
	}
}

// Headers sets the http header info.
func Headers(header http.Header) Option {
	return func(st *sampleStatus) {
		if st.err.DetailInfo.Header == nil {
			st.err.DetailInfo.Header = &statuspb.Header{}
		}
		for key, values := range header {
			for _, value := range values {
				item := &httpstatus.HttpHeader{
					Key:   key,
					Value: value,
				}
				st.err.DetailInfo.Header.Headers = append(st.err.DetailInfo.Header.Headers, item)
			}
		}
	}
}

// ErrorInfo sets the error info.
func ErrorInfo(reason string, domain string, metadata map[string]string) Option {
	return func(st *sampleStatus) {
		st.err.DetailInfo.ErrorInfo = &errdetails.ErrorInfo{
			Reason:   reason,
			Domain:   domain,
			Metadata: metadata,
		}
	}
}

// RetryInfo sets the retry info.
func RetryInfo(retryDelay time.Duration) Option {
	return func(st *sampleStatus) {
		st.err.DetailInfo.RetryInfo = &errdetails.RetryInfo{
			RetryDelay: durationpb.New(retryDelay),
		}
	}
}

// DebugInfo sets the debug info.
func DebugInfo(stackEntries []string, detail string) Option {
	return func(st *sampleStatus) {
		st.err.DetailInfo.DebugInfo = &errdetails.DebugInfo{
			StackEntries: stackEntries,
			Detail:       detail,
		}
	}
}

// QuotaFailure sets the quota failure info.
func QuotaFailure(violations []*errdetails.QuotaFailure_Violation) Option {
	return func(st *sampleStatus) {
		st.err.DetailInfo.QuotaFailure = &errdetails.QuotaFailure{
			Violations: violations,
		}
	}
}

// PreconditionFailure sets the precondition failure info.
func PreconditionFailure(violations []*errdetails.PreconditionFailure_Violation) Option {
	return func(st *sampleStatus) {
		st.err.DetailInfo.PreconditionFailure = &errdetails.PreconditionFailure{
			Violations: violations,
		}
	}
}

// BadRequest sets the bad request info.
func BadRequest(violations []*errdetails.BadRequest_FieldViolation) Option {
	return func(st *sampleStatus) {
		st.err.DetailInfo.BadRequest = &errdetails.BadRequest{
			FieldViolations: violations,
		}
	}
}

// RequestInfo sets the request info.
func RequestInfo(requestId string, servingData string) Option {
	return func(st *sampleStatus) {
		st.err.DetailInfo.RequestInfo = &errdetails.RequestInfo{
			RequestId:   requestId,
			ServingData: servingData,
		}
	}
}

// ResourceInfo sets the resource info.
func ResourceInfo(resourceType string, resourceName string, owner string, description string) Option {
	return func(st *sampleStatus) {
		st.err.DetailInfo.ResourceInfo = &errdetails.ResourceInfo{
			ResourceType: resourceType,
			ResourceName: resourceName,
			Owner:        owner,
			Description:  description,
		}
	}
}

// Help sets the help info.
func Help(links []*errdetails.Help_Link) Option {
	return func(st *sampleStatus) {
		st.err.DetailInfo.Help = &errdetails.Help{
			Links: links,
		}
	}
}

// LocalizedMessage sets the localized message info.
func LocalizedMessage(locale string, message string) Option {
	return func(st *sampleStatus) {
		st.err.DetailInfo.LocalizedMessage = &errdetails.LocalizedMessage{
			Locale:  locale,
			Message: message,
		}
	}
}

// Extra sets the extra info.
func Extra(extra proto.Message) Option {
	return func(st *sampleStatus) {
		switch item := extra.(type) {
		case *wrapperspb.StringValue:
			Message(item.GetValue())(st)
		case *httpstatus.HttpHeader:
			if st.err.DetailInfo.Header == nil {
				st.err.DetailInfo.Header = &statuspb.Header{}
			}
			st.err.DetailInfo.Header.Headers = append(st.err.DetailInfo.Header.Headers, item)
		case *errdetails.ErrorInfo:
			ErrorInfo(item.GetReason(), item.GetDomain(), item.GetMetadata())(st)
		case *errdetails.RetryInfo:
			RetryInfo(item.GetRetryDelay().AsDuration())(st)
		case *errdetails.DebugInfo:
			DebugInfo(item.GetStackEntries(), item.GetDetail())(st)
		case *errdetails.QuotaFailure:
			QuotaFailure(item.GetViolations())(st)
		case *errdetails.PreconditionFailure:
			PreconditionFailure(item.GetViolations())(st)
		case *errdetails.BadRequest:
			BadRequest(item.GetFieldViolations())(st)
		case *errdetails.RequestInfo:
			RequestInfo(item.GetRequestId(), item.GetServingData())(st)
		case *errdetails.ResourceInfo:
			ResourceInfo(item.GetResourceType(), item.GetResourceName(), item.GetOwner(), item.GetDescription())(st)
		case *errdetails.Help:
			Help(item.GetLinks())(st)
		case *errdetails.LocalizedMessage:
			LocalizedMessage(item.GetLocale(), item.GetMessage())(st)
		case *httpstatus.HttpResponse:
			panic(errors.New("status: unsupported HttpResponse"))
		case *rpcstatus.Status:
			panic(errors.New("status: unsupported Status"))
		case *statuspb.DetailInfo:
			panic(errors.New("status: unsupported DetailInfo"))
		case *status.HttpBody:
			panic(errors.New("status: unsupported HttpBody"))
		default:
			value, err := anypb.New(item)
			if err != nil {
				panic(err)
			}
			st.err.DetailInfo.Extra = value
		}
	}
}
