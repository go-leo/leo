package stream

import "context"

// Subscriber is message queue subscriber
type Subscriber interface {
	Topic() string
	Subscribe(ctx context.Context) (<-chan Message, <-chan error)
	Close(ctx context.Context) error
}

type SubscriberDecorator interface {
	Decorate(sub Subscriber) Subscriber
}

type SubscriberDecoratorFunc func(sub Subscriber) Subscriber

func (f SubscriberDecoratorFunc) Decorate(sub Subscriber) Subscriber {
	return f(sub)
}

func chainSubscriber(middlewares []SubscriberDecorator, publisher Subscriber) Subscriber {
	for i := len(middlewares) - 1; i >= 0; i-- {
		publisher = middlewares[i].Decorate(publisher)
	}
	return publisher
}
