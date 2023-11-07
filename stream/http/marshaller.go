package http

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"codeup.aliyun.com/qimao/leo/leo/stream"
)

// Marshaller marshals stream's message to *kafka.Message and unmarshals *kafka.Message to stream's message.
type Marshaller interface {
	Marshal(ctx context.Context, topic string, method string, url string, msg *stream.Message) (*http.Request, error)
	Unmarshal(ctx context.Context, topic string, resp *http.Response) (*stream.Message, error)
}

var _ Marshaller = (*DefaultMarshaller)(nil)

type DefaultMarshaller struct{}

func (d DefaultMarshaller) Marshal(ctx context.Context, topic string, method string, urlStr string, msg *stream.Message) (*http.Request, error) {
	if len(method) <= 0 {
		return nil, errors.New("method is empty")
	}
	if len(urlStr) <= 0 {
		return nil, errors.New("url is empty")
	}

	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	msg.Topic = topic
	if msg.Time.IsZero() {
		msg.Time = time.Now()
	}

	var body io.Reader
	if !strings.EqualFold(method, http.MethodGet) {
		body = bytes.NewReader(msg.Payload)
	} else if len(msg.Payload) > 0 {
		q, err := url.ParseQuery(string(msg.Payload))
		if err != nil {
			return nil, err
		}
		uq := u.Query()
		for key, values := range q {
			for _, value := range values {
				uq.Add(key, value)
			}
		}
		u.RawQuery = uq.Encode()
	}

	request, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, err
	}
	msg.Header.Range(func(key string, values []string) {
		for _, value := range values {
			request.Header.Add(key, value)
		}
	})

	return request, nil
}

func (d DefaultMarshaller) Unmarshal(ctx context.Context, topic string, resp *http.Response) (*stream.Message, error) {
	headerMap := map[string][]string(resp.Header)
	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &stream.Message{
		Time:    time.Now(),
		Payload: payload,
		Header:  headerMap,
		Topic:   topic,
	}, nil
}
