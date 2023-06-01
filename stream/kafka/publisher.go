package kafka

import (
	"codeup.aliyun.com/qimao/leo/leo/stream"
	"context"
	"errors"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"sync"
	"sync/atomic"
)

var _ stream.Publisher = new(Publisher)

type Publisher struct {
	producer *kafka.Producer
	o        *options
	wg       sync.WaitGroup
	closed   atomic.Bool
	topic    string
}

func (pub *Publisher) Topic() string {
	return pub.topic
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

	deliveryChan := make(chan kafka.Event, 1)
	defer close(deliveryChan)

	result := make([]*PublishResult, 0, len(messages))
	for _, msg := range messages {
		kafkaMsg, err := pub.o.Marshaler.Marshal(ctx, pub.topic, msg)
		if err != nil {
			return nil, err
		}
		if err := pub.producer.Produce(kafkaMsg, deliveryChan); err != nil {
			return nil, err
		}
		event, ok := (<-deliveryChan).(*kafka.Message)
		if !ok {
			continue
		}
		produceResult := &PublishResult{
			Topic:     *event.TopicPartition.Topic,
			Partition: event.TopicPartition.Partition,
			Offset:    int64(event.TopicPartition.Offset),
			Msg:       msg,
		}
		result = append(result, produceResult)
	}
	return result, nil
}

func (pub *Publisher) Close(_ context.Context) error {
	if !pub.closed.CompareAndSwap(false, true) {
		return nil
	}
	pub.closed.Store(true)
	pub.wg.Wait()
	pub.producer.Flush(int(pub.o.ShutdownTimeout.Milliseconds()))
	pub.producer.Close()
	return nil
}

type PublishResult struct {
	Topic     string
	Partition int32
	Offset    int64
	Msg       *stream.Message
}

func NewPublisher(topic string, factory func() (*kafka.Producer, error), opts ...Option) (*Publisher, error) {
	if factory == nil {
		return nil, errors.New("factory is nil")
	}
	producer, err := factory()
	if err != nil {
		return nil, err
	}
	o := &options{}
	o.apply(opts...)
	o.init()
	return &Publisher{
		o:        o,
		producer: producer,
		topic:    topic,
	}, nil
}
