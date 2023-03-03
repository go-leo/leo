package stream

import "context"

// Message is a message. Message is emitted by Publisher and received by Subscriber.
type Message interface {
	// ID return message uniq id
	ID() string
	// Payload return message payload
	Payload() []byte
	// Header return message header
	Header() Header
	// Context return message context
	Context() context.Context
	// WithContext return new message with context
	WithContext(ctx context.Context) Message
	// Ack ack message
	Ack() (any, error)
	// Nack nack message
	Nack() (any, error)
	// Acked returns a channel which is closed when ack message. function will be invoked at Ack
	Acked(f func() (any, error)) <-chan struct{}
	// Nacked returns a channel which is closed when nack message. function will be invoked at Nack
	Nacked(f func() (any, error)) <-chan struct{}
}

var _ Message = new(message)

type message struct {
	id      string
	payload []byte
	header  Header
	ctx     context.Context
	ackC    chan struct{}
	nackC   chan struct{}
	onAck   func() (any, error)
	onNack  func() (any, error)
}

func (m *message) ID() string {
	return m.id
}

func (m *message) Payload() []byte {
	return m.payload
}

func (m *message) Header() Header {
	if m.header == nil {
		m.header = make(Header)
	}
	return m.header
}

func (m *message) Context() context.Context {
	return m.ctx
}

func (m *message) WithContext(ctx context.Context) Message {
	cloned := *m
	cloned.ctx = ctx
	return &cloned
}

func (m *message) Ack() (any, error) {
	close(m.ackC)
	if m.onAck != nil {
		return m.onAck()
	}
	return nil, nil
}

func (m *message) Nack() (any, error) {
	close(m.nackC)
	if m.onNack != nil {
		return m.onNack()
	}
	return nil, nil
}

func (m *message) Acked(f func() (any, error)) <-chan struct{} {
	m.onAck = f
	return m.ackC
}

func (m *message) Nacked(f func() (any, error)) <-chan struct{} {
	m.onNack = f
	return m.nackC
}

type MessageBuilder struct {
	id      string
	payload []byte
	header  Header
	ctx     context.Context
}

func (builder *MessageBuilder) SetId(id string) *MessageBuilder {
	builder.id = id
	return builder
}

func (builder *MessageBuilder) SetPayload(payload []byte) *MessageBuilder {
	builder.payload = payload
	return builder
}

func (builder *MessageBuilder) SetHeader(header Header) *MessageBuilder {
	builder.header = header
	return builder
}

func (builder *MessageBuilder) Build(ctx context.Context) Message {
	return &message{
		id:      builder.id,
		payload: builder.payload,
		header:  builder.header,
		ctx:     ctx,
		ackC:    make(chan struct{}),
		nackC:   make(chan struct{}),
		onAck:   func() (any, error) { return nil, nil },
		onNack:  func() (any, error) { return nil, nil },
	}
}

func NewMessageBuilder() *MessageBuilder {
	return &MessageBuilder{}
}
