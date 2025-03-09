package statuspb

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

var (
	_Identifier          = &Identifier{}
	_ErrorInfo           = &errdetails.ErrorInfo{}
	_RetryInfo           = &errdetails.RetryInfo{}
	_DebugInfo           = &errdetails.DebugInfo{}
	_QuotaFailure        = &errdetails.QuotaFailure{}
	_PreconditionFailure = &errdetails.PreconditionFailure{}
	_BadRequest          = &errdetails.BadRequest{}
	_RequestInfo         = &errdetails.RequestInfo{}
	_ResourceInfo        = &errdetails.ResourceInfo{}
	_Help                = &errdetails.Help{}
	_LocalizedMessage    = &errdetails.LocalizedMessage{}
	_Header              = &Header{}
)

func ToHttpDetails(detailInfo *DetailInfo) []*anypb.Any {
	if detailInfo == nil {
		return nil
	}
	var details []*anypb.Any
	if info := detailInfo.GetIdentifier(); info != nil {
		infoAny, err := anypb.New(info)
		if err != nil {
			panic(err)
		}
		details = append(details, infoAny)
	}
	if info := detailInfo.GetErrorInfo(); info != nil {
		infoAny, err := anypb.New(info)
		if err != nil {
			panic(err)
		}
		details = append(details, infoAny)
	}
	if info := detailInfo.GetRetryInfo(); info != nil {
		infoAny, err := anypb.New(info)
		if err != nil {
			panic(err)
		}
		details = append(details, infoAny)
	}
	if info := detailInfo.GetDebugInfo(); info != nil {
		infoAny, err := anypb.New(info)
		if err != nil {
			panic(err)
		}
		details = append(details, infoAny)
	}
	if info := detailInfo.GetQuotaFailure(); info != nil {
		infoAny, err := anypb.New(info)
		if err != nil {
			panic(err)
		}
		details = append(details, infoAny)
	}
	if info := detailInfo.GetPreconditionFailure(); info != nil {
		infoAny, err := anypb.New(info)
		if err != nil {
			panic(err)
		}
		details = append(details, infoAny)
	}
	if info := detailInfo.GetBadRequest(); info != nil {
		infoAny, err := anypb.New(info)
		if err != nil {
			panic(err)
		}
		details = append(details, infoAny)
	}
	if info := detailInfo.GetRequestInfo(); info != nil {
		infoAny, err := anypb.New(info)
		if err != nil {
			panic(err)
		}
		details = append(details, infoAny)
	}
	if info := detailInfo.GetResourceInfo(); info != nil {
		infoAny, err := anypb.New(info)
		if err != nil {
			panic(err)
		}
		details = append(details, infoAny)
	}
	if info := detailInfo.GetHelp(); info != nil {
		infoAny, err := anypb.New(info)
		if err != nil {
			panic(err)
		}
		details = append(details, infoAny)
	}
	if info := detailInfo.GetLocalizedMessage(); info != nil {
		infoAny, err := anypb.New(info)
		if err != nil {
			panic(err)
		}
		details = append(details, infoAny)
	}
	if info := detailInfo.GetExtra(); info != nil {
		infoAny, err := anypb.New(info)
		if err != nil {
			panic(err)
		}
		details = append(details, infoAny)
	}
	return details
}

func (x *DetailInfo) MergeHttpDetails(details *DetailInfo) {
	if info := details.GetIdentifier(); info != nil {
		x.Identifier = info
	}
	if info := details.GetErrorInfo(); info != nil {
		x.ErrorInfo = info
	}
	if info := details.GetRetryInfo(); info != nil {
		x.RetryInfo = info
	}
	if info := details.GetDebugInfo(); info != nil {
		x.DebugInfo = info
	}
	if info := details.GetQuotaFailure(); info != nil {
		x.QuotaFailure = info
	}
	if info := details.GetPreconditionFailure(); info != nil {
		x.PreconditionFailure = info
	}
	if info := details.GetBadRequest(); info != nil {
		x.BadRequest = info
	}
	if info := details.GetRequestInfo(); info != nil {
		x.RequestInfo = info
	}
	if info := details.GetResourceInfo(); info != nil {
		x.ResourceInfo = info
	}
	if info := details.GetHelp(); info != nil {
		x.Help = info
	}
	if info := details.GetLocalizedMessage(); info != nil {
		x.LocalizedMessage = info
	}
	if info := details.GetExtra(); info != nil {
		x.Extra = info
	}
}

func (x *DetailInfo) MergeGrpcDetails(details *DetailInfo) {
	x.MergeHttpDetails(details)
	if info := details.GetHeader(); info != nil {
		x.Header = info
	}
}

func ToGrpcDetails(detailInfo *DetailInfo) []*anypb.Any {
	var details []*anypb.Any
	if info := detailInfo.GetHeader(); info != nil {
		infoAny, err := anypb.New(info)
		if err != nil {
			panic(err)
		}
		details = append(details, infoAny)
	}
	return append(details, ToHttpDetails(detailInfo)...)
}

func FromDetails(details []*anypb.Any) *DetailInfo {
	detailInfo := &DetailInfo{}
	for _, value := range details {
		switch {
		case value.MessageIs(_Identifier):
			detailInfo.Identifier = new(Identifier)
			err := value.UnmarshalTo(detailInfo.Identifier)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(_ErrorInfo):
			detailInfo.ErrorInfo = new(errdetails.ErrorInfo)
			err := value.UnmarshalTo(detailInfo.ErrorInfo)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(_RetryInfo):
			detailInfo.RetryInfo = new(errdetails.RetryInfo)
			err := value.UnmarshalTo(detailInfo.RetryInfo)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(_DebugInfo):
			detailInfo.DebugInfo = new(errdetails.DebugInfo)
			err := value.UnmarshalTo(detailInfo.DebugInfo)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(_QuotaFailure):
			detailInfo.QuotaFailure = new(errdetails.QuotaFailure)
			err := value.UnmarshalTo(detailInfo.QuotaFailure)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(_PreconditionFailure):
			detailInfo.PreconditionFailure = new(errdetails.PreconditionFailure)
			err := value.UnmarshalTo(detailInfo.PreconditionFailure)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(_BadRequest):
			detailInfo.BadRequest = new(errdetails.BadRequest)
			err := value.UnmarshalTo(detailInfo.BadRequest)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(_RequestInfo):
			detailInfo.RequestInfo = new(errdetails.RequestInfo)
			err := value.UnmarshalTo(detailInfo.RequestInfo)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(_ResourceInfo):
			detailInfo.ResourceInfo = new(errdetails.ResourceInfo)
			err := value.UnmarshalTo(detailInfo.ResourceInfo)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(_Help):
			detailInfo.Help = new(errdetails.Help)
			err := value.UnmarshalTo(detailInfo.Help)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(_LocalizedMessage):
			detailInfo.LocalizedMessage = new(errdetails.LocalizedMessage)
			err := value.UnmarshalTo(detailInfo.LocalizedMessage)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(_Header):
			detailInfo.Header = new(Header)
			err := value.UnmarshalTo(detailInfo.Header)
			if err != nil {
				panic(err)
			}
		default:
			detailInfo.Extra = value
		}
	}
	return detailInfo
}

func (x *DetailInfo) Without(detail proto.Message) {
	if x == nil {
		return
	}
	switch detail.(type) {
	case *Identifier:
		x.Identifier = nil
	case *errdetails.ErrorInfo:
		x.ErrorInfo = nil
	case *errdetails.RetryInfo:
		x.RetryInfo = nil
	case *errdetails.DebugInfo:
		x.DebugInfo = nil
	case *errdetails.QuotaFailure:
		x.QuotaFailure = nil
	case *errdetails.PreconditionFailure:
		x.PreconditionFailure = nil
	case *errdetails.BadRequest:
		x.BadRequest = nil
	case *errdetails.RequestInfo:
		x.RequestInfo = nil
	case *errdetails.ResourceInfo:
		x.ResourceInfo = nil
	case *errdetails.Help:
		x.Help = nil
	case *errdetails.LocalizedMessage:
		x.LocalizedMessage = nil
	case *Header:
		x.Header = nil
	default:
		x.Extra = nil
	}
}
