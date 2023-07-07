package stream

import (
	"codeup.aliyun.com/qimao/leo/leo/actuator"
	"codeup.aliyun.com/qimao/leo/leo/actuator/health"
	"codeup.aliyun.com/qimao/leo/leo/log"
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"sync/atomic"
	"time"
)

type options struct {
	Handlers          []Handler
	PubSubHandlers    []PubSubHandler
	MessageBufferSize int
	ErrorHandler      func(err error)
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
	if o.ErrorHandler == nil {
		o.ErrorHandler = func(err error) {}
	}
	if o.Logger == nil {
		o.Logger = log.L()
	}
}

func Handlers(h ...Handler) Option {
	return func(o *options) {
		o.Handlers = append(o.Handlers, h...)
	}
}

func PubSubHandlers(h ...PubSubHandler) Option {
	return func(o *options) {
		o.PubSubHandlers = append(o.PubSubHandlers, h...)
	}
}

func MessageBufferSize(size int) Option {
	return func(o *options) {
		o.MessageBufferSize = size
	}
}

func ErrorHandler(f func(err error)) Option {
	return func(o *options) {
		o.ErrorHandler = f
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
	options         *options
	handlerWrappers []*handlerWrapper
	run             atomic.Bool
	alive           atomic.Bool
}

func (s *Streamer) Run(ctx context.Context) error {
	// check streamer is run
	if !s.run.CompareAndSwap(false, true) {
		return errors.New("streamer was run")
	}

	eg, ctx := errgroup.WithContext(ctx)

	// alive flag
	s.alive.Store(true)
	eg.Go(func() error {
		<-ctx.Done()
		s.alive.Store(false)
		return nil
	})

	// wrap Handler and PubSubHandler
	err := s.addHandles()
	if err != nil {
		return err
	}

	// async run all handlers to subscribe
	for _, handler := range s.handlerWrappers {
		handler := handler
		eg.Go(func() error { return handler.handle(ctx) })
		eg.Go(func() error {
			for {
				select {
				case <-ctx.Done():
					return nil
				case err := <-handler.errC:
					s.options.ErrorHandler(err)
				}
			}
		})
	}
	return eg.Wait()
}

func (s *Streamer) ActuatorHandler() actuator.Handler {
	return &actuatorHandler{streamer: s}
}

func (s *Streamer) HealthChecker() health.Checker {
	return &healthChecker{streamer: s}
}

func (s *Streamer) addHandles() error {
	for _, handler := range s.options.Handlers {
		subscriber, err := handler.Subscriber()
		if err != nil {
			return err
		}
		msgC := make(chan *Message, s.options.MessageBufferSize)
		errC := make(chan error, s.options.MessageBufferSize)
		s.handlerWrappers = append(s.handlerWrappers, &handlerWrapper{
			subscriber: subscriber,
			handleFunc: func(ctx context.Context, msg *Message) ([]*Message, error) {
				return nil, handler.Handle(ctx, msg)
			},
			publisher:       nil,
			msgC:            msgC,
			errC:            errC,
			logger:          s.options.Logger,
			shutdownTimeout: s.options.ShutdownTimeout,
			running:         atomic.Bool{},
		})
	}
	for _, handler := range s.options.PubSubHandlers {
		subscriber, err := handler.Subscriber()
		if err != nil {
			return err
		}
		publisher, err := handler.Publisher()
		if err != nil {
			return err
		}
		msgC := make(chan *Message, s.options.MessageBufferSize)
		errC := make(chan error, s.options.MessageBufferSize)
		s.handlerWrappers = append(s.handlerWrappers, &handlerWrapper{
			subscriber:      subscriber,
			handleFunc:      handler.Handle,
			publisher:       publisher,
			msgC:            msgC,
			errC:            errC,
			logger:          s.options.Logger,
			shutdownTimeout: s.options.ShutdownTimeout,
			running:         atomic.Bool{},
		})
	}
	return nil
}

func (s *Streamer) isAlive() bool {
	return s.alive.Load()
}

func NewStreamer(opts ...Option) *Streamer {
	o := new(options)
	o.apply(opts...)
	o.init()
	return &Streamer{
		options:         o,
		handlerWrappers: nil,
		run:             atomic.Bool{},
		alive:           atomic.Bool{},
	}
}
