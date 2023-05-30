package stream

import (
	"context"
	"errors"
	"sync"
	"time"
)

// Message is a Message. Message is emitted by Publisher and received by Subscriber.
type Message struct {
	ID      string
	Time    time.Time
	Payload []byte
	Header  Header

	sync.Mutex

	ackC    chan any
	ackFunc func(ctx context.Context, msg *Message) error
	acked   bool

	nackC    chan any
	nackFunc func(ctx context.Context, msg *Message) error
	nacked   bool
}

func (m *Message) Ack(ctx context.Context) error {
	m.Lock()
	defer m.Unlock()
	if m.acked {
		return nil
	}
	if m.nacked {
		return errors.New("message nacked")
	}
	if m.ackFunc == nil {
		return errors.New("message is not implement Ack")
	}
	m.acked = true
	err := m.ackFunc(ctx, m)
	m.ackC <- struct{}{}
	return err
}

func (m *Message) Nack(ctx context.Context) error {
	m.Lock()
	defer m.Unlock()
	if m.nacked {
		return nil
	}
	if m.acked {
		return errors.New("message acked")
	}
	if m.nackFunc == nil {
		return errors.New("message is not implement Nack")
	}
	m.nacked = true
	err := m.ackFunc(ctx, m)
	m.nackC <- struct{}{}
	return err
}

func (m *Message) NotifyAck(ackC chan any, ackFunc func(ctx context.Context, msg *Message) error) {
	m.Lock()
	defer m.Unlock()
	m.ackC = ackC
	m.ackFunc = ackFunc
}

func (m *Message) NotifyNack(nackC chan any, nackFunc func(ctx context.Context, msg *Message) error) {
	m.Lock()
	defer m.Unlock()
	m.nackC = nackC
	m.nackFunc = nackFunc
}
