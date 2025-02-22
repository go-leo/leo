package statuspb

import (
	"github.com/go-leo/gox/slicex"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	httpstatus "google.golang.org/genproto/googleapis/rpc/http"
	"google.golang.org/protobuf/types/known/anypb"
)

func (err *Error) HttpDetails() []*anypb.Any {
	return err.GetDetailInfo().ToAnySlice()
}

func (err *Error) ToGrpcDetails() []*anypb.Any {
	r := err.GetDetailInfo().ToAnySlice()
	// add http status info to details
	if info := err.GetHttpStatus(); info != nil {
		infoAny, err := anypb.New(info)
		if err != nil {
			panic(err)
		}
		r = append(r, infoAny)
	}
	return r
}

func (err *Error) FromGrpcDetails(details []*anypb.Any) (*DetailInfo, *httpstatus.HttpResponse) {
	if err == nil {
		return nil, nil
	}
	var detailInfo *DetailInfo
	var httpStatus *httpstatus.HttpResponse
	detailInfo = err.GetDetailInfo().FromAnySlice(details)
	for i, value := range detailInfo.GetDetails() {
		if value.MessageIs(&httpstatus.HttpResponse{}) {
			httpStatus = new(httpstatus.HttpResponse)
			if e := value.UnmarshalTo(httpStatus); e != nil {
				panic(e)
			}
			detailInfo.Details = slicex.RemoveAt(detailInfo.Details, i)
			break
		}
	}
	return detailInfo, httpStatus
}

func (detailInfo *DetailInfo) FromAnySlice(details []*anypb.Any) *DetailInfo {
	if detailInfo == nil {
		detailInfo = &DetailInfo{}
	}
	for _, value := range details {
		switch {
		case value.MessageIs(&Identifier{}):
			detailInfo.Identifier = new(Identifier)
			err := value.UnmarshalTo(detailInfo.Identifier)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(&errdetails.ErrorInfo{}):
			detailInfo.ErrorInfo = new(errdetails.ErrorInfo)
			err := value.UnmarshalTo(detailInfo.ErrorInfo)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(&errdetails.RetryInfo{}):
			detailInfo.RetryInfo = new(errdetails.RetryInfo)
			err := value.UnmarshalTo(detailInfo.RetryInfo)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(&errdetails.DebugInfo{}):
			detailInfo.DebugInfo = new(errdetails.DebugInfo)
			err := value.UnmarshalTo(detailInfo.DebugInfo)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(&errdetails.QuotaFailure{}):
			detailInfo.QuotaFailure = new(errdetails.QuotaFailure)
			err := value.UnmarshalTo(detailInfo.QuotaFailure)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(&errdetails.PreconditionFailure{}):
			detailInfo.PreconditionFailure = new(errdetails.PreconditionFailure)
			err := value.UnmarshalTo(detailInfo.PreconditionFailure)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(&errdetails.BadRequest{}):
			detailInfo.BadRequest = new(errdetails.BadRequest)
			err := value.UnmarshalTo(detailInfo.BadRequest)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(&errdetails.RequestInfo{}):
			detailInfo.RequestInfo = new(errdetails.RequestInfo)
			err := value.UnmarshalTo(detailInfo.RequestInfo)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(&errdetails.ResourceInfo{}):
			detailInfo.ResourceInfo = new(errdetails.ResourceInfo)
			err := value.UnmarshalTo(detailInfo.ResourceInfo)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(&errdetails.Help{}):
			detailInfo.Help = new(errdetails.Help)
			err := value.UnmarshalTo(detailInfo.Help)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(&errdetails.LocalizedMessage{}):
			detailInfo.LocalizedMessage = new(errdetails.LocalizedMessage)
			err := value.UnmarshalTo(detailInfo.LocalizedMessage)
			if err != nil {
				panic(err)
			}
		default:
			detailInfo.Details = append(detailInfo.Details, value)
		}
	}
	return detailInfo
}

func (detailInfo *DetailInfo) ToAnySlice() []*anypb.Any {
	if detailInfo == nil {
		return nil
	}
	var details []*anypb.Any
	if info := detailInfo.GetIdentifier(); info != nil {
		infoAny, e := anypb.New(info)
		if e != nil {
			panic(e)
		}
		details = append(details, infoAny)
	}
	if info := detailInfo.GetErrorInfo(); info != nil {
		infoAny, e := anypb.New(info)
		if e != nil {
			panic(e)
		}
		details = append(details, infoAny)
	}
	if info := detailInfo.GetRetryInfo(); info != nil {
		infoAny, e := anypb.New(info)
		if e != nil {
			panic(e)
		}
		details = append(details, infoAny)
	}
	if info := detailInfo.GetDebugInfo(); info != nil {
		infoAny, e := anypb.New(info)
		if e != nil {
			panic(e)
		}
		details = append(details, infoAny)
	}
	if info := detailInfo.GetQuotaFailure(); info != nil {
		infoAny, e := anypb.New(info)
		if e != nil {
			panic(e)
		}
		details = append(details, infoAny)
	}
	if info := detailInfo.GetPreconditionFailure(); info != nil {
		infoAny, e := anypb.New(info)
		if e != nil {
			panic(e)
		}
		details = append(details, infoAny)
	}
	if info := detailInfo.GetBadRequest(); info != nil {
		infoAny, e := anypb.New(info)
		if e != nil {
			panic(e)
		}
		details = append(details, infoAny)
	}
	if info := detailInfo.GetRequestInfo(); info != nil {
		infoAny, e := anypb.New(info)
		if e != nil {
			panic(e)
		}
		details = append(details, infoAny)
	}
	if info := detailInfo.GetResourceInfo(); info != nil {
		infoAny, e := anypb.New(info)
		if e != nil {
			panic(e)
		}
		details = append(details, infoAny)
	}
	if info := detailInfo.GetHelp(); info != nil {
		infoAny, e := anypb.New(info)
		if e != nil {
			panic(e)
		}
		details = append(details, infoAny)
	}
	if info := detailInfo.GetLocalizedMessage(); info != nil {
		infoAny, e := anypb.New(info)
		if e != nil {
			panic(e)
		}
		details = append(details, infoAny)
	}
	for _, infoAny := range detailInfo.GetDetails() {
		details = append(details, infoAny)
	}
	return details
}
