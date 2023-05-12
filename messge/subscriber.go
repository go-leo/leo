package message

import "context"

// Subscriber is message queue subscriber
type Subscriber interface {
	Subscribe(ctx context.Context, topic string, handler HandlerFunc) <-chan error
	Close(ctx context.Context) error
}

type SubscriberMiddleware interface {
	Middleware(sub Subscriber) Subscriber
}

type SubscriberMiddlewareFunc func(sub Subscriber) Subscriber

func (f SubscriberMiddlewareFunc) Middleware(sub Subscriber) Subscriber {
	return f(sub)
}

func ChainSubscribers(sub Subscriber, mwf ...SubscriberMiddlewareFunc) Subscriber {
	for i := len(mwf) - 1; i >= 0; i-- {
		sub = mwf[i].Middleware(sub)
	}
	return sub
}
