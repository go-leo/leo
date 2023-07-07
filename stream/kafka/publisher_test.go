package kafka_test

import (
	"codeup.aliyun.com/qimao/leo/leo/stream"
	"codeup.aliyun.com/qimao/leo/leo/stream/kafka"
	"context"
	kafka2 "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestKafkaPublisher(t *testing.T) {
	topic := "leo-stream-demo"
	publisher, err := kafka.NewPublisher(topic, func() (*kafka2.Producer, error) {
		return kafka2.NewProducer(&kafka2.ConfigMap{
			"api.version.request":           "true",
			"message.max.bytes":             1000000,
			"linger.ms":                     500,
			"sticky.partitioning.linger.ms": 1000,
			"retries":                       10,
			"retry.backoff.ms":              1000,
			"acks":                          "1",
			"bootstrap.servers":             "localhost:9092",
		})
	})
	assert.NoError(t, err)
	assert.Equal(t, topic, publisher.Topic())
	assert.Equal(t, "kafka", publisher.Queue())

	var messages = []*stream.Message{
		{
			id:      "1",
			Time:    time.Now(),
			Payload: []byte("1"),
			Header:  stream.Header{"index": []string{"1"}},
		},
		{
			id:      "2",
			Time:    time.Now(),
			Payload: []byte("2"),
			Header:  stream.Header{"index": []string{"2"}},
		},
		{
			id:      "3",
			Time:    time.Now(),
			Payload: []byte("3"),
			Header:  stream.Header{"index": []string{"3"}},
		},
		{
			id:      "4",
			Time:    time.Now(),
			Payload: []byte("4"),
			Header:  stream.Header{"index": []string{"4"}},
		},
		{
			id:      "5",
			Time:    time.Now(),
			Payload: []byte("5"),
			Header:  stream.Header{"index": []string{"5"}},
		},
	}

	for _, m := range messages {
		publishResult, err := publisher.Publish(context.Background(), m)
		assert.NoError(t, err)
		t.Log(publishResult)
	}

}

func TestKafkaPublisherMultiMessage(t *testing.T) {
	topic := "leo-stream-demo"
	publisher, err := kafka.NewPublisher(topic, func() (*kafka2.Producer, error) {
		return kafka2.NewProducer(&kafka2.ConfigMap{
			"api.version.request":           "true",
			"message.max.bytes":             1000000,
			"linger.ms":                     500,
			"sticky.partitioning.linger.ms": 1000,
			"retries":                       10,
			"retry.backoff.ms":              1000,
			"acks":                          "1",
			"bootstrap.servers":             "localhost:9092",
		})
	})
	assert.NoError(t, err)
	assert.Equal(t, topic, publisher.Topic())
	assert.Equal(t, "kafka", publisher.Queue())

	var messages = []*stream.Message{
		{
			id:      "1",
			Time:    time.Now(),
			Payload: []byte("1"),
			Header:  stream.Header{"index": []string{"1"}},
		},
		{
			id:      "2",
			Time:    time.Now(),
			Payload: []byte("2"),
			Header:  stream.Header{"index": []string{"2"}},
		},
		{
			id:      "3",
			Time:    time.Now(),
			Payload: []byte("3"),
			Header:  stream.Header{"index": []string{"3"}},
		},
		{
			id:      "4",
			Time:    time.Now(),
			Payload: []byte("4"),
			Header:  stream.Header{"index": []string{"4"}},
		},
		{
			id:      "5",
			Time:    time.Now(),
			Payload: []byte("5"),
			Header:  stream.Header{"index": []string{"5"}},
		},
	}

	publishResult, err := publisher.Publish(context.Background(), messages...)
	assert.NoError(t, err)
	t.Log(publishResult)

}
