package kafka

import (
	"codeup.aliyun.com/qimao/leo/leo/internal/gox/operator"
	"codeup.aliyun.com/qimao/leo/leo/internal/gox/stringx"
	"codeup.aliyun.com/qimao/leo/leo/stream"
	"context"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"strconv"
	"time"
)

const (
	KeyKey           = "Key"
	PartitionKey     = "Partition"
	OffsetKey        = "Offset"
	TimestampKey     = "Timestamp"
	TimestampTypeKey = "TimestampType"
	OpaqueKey        = "Opaque"
	MetadataKey      = "Metadata"
	TopicKey         = "Topic"
	ErrorKey         = "Error"
)

// Marshaler marshals stream's message to *kafka.Message and unmarshals *kafka.Message to stream's message.
type Marshaler interface {
	Marshal(ctx context.Context, topic string, msg *stream.Message) (*kafka.Message, error)
	Unmarshal(kafkaMsg *kafka.Message) (*stream.Message, error)
}

type DefaultMarshaler struct {
}

func (d DefaultMarshaler) Marshal(ctx context.Context, topic string, msg *stream.Message) (*kafka.Message, error) {
	header := msg.Header

	// partition
	partition := kafka.PartitionAny
	if v := header.Get(PartitionKey); stringx.IsNotBlank(v) {
		intV, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		partition = int32(intV)
	}
	header.Del(PartitionKey)

	// key
	var key string
	if v := header.Get(KeyKey); stringx.IsNotBlank(v) {
		key = v
	}
	header.Del(KeyKey)

	// opaque
	opaque := header.Get(OpaqueKey)
	header.Del(OpaqueKey)

	// other header
	kafkaHeaders := make([]kafka.Header, 0, header.Len())
	header.Range(func(key string, values []string) {
		for _, value := range values {
			kafkaHeaders = append(kafkaHeaders, kafka.Header{Key: key, Value: []byte(value)})
		}
	})

	kafkaMsg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: partition,
		},
		Value:         msg.Payload,
		Key:           []byte(key),
		Timestamp:     operator.Ternary(!msg.Time.IsZero(), msg.Time, time.Now()),
		TimestampType: kafka.TimestampCreateTime,
		Opaque:        opaque,
		Headers:       kafkaHeaders,
	}
	return kafkaMsg, nil
}

func (d DefaultMarshaler) Unmarshal(kafkaMsg *kafka.Message) (*stream.Message, error) {
	//TODO implement me
	panic("implement me")
}
