package stream_test

import (
	"codeup.aliyun.com/qimao/leo/leo/internal/gox/mathx/randx"
	"codeup.aliyun.com/qimao/leo/leo/stream"
	"codeup.aliyun.com/qimao/leo/leo/stream/gochan"
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	goChan := make(chan []byte, 1)
	publisher := gochan.NewPublisher("gochantest", goChan)
	go func() {
		for {
			time.Sleep(time.Millisecond)
			result, err := publisher.Publish(context.Background(), &stream.Message{Payload: []byte(time.Now().Format(time.DateTime))})
			assert.NoError(t, err)
			r := result.([]*gochan.PublishResult)
			for _, publishResult := range r {
				t.Log("result: " + string(publishResult.Msg))
			}
		}
	}()
	streamer := stream.NewStreamer(
		stream.Handlers(&GoChanHandler{GoChan: goChan}),
		stream.ErrorHandler(func(err error) { fmt.Println("error: ", err) }),
	)
	err := streamer.Run(context.Background())
	assert.NoError(t, err)
}

func TestPubSubHandler(t *testing.T) {
	originGoChan := make(chan []byte, 1)
	publisher := gochan.NewPublisher("gochantest", originGoChan)
	go func() {
		for {
			time.Sleep(time.Millisecond)
			result, err := publisher.Publish(context.Background(), &stream.Message{Payload: []byte(time.Now().Format(time.DateTime))})
			assert.NoError(t, err)
			r := result.([]*gochan.PublishResult)
			for _, publishResult := range r {
				t.Log("result: " + string(publishResult.Msg))
			}
		}
	}()
	middleGoChan := make(chan []byte, 1)
	streamer := stream.NewStreamer(
		stream.PubSubHandlers(&GoChanPubSubHandler{SubGoChan: originGoChan, PubGoChan: middleGoChan}),
		stream.Handlers(&GoChanHandler{GoChan: middleGoChan}),
		stream.ErrorHandler(func(err error) { fmt.Println("error: ", err) }),
	)
	err := streamer.Run(context.Background())
	assert.NoError(t, err)
}

type GoChanHandler struct {
	GoChan chan []byte
}

func (h GoChanHandler) Subscriber() (stream.Subscriber, error) {
	return gochan.NewSubscriber("gochantest", h.GoChan), nil
}

func (h GoChanHandler) Handle(ctx context.Context, msg *stream.Message) error {
	fmt.Println("GoChanHandler handle:", string(msg.Payload))
	time.Sleep(time.Second)
	if randx.Intn(5) < 1 {
		return errors.New("mock error")
	}
	return nil
}

type GoChanPubSubHandler struct {
	SubGoChan chan []byte
	PubGoChan chan []byte
}

func (h GoChanPubSubHandler) Subscriber() (stream.Subscriber, error) {
	return gochan.NewSubscriber("gochantest", h.SubGoChan), nil
}

func (h GoChanPubSubHandler) Handle(ctx context.Context, msg *stream.Message) ([]*stream.Message, error) {
	fmt.Println("GoChanPubSubHandler handle:", string(msg.Payload))
	time.Sleep(time.Second)
	if randx.Intn(5) < 1 {
		return nil, errors.New("mock error")
	}
	return []*stream.Message{{Payload: []byte(string(msg.Payload) + "-1")}, {Payload: []byte(string(msg.Payload) + "-2")}}, nil
}

func (h GoChanPubSubHandler) Publisher() (stream.Publisher, error) {
	return gochan.NewPublisher("gochantest", h.PubGoChan), nil
}
