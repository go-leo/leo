package stream

import (
	"context"
	"errors"

	"github.com/go-leo/gox/syncx/chanx"
)

type Bridge interface {
	Input() Subscriber
	Output() Publisher
	MessageTransformer() MessageTransformer
	Process(ctx context.Context)
	Close() error
}

type MessageTransformer interface {
	Transform(msg Message) (Message, error)
}

type MessageTransformerFunc func(msg Message) (Message, error)

func (m MessageTransformerFunc) Transform(msg Message) (Message, error) {
	return m(msg)
}

var _ Bridge = new(bridge)

type bridge struct {
	subscriber         Subscriber
	publisher          Publisher
	messageTransformer MessageTransformer
	errC               chan<- error
	ackResultC         chan<- any
	nackResultC        chan<- any
	publishResultC     chan<- any
}

func (b *bridge) Input() Subscriber {
	return b.subscriber
}

func (b *bridge) Output() Publisher {
	return b.publisher
}

func (b *bridge) MessageTransformer() MessageTransformer {
	return b.messageTransformer
}

func (b *bridge) Process(ctx context.Context) {
	msgC, errC := b.subscriber.Subscribe(ctx)
	chanx.AsyncPipe(errC, b.errC)
	for msg := range msgC {
		b.handleMessage(msg)
	}
	b.closeC()
}

func (b *bridge) Close() error {
	return errors.Join(b.subscriber.Close(), b.publisher.Close())
}

func (b *bridge) handleMessage(msg Message) {
	err := b.processMessage(msg)
	if err != nil {
		b.sendError(err)
		b.nack(msg)
		return
	}
	b.ack(msg)
}

func (b *bridge) processMessage(msg Message) error {
	transform, err := b.messageTransformer.Transform(msg)
	if err != nil {
		return err
	}
	publishRes, err := b.publisher.Publish(transform)
	if err != nil {
		return err
	}
	b.sendPublishResult(publishRes)
	return nil
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
	if b.ackResultC == nil {
		return
	}
	b.ackResultC <- ackRes
}

func (b *bridge) sentNack(nackRes any) {
	if b.nackResultC == nil {
		return
	}
	b.nackResultC <- nackRes
}

func (b *bridge) sendError(err error) {
	if b.errC == nil {
		return
	}
	b.errC <- err
}

func (b *bridge) sendPublishResult(publishRes any) {
	if b.publishResultC == nil {
		return
	}
	b.publishResultC <- publishRes
}

func (b *bridge) closeC() {
	if b.ackResultC != nil {
		close(b.ackResultC)
	}
	if b.nackResultC != nil {
		close(b.nackResultC)
	}
	if b.errC != nil {
		close(b.errC)
	}
	if b.publishResultC != nil {
		close(b.publishResultC)
	}
}

type BridgeBuilder struct {
	subscriber         Subscriber
	publisher          Publisher
	messageTransformer MessageTransformer
	errC               chan<- error
	ackResultC         chan<- any
	nackResultC        chan<- any
	publishResultC     chan<- any
}

func (builder *BridgeBuilder) SetSubscriber(subscriber Subscriber) *BridgeBuilder {
	builder.subscriber = subscriber
	return builder
}

func (builder *BridgeBuilder) SetPublisher(publisher Publisher) *BridgeBuilder {
	builder.publisher = publisher
	return builder
}

func (builder *BridgeBuilder) SetMessageTransformer(messageTransformer MessageTransformer) *BridgeBuilder {
	builder.messageTransformer = messageTransformer
	return builder
}

func (builder *BridgeBuilder) SetErrC(errC chan<- error) *BridgeBuilder {
	builder.errC = errC
	return builder
}

func (builder *BridgeBuilder) SetAckResultC(ackResultC chan<- any) *BridgeBuilder {
	builder.ackResultC = ackResultC
	return builder
}

func (builder *BridgeBuilder) SetNackResultC(nackResultC chan<- any) *BridgeBuilder {
	builder.nackResultC = nackResultC
	return builder
}

func (builder *BridgeBuilder) SetPublishResultC(publishResultC chan<- any) *BridgeBuilder {
	builder.publishResultC = publishResultC
	return builder
}

func (builder *BridgeBuilder) Build() Bridge {
	return &bridge{
		subscriber:         builder.subscriber,
		publisher:          builder.publisher,
		messageTransformer: builder.messageTransformer,
		errC:               builder.errC,
		ackResultC:         builder.ackResultC,
		nackResultC:        builder.nackResultC,
		publishResultC:     builder.publishResultC,
	}
}

func NewBridgeBuilder() *BridgeBuilder {
	return &BridgeBuilder{}
}
