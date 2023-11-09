package http_test

import (
	"codeup.aliyun.com/qimao/leo/leo/stream"
	httpstream "codeup.aliyun.com/qimao/leo/leo/stream/http"
	"context"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

func TestSubscriber(t *testing.T) {
	topic := "leo-stream-httpstream-get-demo"
	subscriber, err := httpstream.NewSubscriber(
		topic,
		http.MethodPost,
		"/post",
		httpstream.HttpServer(&http.Server{Addr: ":8080"}),
	)
	assert.NoError(t, err)
	assert.Equal(t, topic, subscriber.Topic())
	assert.Equal(t, "http", subscriber.Queue())
	msgC := make(chan *stream.Message, 1)
	errC := make(chan error, 1)

	go func() {
		err = subscriber.Subscribe(context.Background(), msgC, errC)
		assert.NoError(t, err)
	}()

	go func() {
		for msg := range msgC {
			if rand.Intn(3) < 1 {
				t.Log("ack msg: ", string(msg.Payload))
				ackRes, err := msg.Ack(context.Background())
				assert.NoError(t, err)
				t.Log(ackRes)
			} else {
				ackRes, err := msg.Nack(context.Background())
				assert.NoError(t, err)
				t.Log(ackRes)
			}
		}
	}()

	go func() {
		for err := range errC {
			t.Log(err)
		}
	}()

	select {}
}

func TestSubscriberPublisher(t *testing.T) {
	topic := "leo-stream-httpstream-get-demo"
	publisher := httpstream.NewPublisher(topic, http.MethodPost, "http://localhost:8080/post")
	assert.Equal(t, topic, publisher.Topic())
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
