package gochan_test

import (
	"codeup.aliyun.com/qimao/leo/leo/stream"
	"codeup.aliyun.com/qimao/leo/leo/stream/gochan"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestKafkaPublisher(t *testing.T) {
	topic := "leo-stream-demo"

	goChan := make(chan []byte, 10)

	go func() {
		for msg := range goChan {
			t.Log("msg: ", string(msg))
		}
	}()

	publisher, err := gochan.NewPublisher(topic, goChan)
	assert.NoError(t, err)
	assert.Equal(t, topic, publisher.Topic())
	assert.Equal(t, "gochan", publisher.Queue())
	var messages = []*stream.Message{
		{
			ID:      "1",
			Time:    time.Now(),
			Payload: []byte("1"),
			Header:  stream.Header{"index": []string{"1"}},
		},
		{
			ID:      "2",
			Time:    time.Now(),
			Payload: []byte("2"),
			Header:  stream.Header{"index": []string{"2"}},
		},
		{
			ID:      "3",
			Time:    time.Now(),
			Payload: []byte("3"),
			Header:  stream.Header{"index": []string{"3"}},
		},
		{
			ID:      "4",
			Time:    time.Now(),
			Payload: []byte("4"),
			Header:  stream.Header{"index": []string{"4"}},
		},
		{
			ID:      "5",
			Time:    time.Now(),
			Payload: []byte("5"),
			Header:  stream.Header{"index": []string{"5"}},
		},
	}

	for _, m := range messages {
		publishResult, err := publisher.Publish(context.Background(), m)
		assert.NoError(t, err)
		t.Log(publishResult)
	}

}

func TestKafkaPublisherMultiMessage(t *testing.T) {
	topic := "leo-stream-demo"

	goChan := make(chan []byte, 10)

	go func() {
		for msg := range goChan {
			t.Log("msg: ", string(msg))
		}
	}()

	publisher, err := gochan.NewPublisher(topic, goChan)
	assert.NoError(t, err)
	assert.Equal(t, topic, publisher.Topic())
	assert.Equal(t, "gochan", publisher.Queue())

	var messages = []*stream.Message{
		{
			ID:      "1",
			Time:    time.Now(),
			Payload: []byte("1"),
			Header:  stream.Header{"index": []string{"1"}},
		},
		{
			ID:      "2",
			Time:    time.Now(),
			Payload: []byte("2"),
			Header:  stream.Header{"index": []string{"2"}},
		},
		{
			ID:      "3",
			Time:    time.Now(),
			Payload: []byte("3"),
			Header:  stream.Header{"index": []string{"3"}},
		},
		{
			ID:      "4",
			Time:    time.Now(),
			Payload: []byte("4"),
			Header:  stream.Header{"index": []string{"4"}},
		},
		{
			ID:      "5",
			Time:    time.Now(),
			Payload: []byte("5"),
			Header:  stream.Header{"index": []string{"5"}},
		},
	}

	publishResult, err := publisher.Publish(context.Background(), messages...)
	assert.NoError(t, err)
	t.Log(publishResult)

}
