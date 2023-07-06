package gochan

import (
	"codeup.aliyun.com/qimao/leo/leo/stream"
	"context"
	"errors"
	"sync"
	"sync/atomic"
)

var _ stream.Publisher = new(Publisher)

type Publisher struct {
	o      *options
	wg     sync.WaitGroup
	closed atomic.Bool
	topic  string
	goChan chan<- []byte
}

func (pub *Publisher) Topic() string {
	return pub.topic
}

func (pub *Publisher) Queue() string {
	return "gochan"
}

func (pub *Publisher) Publish(ctx context.Context, messages ...*stream.Message) (any, error) {
	if len(messages) == 0 {
		return nil, nil
	}
	if pub.closed.Load() {
		return nil, stream.ErrPublisherClosed
	}

	pub.wg.Add(1)
	defer pub.wg.Done()

	result := make([]string, 0, len(messages))
	for _, msg := range messages {
		kafkaMsg, err := pub.o.Marshaler.Marshal(ctx, pub.topic, msg)
		if err != nil {
			return nil, err
		}
		pub.goChan <- kafkaMsg
		result = append(result, "ok")
	}
	return result, nil
}

func (pub *Publisher) Close(_ context.Context) error {
	if !pub.closed.CompareAndSwap(false, true) {
		return nil
	}
	pub.closed.Store(true)
	pub.wg.Wait()
	return nil
}

func NewPublisher(topic string, goChan chan<- []byte, opts ...Option) (*Publisher, error) {
	if goChan == nil {
		return nil, errors.New("goChan is nil")
	}
	o := &options{}
	o.apply(opts...)
	o.init()
	return &Publisher{
		o:      o,
		goChan: goChan,
		topic:  topic,
	}, nil
}
