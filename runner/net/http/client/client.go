package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"

	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/go-leo/leo/common/stringx"
	"github.com/go-leo/leo/registry"
	"github.com/go-leo/leo/runner/net/http/header"
	"github.com/go-leo/leo/runner/net/http/internal/codec"
	"github.com/go-leo/leo/runner/net/http/internal/util"
)

type Interface interface {
	Invoke(ctx context.Context, method string, path string, in any, out any) error
}

type CodecType uint8

const (
	JSON CodecType = iota
	Protobuf
)

type Scheme string

const (
	HTTP  Scheme = registry.TransportHTTP
	HTTPS Scheme = registry.TransportHTTPS
)

type options struct {
	Middlewares []Interceptor
	Codec       codec.Codec
	httpClient  *http.Client
	Transport   string
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *options) init() {
	if o.Middlewares == nil {
		o.Middlewares = make([]Interceptor, 0)
	}
	if o.Codec == nil {
		o.Codec = codec.JSONCodec
	}
	if o.httpClient == nil {
		o.httpClient = http.DefaultClient
	}
	if stringx.IsBlank(o.Transport) {
		o.Transport = registry.TransportHTTP
	}
}

type Option func(o *options)

func Middleware(middlewares ...Interceptor) Option {
	return func(o *options) {
		o.Middlewares = append(o.Middlewares, middlewares...)
	}
}

func Codec(c CodecType) Option {
	return func(o *options) {
		switch c {
		case JSON:
			o.Codec = codec.JSONCodec
		case Protobuf:
			o.Codec = codec.ProtobufCodec
		}
	}
}

func NewClient(scheme Scheme, host string, port string, opts ...Option) Interface {
	o := new(options)
	o.apply(opts...)
	o.init()
	tg := fmt.Sprintf("%s://%s", scheme, net.JoinHostPort(host, port))
	return &baseClient{
		o:          o,
		target:     tg,
		middleware: Chain(o.Middlewares...),
	}
}

type baseClient struct {
	o          *options
	target     string
	middleware Interceptor
}

func (baseCli *baseClient) Invoke(ctx context.Context, method string, path string, in any, out any) error {
	// 拼接url
	reqURL := baseCli.getUrl(path)
	// body解码器
	bodyCodec := baseCli.o.Codec
	// 对request编码
	data, err := bodyCodec.Marshal(in)
	if err != nil {
		return err
	}
	// 创建http请求
	req, err := http.NewRequestWithContext(ctx, method, reqURL, bytes.NewReader(data))
	if err != nil {
		return err
	}
	// 设置header
	if header, ok := header.FromContext(ctx); ok {
		for key, vals := range header {
			for _, val := range vals {
				req.Header.Set(key, val)
			}
		}
	}
	// 设置Content-Type和Accept类型
	util.SetContentType(req.Header, bodyCodec.ContentType())
	util.SetAcceptType(req.Header, bodyCodec.ContentType())
	// 如果没有middleware，就直接发送http请求，并对响应进行处理
	if baseCli.middleware == nil {
		return baseCli.do(req, out, nil)
	}
	// 创建HTTPInfo，放入Path、Client、http.Request和http.Response，中间件里可能会用到
	info := new(HTTPInfo)
	info.Path = path
	info.Client = baseCli.o.httpClient
	info.Request = req
	f := func(ctx context.Context, req any, reply any, info *HTTPInfo) (err error) {
		return baseCli.do(info.Request, reply, func(resp *http.Response) { info.Response = resp })
	}
	// 调用中间件与发送http请求
	return baseCli.middleware(ctx, in, out, info, f)
}

func (baseCli *baseClient) Target() string {
	return baseCli.target
}

func (baseCli *baseClient) getUrl(path string) string {
	return baseCli.target + path
}

func (baseCli *baseClient) do(req *http.Request, out any, fn func(response *http.Response)) error {
	resp, err := baseCli.o.httpClient.Do(req)
	if err != nil {
		return err
	}
	if fn != nil {
		fn(resp)
	}
	if resp.StatusCode != http.StatusOK {
		return baseCli.errorHandler(resp)
	}
	return baseCli.decRespBody(resp, out)
}

func (baseCli *baseClient) decRespBody(resp *http.Response, out any) error {
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return codec.GetCodec(util.GetContentType(resp.Header)).Unmarshal(data, out)
}

func (baseCli *baseClient) errorHandler(resp *http.Response) error {
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	pb := spb.Status{}
	err = codec.GetCodec(util.GetContentType(resp.Header)).Unmarshal(data, &pb)
	if err != nil {
		pb.Message = string(data)
		pb.Code = int32(GRPCCodeFromHTTPStatus(resp.StatusCode))
	}
	return status.ErrorProto(&pb)
}

func GRPCCodeFromHTTPStatus(statusCode int) codes.Code {
	switch statusCode {
	case http.StatusOK:
		return codes.OK
	case http.StatusRequestTimeout:
		return codes.Canceled
	case http.StatusBadRequest:
		return codes.Unknown
	case http.StatusGatewayTimeout:
		return codes.DeadlineExceeded
	case http.StatusNotFound:
		return codes.NotFound
	case http.StatusConflict:
		return codes.Unknown
	case http.StatusForbidden:
		return codes.PermissionDenied
	case http.StatusUnauthorized:
		return codes.Unauthenticated
	case http.StatusTooManyRequests:
		return codes.ResourceExhausted
	case http.StatusNotImplemented:
		return codes.Unimplemented
	case http.StatusInternalServerError:
		return codes.Internal
	case http.StatusServiceUnavailable:
		return codes.Unavailable
	}
	return codes.Internal
}
