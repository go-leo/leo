package amqp

import (
	"context"
	"strings"
	"time"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/mathx/randx"
	"codeup.aliyun.com/qimao/leo/leo/stream"

	"github.com/rabbitmq/amqp091-go"
)

const (
	ContentTypeKey     = "Content-Type"
	ContentEncodingKey = "Content-Encoding"
	PriorityKey        = "Priority"
	CorrelationIdKey   = "CorrelationId"
	ReplyToKey         = "ReplyTo"
	ExpirationKey      = "Expiration"
	TypeKey            = "Type"
	UserIdKey          = "User-ID"
	AppIdKey           = "App-ID"
)

// Marshaller marshals stream's message to *kafka.Message and unmarshals *kafka.Message to stream's message.
type Marshaller interface {
	Marshal(ctx context.Context, topic string, msg *stream.Message) (amqp091.Publishing, error)
	Unmarshal(amqpMsg amqp091.Delivery) (*stream.Message, error)
}

type DefaultMarshaller struct{}

func (d DefaultMarshaller) Marshal(ctx context.Context, topic string, msg *stream.Message) (amqp091.Publishing, error) {
	if len(msg.ID) == 0 {
		msg.ID = randx.WordString(32)
	}
	if msg.Time.IsZero() {
		msg.Time = time.Now()
	}
	headers := amqp091.Table{}
	msg.Header.Range(func(key string, values []string) {
		headers[key] = strings.Join(values, ", ")
	})
	return amqp091.Publishing{
		Headers:   headers,
		MessageId: msg.ID,
		Timestamp: msg.Time,
		Body:      msg.Payload,
	}, nil
}

func (d DefaultMarshaller) Unmarshal(amqpMsg amqp091.Delivery) (*stream.Message, error) {
	if len(amqpMsg.MessageId) == 0 {
		amqpMsg.MessageId = randx.WordString(32)
	}
	return &stream.Message{
		ID:      amqpMsg.MessageId,
		Time:    amqpMsg.Timestamp,
		Payload: amqpMsg.Body,
		Header:  nil,
	}, nil
}
