package confluentkafka

import (
	"context"
	"time"

	confluentkafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"

	message "codeup.aliyun.com/qimao/leo/leo/messge"
)

var _ message.Publisher = new(Publisher)

type Publisher struct {
	topic          string
	producer       *confluentkafka.Producer
	messageEncoder func(ctx context.Context, pub message.Publisher, msg *message.Message) (*confluentkafka.Message, error)
	flushTimeout   time.Duration
}

func (pub *Publisher) Topic() string {
	return pub.topic
}

func (pub *Publisher) Publish(ctx context.Context, topic string, messages ...*message.Message) (any, error) {
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

type PublisherBuilder struct {
	topic          string
	producer       *confluentkafka.Producer
	messageEncoder func(topic string, msg message.Message) (*confluentkafka.Message, error)
	flushTimeout   time.Duration
}

func (p *PublisherBuilder) SetTopic(topic string) *PublisherBuilder {
	p.topic = topic
	return p
}

func (p *PublisherBuilder) SetProducer(producer *confluentkafka.Producer) *PublisherBuilder {
	p.producer = producer
	return p
}

func (p *PublisherBuilder) SetMessageEncoder(messageEncoder func(topic string, msg message.Message) (*confluentkafka.Message, error)) *PublisherBuilder {
	p.messageEncoder = messageEncoder
	return p
}

func (p *PublisherBuilder) SetFlushTimeout(flushTimeout time.Duration) *PublisherBuilder {
	p.flushTimeout = flushTimeout
	return p
}

func (p *PublisherBuilder) Build() *Publisher {
	if p.messageEncoder == nil {
		p.messageEncoder = DefaultMessageEncoder
	}
	return &Publisher{
		topic:          p.topic,
		producer:       p.producer,
		messageEncoder: p.messageEncoder,
		flushTimeout:   p.flushTimeout,
	}
}

func NewPublisherBuilder() *PublisherBuilder {
	return &PublisherBuilder{}
}
