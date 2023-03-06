package stream

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-leo/gox/syncx/brave"
	"github.com/go-leo/gox/syncx/chanx"
)

type Router interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	AppendBridge(bridges ...Bridge)
	AppendSubscriberDecorator(decorators ...SubscriberDecorator)
	AppendPublisherDecorator(decorators ...PublisherDecorator)
	AppendHandlerMiddleware(middlewares ...HandlerMiddleware)
	EventChan() <-chan Event
	Bridges() []Bridge
}

type router struct {
	bridges              []*bridge
	subscriberDecorators []SubscriberDecorator
	publisherDecorators  []PublisherDecorator
	middlewares          []HandlerMiddleware
	eventC               chan Event
	bridgeEventCs        []<-chan Event
}

func (r *router) Start(ctx context.Context) error {
	for _, b := range r.bridges {
		r.appendEventChans(b)
		r.process(ctx, b)
	}
	r.pipeEventChans()
	return nil
}

func (r *router) Stop(ctx context.Context) error {
	var errCs []<-chan error
	for _, bridge := range r.bridges {
		errCs = append(errCs, r.asyncClose(ctx, bridge))
	}
	errC := chanx.Combine(errCs...)
	errs := chanx.ReceiveUtilClosed(errC)
	return errors.Join(errs...)
}

func (r *router) AppendBridge(bridges ...Bridge) {
	for _, b := range bridges {
		r.bridges = append(r.bridges, b.(*bridge))
	}
}

func (r *router) AppendSubscriberDecorator(decorators ...SubscriberDecorator) {
	r.subscriberDecorators = append(r.subscriberDecorators, decorators...)
}

func (r *router) AppendPublisherDecorator(decorators ...PublisherDecorator) {
	r.publisherDecorators = append(r.publisherDecorators, decorators...)
}

func (r *router) AppendHandlerMiddleware(middlewares ...HandlerMiddleware) {
	r.middlewares = append(r.middlewares, middlewares...)
}

func (r *router) EventChan() <-chan Event {
	return r.eventC
}

func (r *router) Bridges() []Bridge {
	var bridges []Bridge
	for _, b := range r.bridges {
		bridges = append(bridges, b)
	}
	return bridges
}

func (r *router) process(ctx context.Context, b *bridge) {
	r.insertSubscriberDecorator(b)
	r.insertPublisherDecorators(b)
	r.insertHandlerMiddlewares(b)
	go func() {
		b.process(ctx)
	}()
}

func (r *router) appendEventChans(b *bridge) {
	r.bridgeEventCs = append(r.bridgeEventCs, b.eventChan())
}

func (r *router) pipeEventChans() {
	chanx.Pipe(chanx.Combine(r.bridgeEventCs...), r.eventC)
}

func (r *router) insertSubscriberDecorator(bridge *bridge) {
	for _, subscriberDecorator := range r.subscriberDecorators {
		bridge.unshiftSubscriberDecorator(subscriberDecorator)
	}
}

func (r *router) insertPublisherDecorators(bridge *bridge) {
	for _, publisherDecorator := range r.publisherDecorators {
		bridge.unshiftPublisherDecorator(publisherDecorator)
	}
}

func (r *router) insertHandlerMiddlewares(bridge *bridge) {
	for _, middleware := range r.middlewares {
		bridge.unshiftHandlerMiddleware(middleware)
	}
}

func (r *router) asyncClose(ctx context.Context, bridge *bridge) <-chan error {
	return brave.GoE(
		func() error { return bridge.close(ctx) },
		func(p any) error { return fmt.Errorf("%s", p) },
	)
}

func NewRouter() Router {
	return &router{}
}
