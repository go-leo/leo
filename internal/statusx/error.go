package statusx

import (
	"errors"
	"fmt"
	"golang.org/x/exp/maps"
	httpstatus "google.golang.org/genproto/googleapis/rpc/http"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
)

const (
	kStatusCoderKey  = "X-Leo-Status-Coder"
	kStatusKeysKey   = "X-Leo-Status-Keys"
	kStatusDetailKey = "X-Leo-Status-Detail"
	kStatusCauseKey  = "X-Leo-Status-Cause"
	kStatusGrpcKey   = "X-Leo-Status-Grpc"
)

func (x *Error) Encode() (int, http.Header, []byte) {
	httpProto := x.GetHttpStatus()
	grpcProto := x.GetGrpcStatus()
	if httpProto == nil || grpcProto == nil {
		panic(errors.New("http status or grpc status is nil"))
	}

	// handle headers
	header := make(http.Header)
	if headers := httpProto.GetHeaders(); len(headers) > 0 {
		for _, h := range headers {
			header.Add(h.GetKey(), h.GetValue())
		}
		keys := maps.Keys(header)
		for _, key := range keys {
			header.Add(kStatusKeysKey, key)
		}
	}

	// handle cause error
	if cause := x.GetCause(); cause != nil {
		causeData, err := protojson.Marshal(cause)
		if err != nil {
			panic(err)
		}
		header.Set(kStatusCauseKey, string(causeData))
	}

	// handle detail info
	if detail := x.GetDetail(); detail != nil {
		detailData, err := protojson.Marshal(detail)
		if err != nil {
			panic(err)
		}
		header.Set(kStatusDetailKey, string(detailData))
	}

	// handle grpc status
	grpcProtoData, err := protojson.Marshal(grpcProto)
	if err != nil {
		panic(err)
	}
	header.Set(kStatusGrpcKey, string(grpcProtoData))

	// return
	return int(httpProto.GetStatus()), header, httpProto.GetBody()
}

func (x *Error) Decode(status int, header http.Header, body []byte) {
	if x == nil {
		panic(fmt.Errorf("statusx: error is nil"))
	}

	// handle http status
	x.HttpStatus = &httpstatus.HttpResponse{
		Status: int32(status),
		Reason: http.StatusText(status),
		Body:   body,
	}

	// handle headers
	keys := header.Values(kStatusKeysKey)
	if len(keys) > 0 {
		headers := make([]*httpstatus.HttpHeader, 0, len(keys))
		for _, key := range keys {
			for _, value := range header.Values(key) {
				headers = append(headers, &httpstatus.HttpHeader{Key: key, Value: value})
			}
		}
		x.HttpStatus.Headers = headers
	}

	// handle cause error
	if causeData := header.Get(kStatusCauseKey); causeData != "" {
		var cause Cause
		if err := protojson.Unmarshal([]byte(causeData), &cause); err != nil {
			panic(err)
		}
		x.Cause = &cause
	}

	// handle detail info
	if detailData := header.Get(kStatusDetailKey); detailData != "" {
		var detail Detail
		if err := protojson.Unmarshal([]byte(detailData), &detail); err != nil {
			panic(err)
		}
		x.Detail = &detail
	}

	// handle grpc status
	grpcProtoData := header.Get(kStatusGrpcKey)
	if grpcProtoData == "" {
		panic(errors.New("grpc status is nil"))
	}
	var grpcProto rpcstatus.Status
	if err := protojson.Unmarshal([]byte(grpcProtoData), &grpcProto); err != nil {
		panic(err)
	}
	x.GrpcStatus = &grpcProto
}
