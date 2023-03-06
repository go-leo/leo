package stream

import "context"

type Publisher interface {
	Topic() string
	Publish(message Message) (any, error)
	Close(ctx context.Context) error
}

type PublisherDecorator interface {
	Decorate(pub Publisher) Publisher
}

type PublisherDecoratorFunc func(pub Publisher) Publisher

func (f PublisherDecoratorFunc) Decorate(pub Publisher) Publisher {
	return f(pub)
}

func chainPublisher(middlewares []PublisherDecorator, publisher Publisher) Publisher {
	for i := len(middlewares) - 1; i >= 0; i-- {
		publisher = middlewares[i].Decorate(publisher)
	}
	return publisher
}

type DiscardPublisher struct{}

func (pub DiscardPublisher) Topic() string                   { return "" }
func (pub DiscardPublisher) Publish(_ Message) (any, error)  { return nil, nil }
func (pub DiscardPublisher) Close(ctx context.Context) error { return nil }
