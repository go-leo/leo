package pubsub

import "github.com/ThreeDotsLabs/watermill/message"

type Job struct {
	name                 string
	subscribeTopic       string
	subscriber           message.Subscriber
	publishTopic         string
	publisher            message.Publisher
	handlerFunc          func(msg *message.Message) ([]*message.Message, error)
	noPublishHandlerFunc func(msg *message.Message) error
	handler              *message.Handler
	middlewares          []message.HandlerMiddleware
}

func NewPubSubJob(
	name string,
	subscribeTopic string,
	subscriber message.Subscriber,
	publishTopic string,
	publisher message.Publisher,
	handlerFunc func(msg *message.Message) ([]*message.Message, error),
	middlewares ...message.HandlerMiddleware,
) *Job {
	return &Job{
		name:           name,
		subscribeTopic: subscribeTopic,
		subscriber:     subscriber,
		publishTopic:   publishTopic,
		publisher:      publisher,
		handlerFunc:    handlerFunc,
		middlewares:    middlewares,
	}
}

func NewSubJob(
	name string,
	subscribeTopic string,
	subscriber message.Subscriber,
	handlerFunc func(msg *message.Message) error,
	middlewares ...message.HandlerMiddleware,
) *Job {
	return &Job{
		name:                 name,
		subscribeTopic:       subscribeTopic,
		subscriber:           subscriber,
		noPublishHandlerFunc: handlerFunc,
		middlewares:          middlewares,
	}
}

func (j Job) Name() string {
	return j.name
}

func (j Job) Subscriber() message.Subscriber {
	return j.subscriber
}

func (j Job) SubscribeTopic() string {
	return j.subscribeTopic
}

func (j Job) Publisher() message.Publisher {
	return j.publisher
}

func (j Job) PublishTopic() string {
	return j.publishTopic
}
