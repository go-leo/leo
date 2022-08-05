package client

import (
	"bytes"
	"context"
	"io"
	"net/http"

	statuspb "google.golang.org/genproto/googleapis/rpc/status"
	grpcstatus "google.golang.org/grpc/status"

	"github.com/go-leo/leo/runner/net/http/header"
	"github.com/go-leo/leo/runner/net/http/internal/codec"
	"github.com/go-leo/leo/runner/net/http/internal/status"
	"github.com/go-leo/leo/runner/net/http/internal/util"
)

var _ Interface = new(Client)

type Client struct {
	o          *options
	baseURL    string
	middleware Interceptor
}

func NewClient(baseURL string, opts ...Option) *Client {
	o := new(options)
	o.apply(opts...)
	o.init()
	return &Client{
		o:          o,
		baseURL:    baseURL,
		middleware: Chain(o.Middlewares...),
	}
}

func newClient(baseURL string, o *options) *Client { // nolint
	return &Client{
		o:          o,
		baseURL:    baseURL,
		middleware: Chain(o.Middlewares...),
	}
}

func (cli *Client) Invoke(ctx context.Context, method string, path string, in any, out any) error {
	// 拼接url
	reqURL := cli.getUrl(path)
	// body解码器
	bodyCodec := cli.o.Codec
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
	if hdr, ok := header.FromContext(ctx); ok {
		for key, vals := range hdr {
			for _, val := range vals {
				req.Header.Set(key, val)
			}
		}
	}
	// 设置Content-Type和Accept类型
	util.SetContentType(req.Header, bodyCodec.ContentType())
	util.SetAcceptType(req.Header, bodyCodec.ContentType())
	// 如果没有middleware，就直接发送http请求，并对响应进行处理
	if cli.middleware == nil {
		return cli.do(req, out, nil)
	}
	// 创建HTTPInfo，放入Path、Client、http.Request和http.Response，中间件里可能会用到
	info := new(HTTPInfo)
	info.Path = path
	info.Client = cli.o.HttpClient
	info.Request = req
	f := func(ctx context.Context, req any, reply any, info *HTTPInfo) (err error) {
		return cli.do(info.Request, reply, func(resp *http.Response) { info.Response = resp })
	}
	// 调用中间件与发送http请求
	return cli.middleware(ctx, in, out, info, f)
}

func (cli *Client) getUrl(path string) string {
	return cli.baseURL + path
}

func (cli *Client) do(req *http.Request, out any, fn func(response *http.Response)) error {
	resp, err := cli.o.HttpClient.Do(req)
	if err != nil {
		return err
	}
	if fn != nil {
		fn(resp)
	}
	if resp.StatusCode != http.StatusOK {
		return cli.errorHandler(resp)
	}
	return cli.decRespBody(resp, out)
}

func (cli *Client) decRespBody(resp *http.Response, out any) error {
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyCodec := codec.GetCodec(util.GetContentType(resp.Header))
	return bodyCodec.Unmarshal(data, out)
}

func (cli *Client) errorHandler(resp *http.Response) error {
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	pb := statuspb.Status{}
	bodyCodec := codec.GetCodec(util.GetContentType(resp.Header))
	err = bodyCodec.Unmarshal(data, &pb)
	if err != nil {
		pb.Message = string(data)
		pb.Code = int32(status.GRPCCodeFromHTTPStatus(resp.StatusCode))
	}
	return grpcstatus.ErrorProto(&pb)
}
