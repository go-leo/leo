package kafka

import (
	"codeup.aliyun.com/qimao/leo/leo/stream"
	"context"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var _ stream.Publisher = new(Publisher)

type Publisher struct {
	producer     *kafka.Producer
	flushTimeout time.Duration
}

func (pub *Publisher) Topic() string {
	return pub.topic
}

func (pub *Publisher) Publish(ctx context.Context, topic string, messages ...*stream.Message) (any, error) {
	var errs []error
	deliveryChan := make(chan confluentkafka.Event, len(messages))
	defer close(deliveryChan)

	for _, m := range messages {
		kafkaMsg, err := pub.messageEncoder(ctx, pub, m)
		if err != nil {
			return nil, err
		}
		if err := pub.producer.Produce(kafkaMsg, deliveryChan); err != nil {
			return nil, err
		}
	}

	event := <-deliveryChan
	if msg, ok := event.(*confluentkafka.Message); ok {
		if msg.TopicPartition.Error != nil {
			errs = append(errs, msg.TopicPartition.Error)
		} else {

		}
		return event, nil
	}

	return msg, nil
}

func (pub *Publisher) Close(_ context.Context) error {
	pub.producer.Flush(int(pub.flushTimeout.Milliseconds()))
	pub.producer.Close()
	return nil
}
