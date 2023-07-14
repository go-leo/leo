package amqp

import (
	"context"
	"strings"
	"time"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/convx"

	"codeup.aliyun.com/qimao/leo/leo/stream"

	"github.com/rabbitmq/amqp091-go"
)

// Marshaller marshals stream's message to *kafka.Message and unmarshals *kafka.Message to stream's message.
type Marshaller interface {
	Marshal(ctx context.Context, topic string, msg *stream.Message) (amqp091.Publishing, error)
	Unmarshal(topic string, amqpMsg amqp091.Delivery) (*stream.Message, error)
}

var _ Marshaller = (*DefaultMarshaller)(nil)

type DefaultMarshaller struct{}

func (d DefaultMarshaller) Marshal(ctx context.Context, topic string, msg *stream.Message) (amqp091.Publishing, error) {
	msg.Topic = topic
	if msg.Time.IsZero() {
		msg.Time = time.Now()
	}
	headers := amqp091.Table{}
	msg.Header.Range(func(key string, values []string) {
		headers[key] = strings.Join(values, ", ")
	})
	return amqp091.Publishing{
		Headers:   headers,
		Timestamp: msg.Time,
		Body:      msg.Payload,
	}, nil
}

func (d DefaultMarshaller) Unmarshal(topic string, amqpMsg amqp091.Delivery) (*stream.Message, error) {
	header := stream.Header{}
	for key, value := range amqpMsg.Headers {
		header.Add(key, convx.ToString(value))
	}
	return &stream.Message{
		Time:    amqpMsg.Timestamp,
		Payload: amqpMsg.Body,
		Header:  header,
		Topic:   topic,
	}, nil
}
