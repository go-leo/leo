package http

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"sync/atomic"

	"codeup.aliyun.com/qimao/leo/leo/stream"
)

var _ stream.Publisher = new(Publisher)

type Publisher struct {
	o      *options
	wg     sync.WaitGroup
	closed atomic.Bool
	method string
	uri    *url.URL
}

func (pub *Publisher) Topic() string {
	return topic(pub.method, pub.uri.RequestURI())
}

func (pub *Publisher) Queue() string {
	return "http"
}

func (pub *Publisher) Publish(ctx context.Context, messages ...*stream.Message) (stream.Result, error) {
	if len(messages) == 0 {
		return nil, nil
	}
	if pub.closed.Load() {
		return nil, stream.ErrPublisherClosed
	}

	pub.wg.Add(1)
	defer pub.wg.Done()

	var result stream.Results
	for _, msg := range messages {
		req, err := pub.o.Marshaller.Marshal(ctx, pub.method, pub.uri, msg)
		if err != nil {
			return nil, err
		}
		if pub.o.OnMessageSending != nil {
			req = pub.o.OnMessageSending(msg, req)
		}
		resp, err := pub.o.HttpClient.Do(req)
		if err != nil {
			return nil, err
		}
		produceResult := &PublishResult{
			Req:  req,
			Resp: resp,
			Msg:  msg,
		}
		result = append(result, produceResult)
	}
	return result, nil
}

func (pub *Publisher) Close(_ context.Context) error {
	if !pub.closed.CompareAndSwap(false, true) {
		return nil
	}
	pub.wg.Wait()
	return nil
}

type PublishResult struct {
	Msg  *stream.Message
	Req  *http.Request
	Resp *http.Response
}

func (p PublishResult) String() string {
	if p.Resp != nil {
		content, _ := io.ReadAll(p.Resp.Body)
		return fmt.Sprintf("code: %d, content: %s", p.Resp.StatusCode, content)
	}
	return "response is nil"
}

func NewPublisher(method, rawURL string, opts ...Option) (*Publisher, error) {
	o := &options{}
	o.apply(opts...)
	o.init()
	uri, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	return &Publisher{
		o:      o,
		method: method,
		uri:    uri,
	}, nil
}
