package kafka

import (
	"codeup.aliyun.com/qimao/leo/leo/internal/gox/mathx/randx"
	"codeup.aliyun.com/qimao/leo/leo/internal/gox/stringx"
	"codeup.aliyun.com/qimao/leo/leo/stream"
	"context"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"strconv"
	"time"
)

const (
	KeyKey       = "Key"
	PartitionKey = "Partition"
	OpaqueKey    = "Opaque"
	IDKey        = "X-Leo-Stream-ID"
)

// Marshaller marshals stream's message to *kafka.Message and unmarshals *kafka.Message to stream's message.
type Marshaller interface {
	Marshal(ctx context.Context, topic string, msg *stream.Message) (*kafka.Message, error)
	Unmarshal(kafkaMsg *kafka.Message) (*stream.Message, error)
}

type DefaultMarshaller struct{}

func (d DefaultMarshaller) Marshal(ctx context.Context, topic string, msg *stream.Message) (*kafka.Message, error) {
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

	if len(msg.ID) <= 0 {
		kafkaHeaders = append(kafkaHeaders, kafka.Header{Key: IDKey, Value: []byte(randx.WordString(32))})
	}

	if msg.Time.IsZero() {
		msg.Time = time.Now()
	}

	kafkaMsg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: partition,
		},
		Value:         msg.Payload,
		Key:           []byte(key),
		Timestamp:     msg.Time,
		TimestampType: kafka.TimestampCreateTime,
		Opaque:        opaque,
		Headers:       kafkaHeaders,
	}
	return kafkaMsg, nil
}

func (d DefaultMarshaller) Unmarshal(kafkaMsg *kafka.Message) (*stream.Message, error) {
	var id string
	header := stream.Header{}
	for _, h := range kafkaMsg.Headers {
		if h.Key == IDKey {
			id = string(h.Value)
			continue
		}
		header.Add(h.Key, string(h.Value))
	}
	if len(id) <= 0 {
		id = randx.WordString(32)
	} else {

	}
	return &stream.Message{
		ID:      id,
		Time:    kafkaMsg.Timestamp,
		Payload: kafkaMsg.Value,
		Header:  header,
	}, nil
}
