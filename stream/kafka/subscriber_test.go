package kafka_test

import (
	"context"
	"testing"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/mathx/randx"
	"codeup.aliyun.com/qimao/leo/leo/stream"
	"codeup.aliyun.com/qimao/leo/leo/stream/kafka"

	kafka2 "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/stretchr/testify/assert"
)

func TestSubscriber(t *testing.T) {
	topic := "leo-stream-demo"
	subscriber, err := kafka.NewSubscriber(
		topic,
		func() (*kafka2.Consumer, error) {
			return kafka2.NewConsumer(&kafka2.ConfigMap{
				"api.version.request":       "true",
				"auto.offset.reset":         "latest",
				"heartbeat.interval.ms":     3000,
				"session.timeout.ms":        30000,
				"max.poll.interval.ms":      120000,
				"fetch.max.bytes":           1024000,
				"max.partition.fetch.bytes": 256000,
				"bootstrap.servers":         "localhost:9092",
				"group.id":                  "TestSubscriber",
			})
		},
		kafka.NackHandler(func(msg *stream.Message) {
			t.Log("nack msg: ", string(msg.Payload))
		}))
	assert.NoError(t, err)
	assert.Equal(t, topic, subscriber.Topic())
	assert.Equal(t, "kafka", subscriber.Queue())
	msgC := make(chan *stream.Message, 1)
	errC := make(chan error, 1)

	go func() {
		err = subscriber.Subscribe(context.Background(), msgC, errC)
		assert.NoError(t, err)
	}()

	go func() {
		for msg := range msgC {
			if randx.Intn(3) < 1 {
				t.Log("ack msg: ", string(msg.Payload))
				ackRes, err := msg.Ack(context.Background())
				assert.NoError(t, err)
				t.Log(ackRes)
			} else {
				ackRes, err := msg.Nack(context.Background())
				assert.NoError(t, err)
				t.Log(ackRes)
			}
		}
	}()

	go func() {
		for err := range errC {
			t.Log(err)
		}
	}()

	select {}
}
