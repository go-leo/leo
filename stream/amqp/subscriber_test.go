package amqp_test

import (
	"codeup.aliyun.com/qimao/leo/leo/internal/gox/mathx/randx"
	"codeup.aliyun.com/qimao/leo/leo/stream"
	"codeup.aliyun.com/qimao/leo/leo/stream/amqp"
	"context"
	"github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestHelloWorldSubscriber(t *testing.T) {
	t.Parallel()
	factory := func(topic string) (*amqp091.Connection, error) {
		return amqp091.Dial("amqp://admin:admin123@localhost:5672/")
	}
	subscriber, err := amqp.NewSubscriber(
		"hello",
		factory,
		amqp.QueueName(func(topic string) string { return topic }),
		amqp.QueueOptions(&amqp.QueueOption{
			Durable:    false,
			AutoDelete: false,
			Exclusive:  false,
			NoWait:     false,
			Args:       nil,
		}),
		amqp.ConsumeOptions(amqp.ConsumeOption{
			Tag:       "",
			AutoAck:   true,
			Exclusive: false,
			NoLocal:   false,
			NoWait:    false,
			Table:     nil,
		}),
		amqp.NackHandler(func(msg *stream.Message) {
			t.Log("nack: ", string(msg.Payload))
		}),
	)
	assert.NoError(t, err)

	msgC := make(chan *stream.Message, 1)
	errC := make(chan error, 1)

	go func() {
		err := subscriber.Subscribe(context.Background(), msgC, errC)
		assert.NoError(t, err)
	}()

	go func() {
		for msg := range msgC {
			if randx.Intn(3) < 1 {
				ackRes, err := msg.Nack(context.Background())
				assert.NoError(t, err)
				t.Log(ackRes)
			} else {
				t.Log("ack msg: ", string(msg.Payload))
				ackRes, err := msg.Ack(context.Background())
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

func TestWorkerQueueSubscriber(t *testing.T) {
	t.Parallel()
	factory := func(topic string) (*amqp091.Connection, error) {
		return amqp091.Dial("amqp://admin:admin123@localhost:5672/")
	}
	subscriber, err := amqp.NewSubscriber(
		"task_queue",
		factory,
		amqp.QueueName(func(topic string) string { return topic }),
		amqp.QueueOptions(&amqp.QueueOption{
			Durable:    true,
			AutoDelete: false,
			Exclusive:  false,
			NoWait:     false,
			Args:       nil,
		}),
		amqp.QosOptions(&amqp.QosOption{
			PrefetchSize:  1,
			PrefetchCount: 0,
			Global:        false,
		}),
		amqp.ConsumeOptions(amqp.ConsumeOption{
			Tag:       "",
			AutoAck:   true,
			Exclusive: false,
			NoLocal:   false,
			NoWait:    false,
			Table:     nil,
		}),
		amqp.NackHandler(func(msg *stream.Message) {
			t.Log("nack: ", string(msg.Payload))
		}),
	)
	assert.NoError(t, err)

	msgC := make(chan *stream.Message, 1)
	errC := make(chan error, 1)

	go func() {
		err := subscriber.Subscribe(context.Background(), msgC, errC)
		assert.NoError(t, err)
	}()

	go func() {
		for msg := range msgC {
			if randx.Intn(3) < 1 {
				ackRes, err := msg.Nack(context.Background())
				assert.NoError(t, err)
				t.Log(ackRes)
			} else {
				t.Log("ack msg: ", string(msg.Payload))
				ackRes, err := msg.Ack(context.Background())
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

func TestPubSubSubscriber(t *testing.T) {
	t.Parallel()
	factory := func(topic string) (*amqp091.Connection, error) {
		return amqp091.Dial("amqp://admin:admin123@localhost:5672/")
	}
	subscriber, err := amqp.NewSubscriber(
		"logs",
		factory,
		amqp.ExchangeName(func(topic string) string {
			return topic
		}),
		amqp.QueueName(func(topic string) string {
			return ""
		}),
		amqp.ExchangeOptions(&amqp.ExchangeOption{
			Kind:       "fanout",
			Durable:    true,
			AutoDelete: false,
			Internal:   false,
			NoWait:     false,
			Args:       nil,
		}),
		amqp.QueueOptions(&amqp.QueueOption{
			Durable:    false,
			AutoDelete: false,
			Exclusive:  true,
			NoWait:     false,
			Args:       nil,
		}),
		amqp.QueueBindOptions(&amqp.QueueBindOption{
			NoWait: false,
			Args:   nil,
		}),
		amqp.ConsumeOptions(amqp.ConsumeOption{
			Tag:       "",
			AutoAck:   true,
			Exclusive: false,
			NoLocal:   false,
			NoWait:    false,
			Table:     nil,
		}),
		amqp.NackHandler(func(msg *stream.Message) {
			t.Log("nack: ", string(msg.Payload))
		}),
	)
	assert.NoError(t, err)

	msgC := make(chan *stream.Message, 1)
	errC := make(chan error, 1)

	go func() {
		err := subscriber.Subscribe(context.Background(), msgC, errC)
		assert.NoError(t, err)
	}()

	go func() {
		for msg := range msgC {
			if randx.Intn(3) < 1 {
				ackRes, err := msg.Nack(context.Background())
				assert.NoError(t, err)
				t.Log(ackRes)
			} else {
				t.Log("ack msg: ", string(msg.Payload))
				ackRes, err := msg.Ack(context.Background())
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

func TestRoutingSubscriber(t *testing.T) {
	t.Parallel()
	go func() {
		testRoutingSubscriber(t, []string{"warning", "error"})
	}()
	go func() {
		testRoutingSubscriber(t, []string{"info", "warning", "error"})
	}()
	select {}
}

func testRoutingSubscriber(t *testing.T, routingKeys []string) {
	factory := func(topic string) (*amqp091.Connection, error) {
		return amqp091.Dial("amqp://admin:admin123@localhost:5672/")
	}
	subscriber, err := amqp.NewSubscriber(
		"logs_direct",
		factory,
		amqp.ExchangeName(func(topic string) string { return topic }),
		amqp.QueueName(func(topic string) string { return "" }),
		amqp.RoutingKeys(func(topic string) []string { return routingKeys }),
		amqp.ExchangeOptions(&amqp.ExchangeOption{
			Kind:       "direct",
			Durable:    true,
			AutoDelete: false,
			Internal:   false,
			NoWait:     false,
			Args:       nil,
		}),
		amqp.QueueOptions(&amqp.QueueOption{
			Durable:    false,
			AutoDelete: false,
			Exclusive:  true,
			NoWait:     false,
			Args:       nil,
		}),
		amqp.QueueBindOptions(&amqp.QueueBindOption{
			NoWait: false,
			Args:   nil,
		}),
		amqp.ConsumeOptions(amqp.ConsumeOption{
			Tag:       "",
			AutoAck:   true,
			Exclusive: false,
			NoLocal:   false,
			NoWait:    false,
			Table:     nil,
		}),
		amqp.NackHandler(func(msg *stream.Message) {
			t.Log(strings.Join(routingKeys, ","), " nack: ", string(msg.Payload))
		}),
	)
	assert.NoError(t, err)

	msgC := make(chan *stream.Message, 1)
	errC := make(chan error, 1)

	go func() {
		err := subscriber.Subscribe(context.Background(), msgC, errC)
		assert.NoError(t, err)
	}()

	go func() {
		for msg := range msgC {
			if randx.Intn(3) < 1 {
				ackRes, err := msg.Nack(context.Background())
				assert.NoError(t, err)
				t.Log(ackRes)
			} else {
				t.Log(strings.Join(routingKeys, ","), "ack: ", string(msg.Payload))
				ackRes, err := msg.Ack(context.Background())
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

func TestTopicSubscriber(t *testing.T) {
	t.Parallel()
	go func() {
		testTopicSubscriber(t, []string{"#"})
	}()
	go func() {
		testTopicSubscriber(t, []string{"kern.*"})
	}()
	go func() {
		testTopicSubscriber(t, []string{"*.critical"})
	}()
	go func() {
		testTopicSubscriber(t, []string{"kern.*", "*.critical"})
	}()
	select {}
}

func testTopicSubscriber(t *testing.T, routingKeys []string) {
	factory := func(topic string) (*amqp091.Connection, error) {
		return amqp091.Dial("amqp://admin:admin123@localhost:5672/")
	}
	subscriber, err := amqp.NewSubscriber(
		"logs_topic",
		factory,
		amqp.ExchangeName(func(topic string) string { return topic }),
		amqp.QueueName(func(topic string) string { return "" }),
		amqp.RoutingKeys(func(topic string) []string { return routingKeys }),
		amqp.ExchangeOptions(&amqp.ExchangeOption{
			Kind:       "topic",
			Durable:    true,
			AutoDelete: false,
			Internal:   false,
			NoWait:     false,
			Args:       nil,
		}),
		amqp.QueueOptions(&amqp.QueueOption{
			Durable:    false,
			AutoDelete: false,
			Exclusive:  true,
			NoWait:     false,
			Args:       nil,
		}),
		amqp.QueueBindOptions(&amqp.QueueBindOption{
			NoWait: false,
			Args:   nil,
		}),
		amqp.ConsumeOptions(amqp.ConsumeOption{
			Tag:       "",
			AutoAck:   true,
			Exclusive: false,
			NoLocal:   false,
			NoWait:    false,
			Table:     nil,
		}),
		amqp.NackHandler(func(msg *stream.Message) {
			t.Log(strings.Join(routingKeys, ","), " nack: ", string(msg.Payload))
		}),
	)
	assert.NoError(t, err)

	msgC := make(chan *stream.Message, 1)
	errC := make(chan error, 1)

	go func() {
		err := subscriber.Subscribe(context.Background(), msgC, errC)
		assert.NoError(t, err)
	}()

	go func() {
		for msg := range msgC {
			if randx.Intn(3) < 1 {
				ackRes, err := msg.Nack(context.Background())
				assert.NoError(t, err)
				t.Log(ackRes)
			} else {
				t.Log(strings.Join(routingKeys, ","), "ack: ", string(msg.Payload))
				ackRes, err := msg.Ack(context.Background())
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
