package stream

import (
	"context"
	"errors"
	"sync"
	"time"
)

var (
	ErrMessageAcked       = errors.New("message acked")
	ErrMessageNacked      = errors.New("message nacked")
	ErrAckNotImplemented  = errors.New("message is not implement ack")
	ErrNackNotImplemented = errors.New("message is not implement nack")
)

// Message is a Message. Message is emitted by Publisher and received by Subscriber.
type Message struct {
	ID      string
	Time    time.Time
	Payload []byte
	Header  Header

	m sync.Mutex

	ackC    chan struct{}
	ackFunc func(ctx context.Context, msg *Message) (any, error)
	acked   bool

	nackC    chan struct{}
	nackFunc func(ctx context.Context, msg *Message) (any, error)
	nacked   bool
}

func (m *Message) Ack(ctx context.Context) (any, error) {
	m.m.Lock()
	defer m.m.Unlock()
	return m.ack(ctx)
}

func (m *Message) Nack(ctx context.Context) (any, error) {
	m.m.Lock()
	defer m.m.Unlock()
	return m.nack(ctx)
}

func (m *Message) ack(ctx context.Context) (any, error) {
	if m.acked {
		return nil, ErrMessageAcked
	}
	if m.nacked {
		return nil, ErrMessageNacked
	}
	if m.ackFunc == nil {
		return nil, ErrAckNotImplemented
	}
	m.acked = true
	ackResult, err := m.ackFunc(ctx, m)
	m.ackC <- struct{}{}
	return ackResult, err
}

func (m *Message) nack(ctx context.Context) (any, error) {
	if m.nacked {
		return nil, ErrMessageNacked
	}
	if m.acked {
		return nil, ErrMessageAcked
	}
	if m.nackFunc == nil {
		return nil, ErrNackNotImplemented
	}
	m.nacked = true
	nackResult, err := m.nackFunc(ctx, m)
	m.nackC <- struct{}{}
	return nackResult, err
}

func NotifyAck(m *Message, ackC chan struct{}, ackFunc func(ctx context.Context, msg *Message) (any, error)) {
	m.m.Lock()
	defer m.m.Unlock()
	m.ackC = ackC
	m.ackFunc = ackFunc
}

func NotifyNack(m *Message, nackC chan struct{}, nackFunc func(ctx context.Context, msg *Message) (any, error)) {
	m.m.Lock()
	defer m.m.Unlock()
	m.nackC = nackC
	m.nackFunc = nackFunc
}
