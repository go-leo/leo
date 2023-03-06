package stream

import (
	"context"
	"errors"
)

type Bridge interface {
	Name() string
	Input() Subscriber
	Output() Publisher
}

type EventKind int

const (
	ErrEvent           = 1
	AckResultEvent     = 1
	NackResultEvent    = 1
	PublishResultEvent = 1
)

type Event struct {
	EventKind     EventKind
	Bridge        Bridge
	Err           error
	AckResult     any
	NackResult    any
	PublishResult any
}

var _ Bridge = new(bridge)

type bridge struct {
	name string

	chainedSubscriber Subscriber
	chainedPublisher  Publisher
	chainedHandler    handler
	subscriber        Subscriber
	publisher         Publisher
	handler           handler

	subscriberDecorators []SubscriberDecorator
	publisherDecorators  []PublisherDecorator
	middlewares          []HandlerMiddleware

	eventC chan Event
}

func (b *bridge) Name() string {
	return b.name
}

func (b *bridge) Input() Subscriber {
	return b.subscriber
}

func (b *bridge) Output() Publisher {
	return b.publisher
}

func (b *bridge) close(ctx context.Context) error {
	return errors.Join(b.subscriber.Close(ctx), b.publisher.Close(ctx))
}

func (b *bridge) process(ctx context.Context) {
	b.chainedSubscriber = chainSubscriber(b.subscriberDecorators, b.subscriber)
	b.chainedPublisher = chainPublisher(b.publisherDecorators, b.publisher)
	b.chainedHandler = chainHandler(b.middlewares, b.handler)
	msgC, errC := b.chainedSubscriber.Subscribe(ctx)
	b.pipeErr(errC)
	for msg := range msgC {
		b.handleMessage(msg)
	}
	b.closeC()
}

func (b *bridge) appendSubscriberDecorator(decorators ...SubscriberDecorator) {
	b.subscriberDecorators = append(b.subscriberDecorators, decorators...)
}

func (b *bridge) appendPublisherDecorator(decorators ...PublisherDecorator) {
	b.publisherDecorators = append(b.publisherDecorators, decorators...)
}

func (b *bridge) appendHandlerMiddleware(middlewares ...HandlerMiddleware) {
	b.middlewares = append(b.middlewares, middlewares...)
}

func (b *bridge) unshiftSubscriberDecorator(decorators SubscriberDecorator) {
	b.subscriberDecorators = append([]SubscriberDecorator{decorators}, b.subscriberDecorators...)
}

func (b *bridge) unshiftPublisherDecorator(decorators PublisherDecorator) {
	b.publisherDecorators = append([]PublisherDecorator{decorators}, b.publisherDecorators...)
}

func (b *bridge) unshiftHandlerMiddleware(middleware HandlerMiddleware) {
	b.middlewares = append([]HandlerMiddleware{middleware}, b.middlewares...)
}

func (b *bridge) eventChan() <-chan Event {
	return b.eventC
}

func (b *bridge) handleMessage(msg Message) {
	transformedMsg, err := b.chainedHandler.Handle(msg)
	if err != nil {
		b.sendError(err)
		b.nack(msg)
		return
	}
	if b.publisher == nil {
		b.ack(msg)
		return
	}
	publishRes, err := b.chainedPublisher.Publish(transformedMsg)
	if err != nil {
		b.sendError(err)
		b.nack(msg)
		return
	}
	b.sendPublishResult(publishRes)
	b.ack(msg)
}

func (b *bridge) ack(msg Message) {
	ackRes, err := msg.Ack()
	if err != nil {
		b.sendError(err)
		return
	}
	b.sendAck(ackRes)
}

func (b *bridge) nack(msg Message) {
	nackRes, err := msg.Nack()
	if err != nil {
		b.sendError(err)
		return
	}
	b.sentNack(nackRes)
}

func (b *bridge) sendAck(ackRes any) {
	if ackRes == nil {
		return
	}
	if b.eventC == nil {
		return
	}
	select {
	case b.eventC <- Event{Bridge: b, EventKind: AckResultEvent, AckResult: ackRes}:
	default:
	}
}

func (b *bridge) sentNack(nackRes any) {
	if nackRes == nil {
		return
	}
	if b.eventC == nil {
		return
	}
	select {
	case b.eventC <- Event{Bridge: b, EventKind: NackResultEvent, NackResult: nackRes}:
	default:
	}
}

func (b *bridge) sendError(err error) {
	if err == nil {
		return
	}
	if b.eventC == nil {
		return
	}
	select {
	case b.eventC <- Event{Bridge: b, EventKind: ErrEvent, Err: err}:
	default:
	}
}

func (b *bridge) sendPublishResult(publishRes any) {
	if publishRes == nil {
		return
	}
	if b.eventC == nil {
		return
	}
	select {
	case b.eventC <- Event{Bridge: b, EventKind: PublishResultEvent, PublishResult: publishRes}:
	default:
	}
}

func (b *bridge) closeC() {
	if b.eventC != nil {
		close(b.eventC)
	}
}

func (b *bridge) pipeErr(errC <-chan error) {
	go func() {
		for err := range errC {
			b.sendError(err)
		}
	}()
}

func NewBridge(
	name string,
	subscriber Subscriber,
	publisher Publisher,
	handler HandlerFunc,
) Bridge {
	return &bridge{
		name:       name,
		subscriber: subscriber,
		publisher:  publisher,
		handler:    handler,
		eventC:     make(chan Event),
	}
}

func NewNoPublisherBridge(
	name string,
	subscriber Subscriber,
	handler NoPublishHandlerFunc,
) Bridge {
	return &bridge{
		name:       name,
		subscriber: subscriber,
		handler:    handler,
		eventC:     make(chan Event),
	}
}
