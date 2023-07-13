package gochan

import (
	"context"
	"sync"
	"sync/atomic"

	"codeup.aliyun.com/qimao/leo/leo/stream"
)

var _ stream.Publisher = new(Publisher)

type Publisher struct {
	o      *options
	wg     sync.WaitGroup
	closed atomic.Bool
	topic  string
	goChan chan<- *stream.Message
}

func (pub *Publisher) Topic() string {
	return pub.topic
}

func (pub *Publisher) Queue() string {
	return "gochan"
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
		msg.Topic = pub.topic
		pub.goChan <- msg
		result = append(result, &Result{Msg: msg})
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

type Result struct {
	Msg *stream.Message
}

func (p Result) String() string {
	return "ok"
}

func NewPublisher(topic string, goChan chan<- *stream.Message, opts ...Option) *Publisher {
	o := &options{}
	o.apply(opts...)
	o.init()
	return &Publisher{
		o:      o,
		goChan: goChan,
		topic:  topic,
	}
}
