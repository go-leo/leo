package http_test

import (
	httpstream "codeup.aliyun.com/qimao/leo/leo/stream/http"
	"context"
	"net/http"
	"testing"
	"time"

	"codeup.aliyun.com/qimao/leo/leo/stream"
	"github.com/stretchr/testify/assert"
)

func TestPublisherGet(t *testing.T) {
	publisher, err := httpstream.NewPublisher(http.MethodGet, "http://httpbin.org/get")
	assert.NoError(t, err)
	assert.Equal(t, "http", publisher.Queue())

	messages := []*stream.Message{
		{
			Time:    time.Now(),
			Payload: []byte("number=1"),
			Header:  stream.Header{"index": []string{"1"}},
		},
		{
			Time:    time.Now(),
			Payload: []byte("number=2"),
			Header:  stream.Header{"index": []string{"2"}},
		},
		{
			Time:    time.Now(),
			Payload: []byte("number=3"),
			Header:  stream.Header{"index": []string{"3"}},
		},
		{
			Time:    time.Now(),
			Payload: []byte("number=4"),
			Header:  stream.Header{"index": []string{"4"}},
		},
		{
			Time:    time.Now(),
			Payload: []byte("number=5"),
			Header:  stream.Header{"index": []string{"5"}},
		},
	}

	for _, m := range messages {
		publishResult, err := publisher.Publish(context.Background(), m)
		assert.NoError(t, err)
		t.Log(publishResult)
	}
}

func TestPublisherPost(t *testing.T) {
	publisher, err := httpstream.NewPublisher(http.MethodPost, "http://httpbin.org/post")
	assert.NoError(t, err)
	assert.Equal(t, "http", publisher.Queue())

	messages := []*stream.Message{
		{
			Time:    time.Now(),
			Payload: []byte("number=1"),
			Header:  stream.Header{"index": []string{"1"}},
		},
		{
			Time:    time.Now(),
			Payload: []byte("number=2"),
			Header:  stream.Header{"index": []string{"2"}},
		},
		{
			Time:    time.Now(),
			Payload: []byte("number=3"),
			Header:  stream.Header{"index": []string{"3"}},
		},
		{
			Time:    time.Now(),
			Payload: []byte("number=4"),
			Header:  stream.Header{"index": []string{"4"}},
		},
		{
			Time:    time.Now(),
			Payload: []byte("number=5"),
			Header:  stream.Header{"index": []string{"5"}},
		},
	}

	for _, m := range messages {
		publishResult, err := publisher.Publish(context.Background(), m)
		assert.NoError(t, err)
		t.Log(publishResult)
	}
}
