package amqp_test

import (
	"codeup.aliyun.com/qimao/leo/leo/stream"
	"codeup.aliyun.com/qimao/leo/leo/stream/amqp"
	"context"
	"github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestHelloWorldPublisher(t *testing.T) {
	factory := func(topic string) (*amqp091.Connection, error) {
		return amqp091.Dial("amqp://admin:admin123@localhost:5672/")
	}
	publisher, err := amqp.NewPublisher(
		"hello",
		factory,
		amqp.QueueName(func(topic string) string {
			return topic
		}),
		amqp.RoutingKeys(func(topic string) []string {
			return []string{topic}
		}),
		amqp.QueueOptions(&amqp.QueueOption{
			Durable:    false,
			AutoDelete: false,
			Exclusive:  false,
			NoWait:     false,
			Args:       nil,
		}),
	)
	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	publishRes, err := publisher.Publish(ctx, &stream.Message{Payload: []byte("Hello World!")})
	assert.NoError(t, err)
	t.Log(publishRes)
}

func TestConfirmPublisher(t *testing.T) {
	factory := func(topic string) (*amqp091.Connection, error) {
		return amqp091.Dial("amqp://admin:admin123@localhost:5672/")
	}
	publisher, err := amqp.NewPublisher(
		"hello",
		factory,
		amqp.QueueName(func(topic string) string {
			return topic
		}),
		amqp.RoutingKeys(func(topic string) []string {
			return []string{topic}
		}),
		amqp.QueueOptions(&amqp.QueueOption{
			Durable:    false,
			AutoDelete: false,
			Exclusive:  false,
			NoWait:     false,
			Args:       nil,
		}),
		amqp.ConfirmOptions(&amqp.ConfirmOption{}),
	)
	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	publishRes, err := publisher.Publish(ctx, &stream.Message{Payload: []byte("Hello World!")})
	assert.NoError(t, err)
	t.Log(publishRes)
}

func TestWorkerQueuePublisher(t *testing.T) {
	factory := func(topic string) (*amqp091.Connection, error) {
		return amqp091.Dial("amqp://admin:admin123@localhost:5672/")
	}
	publisher, err := amqp.NewPublisher(
		"task_queue",
		factory,
		amqp.QueueName(func(topic string) string {
			return topic
		}),
		amqp.RoutingKeys(func(topic string) []string {
			return []string{topic}
		}),
		amqp.QueueOptions(&amqp.QueueOption{
			Durable:    true,
			AutoDelete: false,
			Exclusive:  false,
			NoWait:     false,
			Args:       nil,
		}),
	)
	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	publishRes, err := publisher.Publish(ctx, &stream.Message{Payload: []byte("Hello World!")})
	assert.NoError(t, err)
	t.Log(publishRes)

}

func TestPubSubPublisher(t *testing.T) {
	factory := func(topic string) (*amqp091.Connection, error) {
		return amqp091.Dial("amqp://admin:admin123@localhost:5672/")
	}
	publisher, err := amqp.NewPublisher(
		"logs",
		factory,
		amqp.ExchangeName(func(topic string) string {
			return topic
		}),
		amqp.ExchangeOptions(&amqp.ExchangeOption{
			Kind:       "fanout",
			Durable:    true,
			AutoDelete: false,
			Internal:   false,
			NoWait:     false,
			Args:       nil,
		}),
	)
	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	publishRes, err := publisher.Publish(ctx, &stream.Message{Payload: []byte("Hello World!")})
	assert.NoError(t, err)
	t.Log(publishRes)

}

func TestRoutingPublisher(t *testing.T) {
	testRoutingPublisher(t, "info")
	testRoutingPublisher(t, "warning")
	testRoutingPublisher(t, "error")
}

func testRoutingPublisher(t *testing.T, routingKey string) {
	factory := func(topic string) (*amqp091.Connection, error) {
		return amqp091.Dial("amqp://admin:admin123@localhost:5672/")
	}
	publisher, err := amqp.NewPublisher(
		"logs_direct",
		factory,
		amqp.RoutingKeys(func(topic string) []string {
			return []string{routingKey}
		}),
		amqp.ExchangeName(func(topic string) string { return topic }),
		amqp.ExchangeOptions(&amqp.ExchangeOption{
			Kind:       "direct",
			Durable:    true,
			AutoDelete: false,
			Internal:   false,
			NoWait:     false,
			Args:       nil,
		}),
	)
	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	publishRes, err := publisher.Publish(ctx, &stream.Message{Payload: []byte("Hello World! " + routingKey)})
	assert.NoError(t, err)
	t.Log(publishRes)
}

func TestTopicPublisher(t *testing.T) {
	testTopicPublisher(t, "kern.critical")
	testTopicPublisher(t, "kern.error")
	testTopicPublisher(t, "app.error")
}

func testTopicPublisher(t *testing.T, routingKey string) {
	factory := func(topic string) (*amqp091.Connection, error) {
		return amqp091.Dial("amqp://admin:admin123@localhost:5672/")
	}
	publisher, err := amqp.NewPublisher(
		"logs_topic",
		factory,
		amqp.RoutingKeys(func(topic string) []string {
			return []string{routingKey}
		}),
		amqp.ExchangeName(func(topic string) string { return topic }),
		amqp.ExchangeOptions(&amqp.ExchangeOption{
			Kind:       "topic",
			Durable:    true,
			AutoDelete: false,
			Internal:   false,
			NoWait:     false,
			Args:       nil,
		}),
	)
	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	publishRes, err := publisher.Publish(ctx, &stream.Message{Payload: []byte("Hello World! " + routingKey)})
	assert.NoError(t, err)
	t.Log(publishRes)
}
