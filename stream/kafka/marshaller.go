package kafka

import (
	"context"
	"time"

	"codeup.aliyun.com/qimao/leo/leo/stream"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const (
	KeyKey       = "Key"
	PartitionKey = "Partition"
	OpaqueKey    = "Opaque"
)

// Marshaller marshals stream's message to *kafka.Message and unmarshals *kafka.Message to stream's message.
type Marshaller interface {
	Marshal(ctx context.Context, topic string, msg *stream.Message) (*kafka.Message, error)
	Unmarshal(topic string, kafkaMsg *kafka.Message) (*stream.Message, error)
}

var _ Marshaller = (*DefaultMarshaller)(nil)

type DefaultMarshaller struct{}

func (d DefaultMarshaller) Marshal(ctx context.Context, topic string, msg *stream.Message) (*kafka.Message, error) {
	msg.Topic = topic

	if msg.Time.IsZero() {
		msg.Time = time.Now()
	}

	kafkaHeaders := make([]kafka.Header, 0, msg.Header.Len())
	msg.Header.Range(func(key string, values []string) {
		for _, value := range values {
			kafkaHeaders = append(kafkaHeaders, kafka.Header{Key: key, Value: []byte(value)})
		}
	})

	kafkaMsg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value:         msg.Payload,
		Timestamp:     msg.Time,
		TimestampType: kafka.TimestampCreateTime,
		Headers:       kafkaHeaders,
	}
	return kafkaMsg, nil
}

func (d DefaultMarshaller) Unmarshal(topic string, kafkaMsg *kafka.Message) (*stream.Message, error) {
	header := stream.Header{}
	for _, h := range kafkaMsg.Headers {
		header.Add(h.Key, string(h.Value))
	}
	return &stream.Message{
		Time:    kafkaMsg.Timestamp,
		Payload: kafkaMsg.Value,
		Header:  header,
		Topic:   topic,
	}, nil
}
