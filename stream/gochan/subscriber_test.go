package gochan_test

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"codeup.aliyun.com/qimao/leo/leo/stream"
	"codeup.aliyun.com/qimao/leo/leo/stream/gochan"

	"github.com/stretchr/testify/assert"
)

func TestSubscriber(t *testing.T) {
	topic := "leo-stream-demo"
	goChan := make(chan *stream.Message, 10)

	go func() {
		for {
			time.Sleep(time.Millisecond)
			goChan <- &stream.Message{Payload: []byte(time.Now().String())}
		}
	}()

	subscriber := gochan.NewSubscriber(topic, goChan, gochan.NackHandler(func(msg *stream.Message) {
		t.Log("nack msg: ", string(msg.Payload))
	}))
	assert.Equal(t, topic, subscriber.Topic())
	assert.Equal(t, "gochan", subscriber.Queue())
	msgC := make(chan *stream.Message, 1)
	errC := make(chan error, 1)

	go func() {
		err := subscriber.Subscribe(context.Background(), msgC, errC)
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
