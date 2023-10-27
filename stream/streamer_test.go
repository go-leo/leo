package stream_test

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"codeup.aliyun.com/qimao/leo/leo/stream"
	"codeup.aliyun.com/qimao/leo/leo/stream/gochan"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	goChan := make(chan *stream.Message, 1)
	publisher := gochan.NewPublisher("gochantest", goChan)
	go func() {
		for {
			time.Sleep(time.Millisecond)
			result, err := publisher.Publish(context.Background(), &stream.Message{Payload: []byte(time.Now().Format(time.DateTime))})
			assert.NoError(t, err)
			r := result.(stream.Results)
			for _, publishResult := range r {
				t.Log("result: " + string(publishResult.(*gochan.Result).Msg.Payload))
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
	originGoChan := make(chan *stream.Message, 1)
	publisher := gochan.NewPublisher("gochantest", originGoChan)
	go func() {
		for {
			time.Sleep(time.Millisecond)
			result, err := publisher.Publish(context.Background(), &stream.Message{Payload: []byte(time.Now().Format(time.DateTime))})
			assert.NoError(t, err)
			r := result.(stream.Results)
			for _, publishResult := range r {
				t.Log("result: " + string(publishResult.(*gochan.Result).Msg.Payload))
			}
		}
	}()
	middleGoChan := make(chan *stream.Message, 1)
	streamer := stream.NewStreamer(
		stream.PubSubHandlers(&GoChanPubSubHandler{SubGoChan: originGoChan, PubGoChan: middleGoChan}),
		stream.Handlers(&GoChanHandler{GoChan: middleGoChan}),
		stream.ErrorHandler(func(err error) { fmt.Println("error: ", err) }),
	)
	err := streamer.Run(context.Background())
	assert.NoError(t, err)
}

func TestInterceptor(t *testing.T) {
	goChan := make(chan *stream.Message, 1)
	publisher := gochan.NewPublisher("gochantest", goChan)
	go func() {
		for {
			time.Sleep(time.Millisecond)
			result, err := publisher.Publish(context.Background(), &stream.Message{Payload: []byte(time.Now().Format(time.DateTime))})
			assert.NoError(t, err)
			r := result.(stream.Results)
			for _, publishResult := range r {
				t.Log("result: " + string(publishResult.(*gochan.Result).Msg.Payload))
			}
		}
	}()
	streamer := stream.NewStreamer(
		stream.Handlers(&GoChanHandler{GoChan: goChan}),
		stream.Interceptors(
			func(ctx context.Context, msg *stream.Message, channel stream.Channel, invoker stream.Invoker) error {
				t.Log("Interceptor ==== 1 ==== before handle")
				err := invoker(ctx, msg, channel)
				t.Log("Interceptor ==== 1 ==== after handle")
				return err
			},

			func(ctx context.Context, msg *stream.Message, channel stream.Channel, invoker stream.Invoker) error {
				t.Log("Interceptor ==== 2 ==== before handle")
				err := invoker(ctx, msg, channel)
				t.Log("Interceptor ==== 2 ==== after handle")
				return err
			},

			func(ctx context.Context, msg *stream.Message, channel stream.Channel, invoker stream.Invoker) error {
				t.Log("Interceptor ==== 3 ==== before handle")
				err := invoker(ctx, msg, channel)
				t.Log("Interceptor ==== 3 ==== after handle")
				return err
			},
		),
		stream.ErrorHandler(func(err error) { fmt.Println("error: ", err) }),
	)
	err := streamer.Run(context.Background())
	assert.NoError(t, err)
}

func TestInterceptor_MiddleNotInvoker(t *testing.T) {
	goChan := make(chan *stream.Message, 1)
	publisher := gochan.NewPublisher("gochantest", goChan)
	go func() {
		for {
			time.Sleep(time.Millisecond)
			result, err := publisher.Publish(context.Background(), &stream.Message{Payload: []byte(time.Now().Format(time.DateTime))})
			assert.NoError(t, err)
			r := result.(stream.Results)
			for _, publishResult := range r {
				t.Log("result: " + string(publishResult.(*gochan.Result).Msg.Payload))
			}
		}
	}()
	streamer := stream.NewStreamer(
		stream.Handlers(&GoChanHandler{GoChan: goChan}),
		stream.Interceptors(
			func(ctx context.Context, msg *stream.Message, channel stream.Channel, invoker stream.Invoker) error {
				t.Log("Interceptor ==== 1 ==== before handle")
				err := invoker(ctx, msg, channel)
				t.Log("Interceptor ==== 1 ==== after handle")
				return err
			},

			func(ctx context.Context, msg *stream.Message, channel stream.Channel, invoker stream.Invoker) error {
				t.Log("Interceptor ==== 2 ====")
				return nil
			},

			func(ctx context.Context, msg *stream.Message, channel stream.Channel, invoker stream.Invoker) error {
				t.Log("Interceptor ==== 3 ==== before handle")
				err := invoker(ctx, msg, channel)
				t.Log("Interceptor ==== 3 ==== after handle")
				return err
			},
		),
		stream.ErrorHandler(func(err error) { fmt.Println("error: ", err) }),
	)
	err := streamer.Run(context.Background())
	assert.NoError(t, err)
}

func TestInterceptor_FinnalNotInvoker(t *testing.T) {
	goChan := make(chan *stream.Message, 1)
	publisher := gochan.NewPublisher("gochantest", goChan)
	go func() {
		for {
			time.Sleep(time.Millisecond)
			result, err := publisher.Publish(context.Background(), &stream.Message{Payload: []byte(time.Now().Format(time.DateTime))})
			assert.NoError(t, err)
			r := result.(stream.Results)
			for _, publishResult := range r {
				t.Log("result: " + string(publishResult.(*gochan.Result).Msg.Payload))
			}
		}
	}()
	streamer := stream.NewStreamer(
		stream.Handlers(&GoChanHandler{GoChan: goChan}),
		stream.Interceptors(
			func(ctx context.Context, msg *stream.Message, channel stream.Channel, invoker stream.Invoker) error {
				t.Log("Interceptor ==== 1 ==== before handle")
				err := invoker(ctx, msg, channel)
				t.Log("Interceptor ==== 1 ==== after handle")
				return err
			},

			func(ctx context.Context, msg *stream.Message, channel stream.Channel, invoker stream.Invoker) error {
				t.Log("Interceptor ==== 2 ==== before handle")
				err := invoker(ctx, msg, channel)
				t.Log("Interceptor ==== 2 ==== after handle")
				return err
			},

			func(ctx context.Context, msg *stream.Message, channel stream.Channel, invoker stream.Invoker) error {
				t.Log("Interceptor ==== 3 ====")
				return nil
			},
		),
		stream.ErrorHandler(func(err error) { fmt.Println("error: ", err) }),
	)
	err := streamer.Run(context.Background())
	assert.NoError(t, err)
}

type GoChanHandler struct {
	GoChan chan *stream.Message
}

func (h GoChanHandler) Subscriber() (stream.Subscriber, error) {
	return gochan.NewSubscriber("gochantest", h.GoChan), nil
}

func (h GoChanHandler) Handle(ctx context.Context, msg *stream.Message) error {
	fmt.Println("GoChanHandler handle:", string(msg.Payload))
	time.Sleep(time.Second)
	if rand.Intn(5) < 1 {
		return errors.New("mock error")
	}
	return nil
}

type GoChanPubSubHandler struct {
	SubGoChan chan *stream.Message
	PubGoChan chan *stream.Message
}

func (h GoChanPubSubHandler) Subscriber() (stream.Subscriber, error) {
	return gochan.NewSubscriber("gochantest", h.SubGoChan), nil
}

func (h GoChanPubSubHandler) Handle(ctx context.Context, msg *stream.Message) ([]*stream.Message, error) {
	fmt.Println("GoChanPubSubHandler handle:", string(msg.Payload))
	time.Sleep(time.Second)
	if rand.Intn(5) < 1 {
		return nil, errors.New("mock error")
	}
	return []*stream.Message{{Payload: []byte(string(msg.Payload) + "-1")}, {Payload: []byte(string(msg.Payload) + "-2")}}, nil
}

func (h GoChanPubSubHandler) Publisher() (stream.Publisher, error) {
	return gochan.NewPublisher("gochantest", h.PubGoChan), nil
}
