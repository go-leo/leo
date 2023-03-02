package stream

import "context"

type Bridge interface {
	Input() Subscriber
	Output() Publisher
	MessageTransformer() MessageTransformer
	Process(ctx context.Context)
}

type MessageTransformer interface {
	Transform(msg Message) (Message, error)
}

type MessageTransformerFunc func(msg Message) (Message, error)

func (m MessageTransformerFunc) Transform(msg Message) (Message, error) {
	return m(msg)
}
