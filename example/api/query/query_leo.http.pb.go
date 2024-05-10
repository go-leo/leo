// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package query

import (
	context "context"
	errors "errors"
	fmt "fmt"
	endpoint "github.com/go-kit/kit/endpoint"
	http "github.com/go-kit/kit/transport/http"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	mux "github.com/gorilla/mux"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
	http1 "net/http"
	url "net/url"
	strconv "strconv"
)

type queryHTTPClient struct {
	query endpoint.Endpoint
}

func (c *queryHTTPClient) Query(ctx context.Context, request *QueryRequest) (*emptypb.Empty, error) {
	rep, err := c.query(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func NewQueryHTTPClient(
	scheme string,
	instance string,
	mdw []endpoint.Middleware,
	opts ...http.ClientOption,
) interface {
	Query(ctx context.Context, request *QueryRequest) (*emptypb.Empty, error)
} {
	router := mux.NewRouter()
	router.NewRoute().
		Name("/leo.example.query.v1.Query/Query").
		Methods("GET").
		Path("/v1/query")
	return &queryHTTPClient{
		query: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					if obj == nil {
						return nil, errors.New("request object is nil")
					}
					req, ok := obj.(*QueryRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					_ = req
					var body io.Reader
					var pairs []string
					path, err := router.Get("/leo.example.query.v1.Query/Query").URLPath(pairs...)
					if err != nil {
						return nil, err
					}
					queries := url.Values{}
					queries.Add("bool", strconv.FormatBool(req.Bool))
					queries.Add("int32", strconv.FormatInt(int64(req.Int32), 10))
					queries.Add("sint32", strconv.FormatInt(int64(req.Sint32), 10))
					queries.Add("uint32", strconv.FormatUint(uint64(req.Uint32), 10))
					queries.Add("int64", strconv.FormatInt(req.Int64, 10))
					queries.Add("sint64", strconv.FormatInt(req.Sint64, 10))
					queries.Add("uint64", strconv.FormatUint(req.Uint64, 10))
					queries.Add("sfixed32", strconv.FormatInt(int64(req.Sfixed32), 10))
					queries.Add("fixed32", strconv.FormatUint(uint64(req.Fixed32), 10))
					queries.Add("float", strconv.FormatFloat(float64(req.Float), 'f', -1, 32))
					queries.Add("sfixed64", strconv.FormatInt(req.Sfixed64, 10))
					queries.Add("fixed64", strconv.FormatUint(req.Fixed64, 10))
					queries.Add("double", strconv.FormatFloat(req.Double, 'f', -1, 64))
					queries.Add("string", req.String_)
					if req.OptBool == nil {
						queries.Add("opt_bool", strconv.FormatBool(*req.OptBool))
					}
					if req.OptInt32 == nil {
						queries.Add("opt_int32", strconv.FormatInt(int64(*req.OptInt32), 10))
					}
					if req.OptSint32 == nil {
						queries.Add("opt_sint32", strconv.FormatInt(int64(*req.OptSint32), 10))
					}
					if req.OptUint32 == nil {
						queries.Add("opt_uint32", strconv.FormatUint(uint64(*req.OptUint32), 10))
					}
					if req.OptInt64 == nil {
						queries.Add("opt_int64", strconv.FormatInt(*req.OptInt64, 10))
					}
					if req.OptSint64 == nil {
						queries.Add("opt_sint64", strconv.FormatInt(*req.OptSint64, 10))
					}
					if req.OptUint64 == nil {
						queries.Add("opt_uint64", strconv.FormatUint(*req.OptUint64, 10))
					}
					if req.OptSfixed32 == nil {
						queries.Add("opt_sfixed32", strconv.FormatInt(int64(*req.OptSfixed32), 10))
					}
					if req.OptFixed32 == nil {
						queries.Add("opt_fixed32", strconv.FormatUint(uint64(*req.OptFixed32), 10))
					}
					if req.OptFloat == nil {
						queries.Add("opt_float", strconv.FormatFloat(float64(*req.OptFloat), 'f', -1, 32))
					}
					if req.OptSfixed64 == nil {
						queries.Add("opt_sfixed64", strconv.FormatInt(*req.OptSfixed64, 10))
					}
					if req.OptFixed64 == nil {
						queries.Add("opt_fixed64", strconv.FormatUint(*req.OptFixed64, 10))
					}
					if req.OptDouble == nil {
						queries.Add("opt_double", strconv.FormatFloat(*req.OptDouble, 'f', -1, 64))
					}
					if req.OptString == nil {
						queries.Add("opt_string", *req.OptString)
					}
					for _, item := range req.RepBool {
						queries.Add("rep_bool", strconv.FormatBool(item))
					}
					for _, item := range req.RepInt32 {
						queries.Add("rep_int32", strconv.FormatInt(int64(item), 10))
					}
					for _, item := range req.RepSint32 {
						queries.Add("rep_sint32", strconv.FormatInt(int64(item), 10))
					}
					for _, item := range req.RepUint32 {
						queries.Add("rep_uint32", strconv.FormatUint(uint64(item), 10))
					}
					for _, item := range req.RepInt64 {
						queries.Add("rep_int64", strconv.FormatInt(item, 10))
					}
					for _, item := range req.RepSint64 {
						queries.Add("rep_sint64", strconv.FormatInt(item, 10))
					}
					for _, item := range req.RepUint64 {
						queries.Add("rep_uint64", strconv.FormatUint(item, 10))
					}
					for _, item := range req.RepSfixed32 {
						queries.Add("rep_sfixed32", strconv.FormatInt(int64(item), 10))
					}
					for _, item := range req.RepFixed32 {
						queries.Add("rep_fixed32", strconv.FormatUint(uint64(item), 10))
					}
					for _, item := range req.RepFloat {
						queries.Add("rep_float", strconv.FormatFloat(float64(item), 'f', -1, 32))
					}
					for _, item := range req.RepSfixed64 {
						queries.Add("rep_sfixed64", strconv.FormatInt(item, 10))
					}
					for _, item := range req.RepFixed64 {
						queries.Add("rep_fixed64", strconv.FormatUint(item, 10))
					}
					for _, item := range req.RepDouble {
						queries.Add("rep_double", strconv.FormatFloat(item, 'f', -1, 64))
					}
					for _, item := range req.RepString {
						queries.Add("rep_string", item)
					}
					if req.WrapDouble == nil {
						queries.Add("wrap_double", strconv.FormatFloat(req.WrapDouble.Value, 'f', -1, 64))
					}
					if req.WrapFloat == nil {
						queries.Add("wrap_float", strconv.FormatFloat(float64(req.WrapFloat.Value), 'f', -1, 32))
					}
					if req.WrapInt64 == nil {
						queries.Add("wrap_int64", strconv.FormatInt(req.WrapInt64.Value, 10))
					}
					if req.WrapUint64 == nil {
						queries.Add("wrap_uint64", strconv.FormatUint(req.WrapUint64.Value, 10))
					}
					if req.WrapInt32 == nil {
						queries.Add("wrap_int32", strconv.FormatInt(int64(req.WrapInt32.Value), 10))
					}
					if req.WrapUint32 == nil {
						queries.Add("wrap_uint32", strconv.FormatUint(uint64(req.WrapUint32.Value), 10))
					}
					if req.WrapBool == nil {
						queries.Add("wrap_bool", strconv.FormatBool(req.WrapBool.Value))
					}
					if req.WrapString == nil {
						queries.Add("wrap_string", req.WrapString.Value)
					}
					if req.OptWrapDouble == nil {
						queries.Add("opt_wrap_double", strconv.FormatFloat(req.OptWrapDouble.Value, 'f', -1, 64))
					}
					if req.OptWrapFloat == nil {
						queries.Add("opt_wrap_float", strconv.FormatFloat(float64(req.OptWrapFloat.Value), 'f', -1, 32))
					}
					if req.OptWrapInt64 == nil {
						queries.Add("opt_wrap_int64", strconv.FormatInt(req.OptWrapInt64.Value, 10))
					}
					if req.OptWrapUint64 == nil {
						queries.Add("opt_wrap_uint64", strconv.FormatUint(req.OptWrapUint64.Value, 10))
					}
					if req.OptWrapInt32 == nil {
						queries.Add("opt_wrap_int32", strconv.FormatInt(int64(req.OptWrapInt32.Value), 10))
					}
					if req.OptWrapUint32 == nil {
						queries.Add("opt_wrap_uint32", strconv.FormatUint(uint64(req.OptWrapUint32.Value), 10))
					}
					if req.OptWrapBool == nil {
						queries.Add("opt_wrap_bool", strconv.FormatBool(req.OptWrapBool.Value))
					}
					if req.OptWrapString == nil {
						queries.Add("opt_wrap_string", req.OptWrapString.Value)
					}
					for _, item := range req.RepWrapDouble {
						queries.Add("rep_wrap_double", strconv.FormatFloat(item.Value, 'f', -1, 64))
					}
					for _, item := range req.RepWrapFloat {
						queries.Add("rep_wrap_float", strconv.FormatFloat(float64(item.Value), 'f', -1, 32))
					}
					for _, item := range req.RepWrapInt64 {
						queries.Add("rep_wrap_int64", strconv.FormatInt(item.Value, 10))
					}
					for _, item := range req.RepWrapUint64 {
						queries.Add("rep_wrap_uint64", strconv.FormatUint(item.Value, 10))
					}
					for _, item := range req.RepWrapInt32 {
						queries.Add("rep_wrap_int32", strconv.FormatInt(int64(item.Value), 10))
					}
					for _, item := range req.RepWrapUint32 {
						queries.Add("rep_wrap_uint32", strconv.FormatUint(uint64(item.Value), 10))
					}
					for _, item := range req.RepWrapBool {
						queries.Add("rep_wrap_bool", strconv.FormatBool(item.Value))
					}
					for _, item := range req.RepWrapString {
						queries.Add("rep_wrap_string", item.Value)
					}
					queries.Add("status", strconv.FormatInt(int64(req.Status), 10))
					if req.OptStatus == nil {
						queries.Add("opt_status", strconv.FormatInt(int64(*req.OptStatus), 10))
					}
					for _, item := range req.RepStatus {
						queries.Add("rep_status", strconv.FormatInt(int64(item), 10))
					}
					target := &url.URL{
						Scheme:   scheme,
						Host:     instance,
						Path:     path.Path,
						RawQuery: queries.Encode(),
					}
					r, err := http1.NewRequestWithContext(ctx, "GET", target.String(), body)
					if err != nil {
						return nil, err
					}
					return r, nil
				},
				func(ctx context.Context, r *http1.Response) (interface{}, error) {
					return nil, nil
				},
				opts...,
			).Endpoint(),
			mdw...),
	}
}
