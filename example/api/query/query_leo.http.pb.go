// Code generated by protoc-gen-leo-http. DO NOT EDIT.

package query

import (
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
	jsonx "github.com/go-leo/gox/encodingx/jsonx"
	errorx "github.com/go-leo/gox/errorx"
	urlx "github.com/go-leo/gox/netx/urlx"
	protox "github.com/go-leo/gox/protox"
	strconvx "github.com/go-leo/gox/strconvx"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	statusx "github.com/go-leo/leo/v3/statusx"
	httpx "github.com/go-leo/leo/v3/transportx/httpx"
	mux "github.com/gorilla/mux"
	proto "google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	io "io"
	http "net/http"
	url "net/url"
)

// =========================== http router ===========================

func appendQueryHttpRoutes(router *mux.Router) *mux.Router {
	router.NewRoute().Name("/leo.example.query.v1.Query/Query").Methods("GET").Path("/v1/query")
	return router
}
func AppendQueryHttpRoutes(router *mux.Router, svc QueryService, middlewares ...endpoint.Middleware) *mux.Router {
	transports := newQueryHttpServerTransports(svc, middlewares...)
	router = appendQueryHttpRoutes(router)
	router.Get("/leo.example.query.v1.Query/Query").Handler(transports.Query())
	return router
}

func NewQueryHttpClient(target string, opts ...httpx.ClientOption) QueryService {
	options := httpx.NewClientOptions(opts...)
	transports := newQueryHttpClientTransports(options.Scheme(), options.ClientTransportOptions(), options.Middlewares())
	endpoints := newQueryClientEndpoints(target, transports, options.InstancerFactory(), options.EndpointerOptions(), options.BalancerFactory(), options.Logger())
	return newQueryClientService(endpoints, httpx.HttpClient)
}

// =========================== http server ===========================

type QueryHttpServerTransports interface {
	Query() http.Handler
}

type QueryHttpServerRequestDecoder interface {
	Query() http1.DecodeRequestFunc
}

type QueryHttpServerResponseEncoder interface {
	Query() http1.EncodeResponseFunc
}

type QueryHttpClientRequestEncoder interface {
	Query() http1.CreateRequestFunc
}

type QueryHttpClientResponseDecoder interface {
	Query() http1.DecodeResponseFunc
}

type queryHttpServerTransports struct {
	endpoints       QueryServerEndpoints
	requestDecoder  QueryHttpServerRequestDecoder
	responseEncoder QueryHttpServerResponseEncoder
}

func (t *queryHttpServerTransports) Query() http.Handler {
	return http1.NewServer(
		t.endpoints.Query(context.TODO()),
		t.requestDecoder.Query(),
		t.responseEncoder.Query(),
		http1.ServerBefore(httpx.EndpointInjector("/leo.example.query.v1.Query/Query")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(httpx.IncomingTimeLimitInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerFinalizer(httpx.CancelInvoker),
		http1.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func newQueryHttpServerTransports(svc QueryService, middlewares ...endpoint.Middleware) QueryHttpServerTransports {
	endpoints := newQueryServerEndpoints(svc, middlewares...)
	return &queryHttpServerTransports{
		endpoints:       endpoints,
		requestDecoder:  queryHttpServerRequestDecoder{},
		responseEncoder: queryHttpServerResponseEncoder{},
	}
}

type queryHttpServerRequestDecoder struct{}

func (queryHttpServerRequestDecoder) Query() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &QueryRequest{}
		queries := r.URL.Query()
		var queryErr error
		req.Bool, queryErr = errorx.Break[bool](queryErr)(urlx.GetBool(queries, "bool"))
		req.Int32, queryErr = errorx.Break[int32](queryErr)(urlx.GetInt[int32](queries, "int32"))
		req.Sint32, queryErr = errorx.Break[int32](queryErr)(urlx.GetInt[int32](queries, "sint32"))
		req.Uint32, queryErr = errorx.Break[uint32](queryErr)(urlx.GetUint[uint32](queries, "uint32"))
		req.Int64, queryErr = errorx.Break[int64](queryErr)(urlx.GetInt[int64](queries, "int64"))
		req.Sint64, queryErr = errorx.Break[int64](queryErr)(urlx.GetInt[int64](queries, "sint64"))
		req.Uint64, queryErr = errorx.Break[uint64](queryErr)(urlx.GetUint[uint64](queries, "uint64"))
		req.Sfixed32, queryErr = errorx.Break[int32](queryErr)(urlx.GetInt[int32](queries, "sfixed32"))
		req.Fixed32, queryErr = errorx.Break[uint32](queryErr)(urlx.GetUint[uint32](queries, "fixed32"))
		req.Float, queryErr = errorx.Break[float32](queryErr)(urlx.GetFloat[float32](queries, "float"))
		req.Sfixed64, queryErr = errorx.Break[int64](queryErr)(urlx.GetInt[int64](queries, "sfixed64"))
		req.Fixed64, queryErr = errorx.Break[uint64](queryErr)(urlx.GetUint[uint64](queries, "fixed64"))
		req.Double, queryErr = errorx.Break[float64](queryErr)(urlx.GetFloat[float64](queries, "double"))
		req.String_ = queries.Get("string")
		req.OptBool, queryErr = errorx.Break[*bool](queryErr)(urlx.GetBoolPtr(queries, "opt_bool"))
		req.OptInt32, queryErr = errorx.Break[*int32](queryErr)(urlx.GetIntPtr[int32](queries, "opt_int32"))
		req.OptSint32, queryErr = errorx.Break[*int32](queryErr)(urlx.GetIntPtr[int32](queries, "opt_sint32"))
		req.OptUint32, queryErr = errorx.Break[*uint32](queryErr)(urlx.GetUintPtr[uint32](queries, "opt_uint32"))
		req.OptInt64, queryErr = errorx.Break[*int64](queryErr)(urlx.GetIntPtr[int64](queries, "opt_int64"))
		req.OptSint64, queryErr = errorx.Break[*int64](queryErr)(urlx.GetIntPtr[int64](queries, "opt_sint64"))
		req.OptUint64, queryErr = errorx.Break[*uint64](queryErr)(urlx.GetUintPtr[uint64](queries, "opt_uint64"))
		req.OptSfixed32, queryErr = errorx.Break[*int32](queryErr)(urlx.GetIntPtr[int32](queries, "opt_sfixed32"))
		req.OptFixed32, queryErr = errorx.Break[*uint32](queryErr)(urlx.GetUintPtr[uint32](queries, "opt_fixed32"))
		req.OptFloat, queryErr = errorx.Break[*float32](queryErr)(urlx.GetFloatPtr[float32](queries, "opt_float"))
		req.OptSfixed64, queryErr = errorx.Break[*int64](queryErr)(urlx.GetIntPtr[int64](queries, "opt_sfixed64"))
		req.OptFixed64, queryErr = errorx.Break[*uint64](queryErr)(urlx.GetUintPtr[uint64](queries, "opt_fixed64"))
		req.OptDouble, queryErr = errorx.Break[*float64](queryErr)(urlx.GetFloatPtr[float64](queries, "opt_double"))
		req.OptString = proto.String(queries.Get("opt_string"))
		req.RepBool, queryErr = errorx.Break[[]bool](queryErr)(urlx.GetBoolSlice(queries, "rep_bool"))
		req.RepInt32, queryErr = errorx.Break[[]int32](queryErr)(urlx.GetIntSlice[int32](queries, "rep_int32"))
		req.RepSint32, queryErr = errorx.Break[[]int32](queryErr)(urlx.GetIntSlice[int32](queries, "rep_sint32"))
		req.RepUint32, queryErr = errorx.Break[[]uint32](queryErr)(urlx.GetUintSlice[uint32](queries, "rep_uint32"))
		req.RepInt64, queryErr = errorx.Break[[]int64](queryErr)(urlx.GetIntSlice[int64](queries, "rep_int64"))
		req.RepSint64, queryErr = errorx.Break[[]int64](queryErr)(urlx.GetIntSlice[int64](queries, "rep_sint64"))
		req.RepUint64, queryErr = errorx.Break[[]uint64](queryErr)(urlx.GetUintSlice[uint64](queries, "rep_uint64"))
		req.RepSfixed32, queryErr = errorx.Break[[]int32](queryErr)(urlx.GetIntSlice[int32](queries, "rep_sfixed32"))
		req.RepFixed32, queryErr = errorx.Break[[]uint32](queryErr)(urlx.GetUintSlice[uint32](queries, "rep_fixed32"))
		req.RepFloat, queryErr = errorx.Break[[]float32](queryErr)(urlx.GetFloatSlice[float32](queries, "rep_float"))
		req.RepSfixed64, queryErr = errorx.Break[[]int64](queryErr)(urlx.GetIntSlice[int64](queries, "rep_sfixed64"))
		req.RepFixed64, queryErr = errorx.Break[[]uint64](queryErr)(urlx.GetUintSlice[uint64](queries, "rep_fixed64"))
		req.RepDouble, queryErr = errorx.Break[[]float64](queryErr)(urlx.GetFloatSlice[float64](queries, "rep_double"))
		req.RepString = queries["rep_string"]
		req.WrapDouble, queryErr = errorx.Break[*wrapperspb.DoubleValue](queryErr)(urlx.GetFloat64Value(queries, "wrap_double"))
		req.WrapFloat, queryErr = errorx.Break[*wrapperspb.FloatValue](queryErr)(urlx.GetFloat32Value(queries, "wrap_float"))
		req.WrapInt64, queryErr = errorx.Break[*wrapperspb.Int64Value](queryErr)(urlx.GetInt64Value(queries, "wrap_int64"))
		req.WrapUint64, queryErr = errorx.Break[*wrapperspb.UInt64Value](queryErr)(urlx.GetUint64Value(queries, "wrap_uint64"))
		req.WrapInt32, queryErr = errorx.Break[*wrapperspb.Int32Value](queryErr)(urlx.GetInt32Value(queries, "wrap_int32"))
		req.WrapUint32, queryErr = errorx.Break[*wrapperspb.UInt32Value](queryErr)(urlx.GetUint32Value(queries, "wrap_uint32"))
		req.WrapBool, queryErr = errorx.Break[*wrapperspb.BoolValue](queryErr)(urlx.GetBoolValue(queries, "wrap_bool"))
		req.WrapString = wrapperspb.String(queries.Get("wrap_string"))
		req.OptWrapDouble, queryErr = errorx.Break[*wrapperspb.DoubleValue](queryErr)(urlx.GetFloat64Value(queries, "opt_wrap_double"))
		req.OptWrapFloat, queryErr = errorx.Break[*wrapperspb.FloatValue](queryErr)(urlx.GetFloat32Value(queries, "opt_wrap_float"))
		req.OptWrapInt64, queryErr = errorx.Break[*wrapperspb.Int64Value](queryErr)(urlx.GetInt64Value(queries, "opt_wrap_int64"))
		req.OptWrapUint64, queryErr = errorx.Break[*wrapperspb.UInt64Value](queryErr)(urlx.GetUint64Value(queries, "opt_wrap_uint64"))
		req.OptWrapInt32, queryErr = errorx.Break[*wrapperspb.Int32Value](queryErr)(urlx.GetInt32Value(queries, "opt_wrap_int32"))
		req.OptWrapUint32, queryErr = errorx.Break[*wrapperspb.UInt32Value](queryErr)(urlx.GetUint32Value(queries, "opt_wrap_uint32"))
		req.OptWrapBool, queryErr = errorx.Break[*wrapperspb.BoolValue](queryErr)(urlx.GetBoolValue(queries, "opt_wrap_bool"))
		req.OptWrapString = wrapperspb.String(queries.Get("opt_wrap_string"))
		req.RepWrapDouble, queryErr = errorx.Break[[]*wrapperspb.DoubleValue](queryErr)(urlx.GetFloat64ValueSlice(queries, "rep_wrap_double"))
		req.RepWrapFloat, queryErr = errorx.Break[[]*wrapperspb.FloatValue](queryErr)(urlx.GetFloat32ValueSlice(queries, "rep_wrap_float"))
		req.RepWrapInt64, queryErr = errorx.Break[[]*wrapperspb.Int64Value](queryErr)(urlx.GetInt64ValueSlice(queries, "rep_wrap_int64"))
		req.RepWrapUint64, queryErr = errorx.Break[[]*wrapperspb.UInt64Value](queryErr)(urlx.GetUint64ValueSlice(queries, "rep_wrap_uint64"))
		req.RepWrapInt32, queryErr = errorx.Break[[]*wrapperspb.Int32Value](queryErr)(urlx.GetInt32ValueSlice(queries, "rep_wrap_int32"))
		req.RepWrapUint32, queryErr = errorx.Break[[]*wrapperspb.UInt32Value](queryErr)(urlx.GetUint32ValueSlice(queries, "rep_wrap_uint32"))
		req.RepWrapBool, queryErr = errorx.Break[[]*wrapperspb.BoolValue](queryErr)(urlx.GetBoolValueSlice(queries, "rep_wrap_bool"))
		req.RepWrapString = protox.WrapStringSlice(queries["rep_wrap_string"])
		req.Status, queryErr = errorx.Break[QueryRequest_Status](queryErr)(urlx.GetInt[QueryRequest_Status](queries, "status"))
		req.OptStatus, queryErr = errorx.Break[*QueryRequest_Status](queryErr)(urlx.GetIntPtr[QueryRequest_Status](queries, "opt_status"))
		req.RepStatus, queryErr = errorx.Break[[]QueryRequest_Status](queryErr)(urlx.GetIntSlice[QueryRequest_Status](queries, "rep_status"))
		if queryErr != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(queryErr))
		}
		return req, nil
	}
}

type queryHttpServerResponseEncoder struct{}

func (queryHttpServerResponseEncoder) Query() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*emptypb.Empty)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
			return statusx.ErrInternal.With(statusx.Wrap(err))
		}
		return nil
	}
}

type queryHttpClientTransports struct {
	scheme          string
	router          *mux.Router
	clientOptions   []http1.ClientOption
	middlewares     []endpoint.Middleware
	requestEncoder  QueryHttpClientRequestEncoder
	responseDecoder QueryHttpClientResponseDecoder
}

func (t *queryHttpClientTransports) Query(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(httpx.OutgoingTimeLimitInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		t.requestEncoder.Query(),
		t.responseDecoder.Query(),
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func newQueryHttpClientTransports(scheme string, clientOptions []http1.ClientOption, middlewares []endpoint.Middleware) QueryClientTransports {
	return &queryHttpClientTransports{
		scheme:          scheme,
		router:          appendQueryHttpRoutes(mux.NewRouter()),
		clientOptions:   clientOptions,
		middlewares:     middlewares,
		requestEncoder:  nil,
		responseDecoder: queryHttpClientResponseDecoder{},
	}
}

type queryHttpClientResponseDecoder struct{}

func (queryHttpClientResponseDecoder) Query() http1.DecodeResponseFunc {
	return func(ctx context.Context, r *http.Response) (any, error) {
		if httpx.IsErrorResponse(r) {
			return nil, httpx.ErrorDecoder(ctx, r)
		}
		resp := &emptypb.Empty{}
		if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
			return nil, err
		}
		return resp, nil
	}
}

// =========================== http coder ===========================

func _Query_Query_HttpClient_RequestEncoder(router *mux.Router) func(scheme string, instance string) http1.CreateRequestFunc {
	return func(scheme string, instance string) http1.CreateRequestFunc {
		return func(ctx context.Context, obj any) (*http.Request, error) {
			if obj == nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
			}
			req, ok := obj.(*QueryRequest)
			if !ok {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid request type, %T", obj))
			}
			_ = req
			var body io.Reader
			var pairs []string
			path, err := router.Get("/leo.example.query.v1.Query/Query").URLPath(pairs...)
			if err != nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
			}
			queries := url.Values{}
			queries["bool"] = append(queries["bool"], strconvx.FormatBool(req.GetBool()))
			queries["int32"] = append(queries["int32"], strconvx.FormatInt(req.GetInt32(), 10))
			queries["sint32"] = append(queries["sint32"], strconvx.FormatInt(req.GetSint32(), 10))
			queries["uint32"] = append(queries["uint32"], strconvx.FormatUint(req.GetUint32(), 10))
			queries["int64"] = append(queries["int64"], strconvx.FormatInt(req.GetInt64(), 10))
			queries["sint64"] = append(queries["sint64"], strconvx.FormatInt(req.GetSint64(), 10))
			queries["uint64"] = append(queries["uint64"], strconvx.FormatUint(req.GetUint64(), 10))
			queries["sfixed32"] = append(queries["sfixed32"], strconvx.FormatInt(req.GetSfixed32(), 10))
			queries["fixed32"] = append(queries["fixed32"], strconvx.FormatUint(req.GetFixed32(), 10))
			queries["float"] = append(queries["float"], strconvx.FormatFloat(req.GetFloat(), 'f', -1, 32))
			queries["sfixed64"] = append(queries["sfixed64"], strconvx.FormatInt(req.GetSfixed64(), 10))
			queries["fixed64"] = append(queries["fixed64"], strconvx.FormatUint(req.GetFixed64(), 10))
			queries["double"] = append(queries["double"], strconvx.FormatFloat(req.GetDouble(), 'f', -1, 64))
			queries["string"] = append(queries["string"], req.GetString_())
			queries["opt_bool"] = append(queries["opt_bool"], strconvx.FormatBool(req.GetOptBool()))
			queries["opt_int32"] = append(queries["opt_int32"], strconvx.FormatInt(req.GetOptInt32(), 10))
			queries["opt_sint32"] = append(queries["opt_sint32"], strconvx.FormatInt(req.GetOptSint32(), 10))
			queries["opt_uint32"] = append(queries["opt_uint32"], strconvx.FormatUint(req.GetOptUint32(), 10))
			queries["opt_int64"] = append(queries["opt_int64"], strconvx.FormatInt(req.GetOptInt64(), 10))
			queries["opt_sint64"] = append(queries["opt_sint64"], strconvx.FormatInt(req.GetOptSint64(), 10))
			queries["opt_uint64"] = append(queries["opt_uint64"], strconvx.FormatUint(req.GetOptUint64(), 10))
			queries["opt_sfixed32"] = append(queries["opt_sfixed32"], strconvx.FormatInt(req.GetOptSfixed32(), 10))
			queries["opt_fixed32"] = append(queries["opt_fixed32"], strconvx.FormatUint(req.GetOptFixed32(), 10))
			queries["opt_float"] = append(queries["opt_float"], strconvx.FormatFloat(req.GetOptFloat(), 'f', -1, 32))
			queries["opt_sfixed64"] = append(queries["opt_sfixed64"], strconvx.FormatInt(req.GetOptSfixed64(), 10))
			queries["opt_fixed64"] = append(queries["opt_fixed64"], strconvx.FormatUint(req.GetOptFixed64(), 10))
			queries["opt_double"] = append(queries["opt_double"], strconvx.FormatFloat(req.GetOptDouble(), 'f', -1, 64))
			queries["opt_string"] = append(queries["opt_string"], req.GetOptString())
			queries["rep_bool"] = append(queries["rep_bool"], strconvx.FormatBoolSlice(req.GetRepBool())...)
			queries["rep_int32"] = append(queries["rep_int32"], strconvx.FormatIntSlice(req.GetRepInt32(), 10)...)
			queries["rep_sint32"] = append(queries["rep_sint32"], strconvx.FormatIntSlice(req.GetRepSint32(), 10)...)
			queries["rep_uint32"] = append(queries["rep_uint32"], strconvx.FormatUintSlice(req.GetRepUint32(), 10)...)
			queries["rep_int64"] = append(queries["rep_int64"], strconvx.FormatIntSlice(req.GetRepInt64(), 10)...)
			queries["rep_sint64"] = append(queries["rep_sint64"], strconvx.FormatIntSlice(req.GetRepSint64(), 10)...)
			queries["rep_uint64"] = append(queries["rep_uint64"], strconvx.FormatUintSlice(req.GetRepUint64(), 10)...)
			queries["rep_sfixed32"] = append(queries["rep_sfixed32"], strconvx.FormatIntSlice(req.GetRepSfixed32(), 10)...)
			queries["rep_fixed32"] = append(queries["rep_fixed32"], strconvx.FormatUintSlice(req.GetRepFixed32(), 10)...)
			queries["rep_float"] = append(queries["rep_float"], strconvx.FormatFloatSlice(req.GetRepFloat(), 'f', -1, 32)...)
			queries["rep_sfixed64"] = append(queries["rep_sfixed64"], strconvx.FormatIntSlice(req.GetRepSfixed64(), 10)...)
			queries["rep_fixed64"] = append(queries["rep_fixed64"], strconvx.FormatUintSlice(req.GetRepFixed64(), 10)...)
			queries["rep_double"] = append(queries["rep_double"], strconvx.FormatFloatSlice(req.GetRepDouble(), 'f', -1, 64)...)
			queries["rep_string"] = append(queries["rep_string"], req.GetRepString()...)
			queries["wrap_double"] = append(queries["wrap_double"], strconvx.FormatFloat(req.GetWrapDouble().GetValue(), 'f', -1, 64))
			queries["wrap_float"] = append(queries["wrap_float"], strconvx.FormatFloat(req.GetWrapFloat().GetValue(), 'f', -1, 32))
			queries["wrap_int64"] = append(queries["wrap_int64"], strconvx.FormatInt(req.GetWrapInt64().GetValue(), 10))
			queries["wrap_uint64"] = append(queries["wrap_uint64"], strconvx.FormatUint(req.GetWrapUint64().GetValue(), 10))
			queries["wrap_int32"] = append(queries["wrap_int32"], strconvx.FormatInt(req.GetWrapInt32().GetValue(), 10))
			queries["wrap_uint32"] = append(queries["wrap_uint32"], strconvx.FormatUint(req.GetWrapUint32().GetValue(), 10))
			queries["wrap_bool"] = append(queries["wrap_bool"], strconvx.FormatBool(req.GetWrapBool().GetValue()))
			queries["wrap_string"] = append(queries["wrap_string"], req.GetWrapString().GetValue())
			queries["opt_wrap_double"] = append(queries["opt_wrap_double"], strconvx.FormatFloat(req.GetOptWrapDouble().GetValue(), 'f', -1, 64))
			queries["opt_wrap_float"] = append(queries["opt_wrap_float"], strconvx.FormatFloat(req.GetOptWrapFloat().GetValue(), 'f', -1, 32))
			queries["opt_wrap_int64"] = append(queries["opt_wrap_int64"], strconvx.FormatInt(req.GetOptWrapInt64().GetValue(), 10))
			queries["opt_wrap_uint64"] = append(queries["opt_wrap_uint64"], strconvx.FormatUint(req.GetOptWrapUint64().GetValue(), 10))
			queries["opt_wrap_int32"] = append(queries["opt_wrap_int32"], strconvx.FormatInt(req.GetOptWrapInt32().GetValue(), 10))
			queries["opt_wrap_uint32"] = append(queries["opt_wrap_uint32"], strconvx.FormatUint(req.GetOptWrapUint32().GetValue(), 10))
			queries["opt_wrap_bool"] = append(queries["opt_wrap_bool"], strconvx.FormatBool(req.GetOptWrapBool().GetValue()))
			queries["opt_wrap_string"] = append(queries["opt_wrap_string"], req.GetOptWrapString().GetValue())
			queries["rep_wrap_double"] = append(queries["rep_wrap_double"], strconvx.FormatFloatSlice(protox.UnwrapFloat64Slice(req.GetRepWrapDouble()), 'f', -1, 64)...)
			queries["rep_wrap_float"] = append(queries["rep_wrap_float"], strconvx.FormatFloatSlice(protox.UnwrapFloat32Slice(req.GetRepWrapFloat()), 'f', -1, 32)...)
			queries["rep_wrap_int64"] = append(queries["rep_wrap_int64"], strconvx.FormatIntSlice(protox.UnwrapInt64Slice(req.GetRepWrapInt64()), 10)...)
			queries["rep_wrap_uint64"] = append(queries["rep_wrap_uint64"], strconvx.FormatUintSlice(protox.UnwrapUint64Slice(req.GetRepWrapUint64()), 10)...)
			queries["rep_wrap_int32"] = append(queries["rep_wrap_int32"], strconvx.FormatIntSlice(protox.UnwrapInt32Slice(req.GetRepWrapInt32()), 10)...)
			queries["rep_wrap_uint32"] = append(queries["rep_wrap_uint32"], strconvx.FormatUintSlice(protox.UnwrapUint32Slice(req.GetRepWrapUint32()), 10)...)
			queries["rep_wrap_bool"] = append(queries["rep_wrap_bool"], strconvx.FormatBoolSlice(protox.UnwrapBoolSlice(req.GetRepWrapBool()))...)
			queries["rep_wrap_string"] = append(queries["rep_wrap_string"], protox.UnwrapStringSlice(req.GetRepWrapString())...)
			queries["status"] = append(queries["status"], strconvx.FormatInt(req.GetStatus(), 10))
			queries["opt_status"] = append(queries["opt_status"], strconvx.FormatInt(req.GetOptStatus(), 10))
			queries["rep_status"] = append(queries["rep_status"], strconvx.FormatIntSlice(req.GetRepStatus(), 10)...)
			target := &url.URL{
				Scheme:   scheme,
				Host:     instance,
				Path:     path.Path,
				RawQuery: queries.Encode(),
			}
			r, err := http.NewRequestWithContext(ctx, "GET", target.String(), body)
			if err != nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
			}
			return r, nil
		}
	}
}
