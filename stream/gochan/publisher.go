package gochan

import (
	"codeup.aliyun.com/qimao/leo/leo/stream"
	"context"
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

	result := make([]*PublishResult, 0, len(messages))
	for _, msg := range messages {
		goChanMsg, err := pub.o.Marshaller.Marshal(ctx, pub.topic, msg)
		if err != nil {
			return nil, err
		}
		pub.goChan <- goChanMsg
		result = append(result, &PublishResult{Msg: goChanMsg})
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
	Msg []byte
}

func NewPublisher(topic string, goChan chan<- []byte, opts ...Option) *Publisher {
	o := &options{}
	o.apply(opts...)
	o.init()
	return &Publisher{
		o:      o,
		goChan: goChan,
		topic:  topic,
	}
}
