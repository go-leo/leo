package message

import "context"

// Publisher is message queue publisher
type Publisher interface {
	Publish(ctx context.Context, topic string, msg ...*Message) (any, error)
	Close(ctx context.Context) error
}

type PublisherDecorator interface {
	Decorate(pub Publisher) Publisher
}

type PublisherDecoratorFunc func(pub Publisher) Publisher

func (f PublisherDecoratorFunc) Decorate(pub Publisher) Publisher {
	return f(pub)
}

func ChainPublishers(pub Publisher, mwf ...PublisherDecoratorFunc) Publisher {
	for i := len(mwf) - 1; i >= 0; i-- {
		pub = mwf[i].Decorate(pub)
	}
	return pub
}
