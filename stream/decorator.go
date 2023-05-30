package stream

type SubscriberDecorator interface {
	Decorate(sub Subscriber) Subscriber
}

type SubscriberDecoratorFunc func(sub Subscriber) Subscriber

func (f SubscriberDecoratorFunc) Decorate(sub Subscriber) Subscriber {
	return f(sub)
}

func DecorateSubscribers(sub Subscriber, mwf ...SubscriberDecoratorFunc) Subscriber {
	for i := len(mwf) - 1; i >= 0; i-- {
		sub = mwf[i].Decorate(sub)
	}
	return sub
}

type PublisherDecorator interface {
	Decorate(pub Publisher) Publisher
}

type PublisherDecoratorFunc func(pub Publisher) Publisher

func (f PublisherDecoratorFunc) Decorate(pub Publisher) Publisher {
	return f(pub)
}

func DecoratePublishers(pub Publisher, mwf ...PublisherDecoratorFunc) Publisher {
	for i := len(mwf) - 1; i >= 0; i-- {
		pub = mwf[i].Decorate(pub)
	}
	return pub
}
