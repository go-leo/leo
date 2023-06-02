package stream

import (
	"codeup.aliyun.com/qimao/leo/leo/actuator"
	"codeup.aliyun.com/qimao/leo/leo/actuator/health"
	"codeup.aliyun.com/qimao/leo/leo/log"
	"context"
	"errors"
	"fmt"
	"github.com/go-leo/gox/slicex"
	"os/signal"
	"sync/atomic"
	"time"

	"golang.org/x/sync/errgroup"
)

type options struct {
	Handlers          []*handler
	MessageBufferSize int
	ErrorC            chan error
	Logger            log.Logger
	ShutdownTimeout   time.Duration
}

type Option func(o *options)

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *options) init() {
	if o.MessageBufferSize <= 0 {
		o.MessageBufferSize = 1
	}
	if o.Logger == nil {
		o.Logger = log.L()
	}
}

func Handler(
	subscriber Subscriber,
	handleFunc func(ctx context.Context, msg *Message) ([]*Message, error),
	publisher Publisher,
) Option {
	return func(o *options) {
		o.Handlers = append(o.Handlers, &handler{
			Subscriber: subscriber,
			HandleFunc: handleFunc,
			Publisher:  publisher,
		})
	}
}

func NoPublishHandler(
	subscriber Subscriber,
	handleFunc func(ctx context.Context, msg *Message) error,
) Option {
	return func(o *options) {
		o.Handlers = append(o.Handlers, &handler{
			Subscriber:          subscriber,
			NoPublishHandleFunc: handleFunc,
		})
	}
}

func MessageBufferSize(size int) Option {
	return func(o *options) {
		o.MessageBufferSize = size
	}
}

func ErrorC(errC chan error) Option {
	return func(o *options) {
		o.ErrorC = errC
	}
}

func Logger(l log.Logger) Option {
	return func(o *options) {
		o.Logger = l
	}
}

func ShutdownTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.ShutdownTimeout = timeout
	}
}

type Streamer struct {
	options *options
	run     atomic.Bool
	alive   atomic.Bool
}

func (s *Streamer) Run(ctx context.Context) error {
	// check streamer is run
	if !s.run.CompareAndSwap(false, true) {
		return errors.New("scheduler is run")
	}
	s.alive.Store(true)
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		<-ctx.Done()
		s.alive.Store(false)
		return nil
	})
	// async run all handlers to subscribe
	for _, handler := range s.options.Handlers {
		handler := handler
		if handler.Subscriber == nil {
			return errors.New("subscriber is nil")
		}
		if handler.HandleFunc == nil && handler.NoPublishHandleFunc == nil {
			return errors.New("handle func is nil")
		}
		handler.Streamer = s
		handler.MessageC = make(chan *Message, s.options.MessageBufferSize)
		handler.ErrorC = s.options.ErrorC
		if handler.NoPublishHandleFunc != nil && handler.HandleFunc == nil {
			handler.HandleFunc = func(ctx context.Context, msg *Message) ([]*Message, error) {
				return nil, handler.NoPublishHandleFunc(ctx, msg)
			}
		}
		handler.Logger = s.options.Logger
		eg.Go(func() error { return handler.Handle(ctx) })
	}
	err := eg.Wait()
	return err
}

func (s *Streamer) ActuatorHandler() actuator.Handler {
	return &actuatorHandler{streamer: s}
}

func (s *Streamer) HealthChecker() health.Checker {
	return &healthChecker{streamer: s}
}

func (s *Streamer) isAlive() bool {
	return s.alive.Load()
}

type handler struct {
	Streamer            *Streamer
	Subscriber          Subscriber
	MessageC            chan *Message
	ErrorC              chan error
	HandleFunc          func(ctx context.Context, msg *Message) ([]*Message, error)
	Publisher           Publisher
	NoPublishHandleFunc func(ctx context.Context, msg *Message) error
	Logger              log.Logger
}

func (h *handler) Handle(ctx context.Context) error {
	if h.Subscriber == nil {
		return errors.New("subscriber is nil")
	}
	err := h.Subscriber.Subscribe(ctx, h.MessageC, h.ErrorC)
	if err != nil {
		return err
	}
	for {
		select {
		case msg := <-h.MessageC:
			h.handleMessage(ctx, msg)
		case <-ctx.Done():
			return h.close()
		}
	}
}

func (h *handler) handleMessage(ctx context.Context, msg *Message) {
	// handle message
	messages, err := h.HandleFunc(ctx, msg)
	if err != nil {
		h.Logger.ErrorF(h.msgIDField(msg), log.ErrField(fmt.Errorf("failed to handle message, %w", err)))
		h.nackMessage(ctx, msg)
		return
	}
	// if messages is empty, ack message
	if slicex.IsEmpty(messages) {
		h.ackMessage(ctx, msg)
		return
	}
	// if publisher is nil, ack message
	if h.Publisher == nil {
		h.ackMessage(ctx, msg)
		return
	}
	// publish message
	publish, err := h.Publisher.Publish(ctx, messages...)
	if err != nil {
		h.Logger.ErrorF(h.msgIDField(msg), log.ErrField(fmt.Errorf("failed to publish message, %w", err)))
		h.nackMessage(ctx, msg)
		return
	}
	h.Logger.DebugF(h.msgIDField(msg), log.MsgField(fmt.Sprintf("successfully published message, %v", publish)))
	h.ackMessage(ctx, msg)
	return
}

func (h *handler) msgIDField(msg *Message) log.F {
	return log.F{K: "msg_id", V: msg.ID}
}

func (h *handler) ackMessage(ctx context.Context, msg *Message) {
	res, err := msg.Ack(ctx)
	if err != nil {
		h.Logger.ErrorF(h.msgIDField(msg), log.ErrField(fmt.Errorf("failed to ack message: %w", err)))
		return
	}
	h.Logger.DebugF(h.msgIDField(msg), log.MsgField(fmt.Sprintf("successfully acked message: %v", res)))
	return
}

func (h *handler) nackMessage(ctx context.Context, msg *Message) {
	res, err := msg.Nack(ctx)
	if err != nil {
		h.Logger.ErrorF(h.msgIDField(msg), log.ErrField(fmt.Errorf("failed to nack message: %w", err)))
		return
	}
	h.Logger.DebugF(h.msgIDField(msg), log.MsgField(fmt.Sprintf("successfully nacked message: %v", res)))
}

func (h *handler) close() error {
	ctx, cancel := signal.NotifyContext(context.Background())
	defer cancel()
	shutdownTimeout := h.Streamer.options.ShutdownTimeout
	if shutdownTimeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, shutdownTimeout)
		defer cancel()
	}
	var errs []error
	if err := h.Subscriber.Close(ctx); err != nil {
		errs = append(errs, err)
	}
	if h.Publisher == nil {
		return errors.Join(errs...)
	}
	if err := h.Publisher.Close(ctx); err != nil {
		errs = append(errs, err)
	}
	return errors.Join(errs...)
}
