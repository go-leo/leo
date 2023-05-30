package stream

import (
	"context"
	"errors"
)

var ErrPublisherClosed = errors.New("publisher is closed")

// Publisher is message queue publisher
type Publisher interface {
	Publish(ctx context.Context, topic string, msg ...*Message) (any, error)
	Close(ctx context.Context) error
}

type noopPublisher struct{}

func (pub *noopPublisher) Publish(ctx context.Context, topic string, msg ...*Message) (any, error) {
	return struct{}{}, nil
}

func (pub *noopPublisher) Close(ctx context.Context) error {
	return nil
}

type multiPublisher struct {
	publishers []Publisher
}

func (pub *multiPublisher) Publish(ctx context.Context, topic string, msg ...*Message) (any, error) {
	var allRes []any
	for _, w := range pub.publishers {
		res, err := w.Publish(ctx, topic, msg...)
		if err != nil {
			return nil, err
		}
		allRes = append(allRes, res)
	}
	return allRes, nil
}

func (pub *multiPublisher) Close(ctx context.Context) error {
	for _, w := range pub.publishers {
		err := w.Close(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func MultiPublisher(publishers ...Publisher) Publisher {
	allPublishers := make([]Publisher, 0, len(publishers))
	for _, w := range publishers {
		if mw, ok := w.(*multiPublisher); ok {
			allPublishers = append(allPublishers, mw.publishers...)
		} else {
			allPublishers = append(allPublishers, w)
		}
	}
	return &multiPublisher{publishers: allPublishers}
}
