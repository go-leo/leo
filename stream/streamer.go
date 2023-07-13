package stream

import (
	"codeup.aliyun.com/qimao/leo/leo/actuator"
	"codeup.aliyun.com/qimao/leo/leo/actuator/health"
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"sync/atomic"
)

type Streamer struct {
	options  *options
	channels []*channel
	run      atomic.Bool
	alive    atomic.Bool
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

	// add channel
	err := s.addChannels()
	if err != nil {
		return err
	}

	// async run all handlers to subscribe
	for _, channel := range s.channels {
		channel := channel
		eg.Go(func() error { return channel.pipe(ctx) })
		eg.Go(func() error {
			for {
				select {
				case <-ctx.Done():
					return nil
				case err := <-channel.ErrC:
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

func (s *Streamer) addChannels() error {
	for _, handler := range s.options.Handlers {
		subscriber, err := handler.Subscriber()
		if err != nil {
			return err
		}
		msgC := make(chan *Message, s.options.MessageBufferSize)
		errC := make(chan error, s.options.MessageBufferSize)
		s.channels = append(s.channels, &channel{
			subscriber:      subscriber,
			HandleFunc:      func(ctx context.Context, msg *Message) ([]*Message, error) { return nil, handler.Handle(ctx, msg) },
			publisher:       nil,
			MsgC:            msgC,
			ErrC:            errC,
			ShutdownTimeout: s.options.ShutdownTimeout,
			Interceptor:     chainInterceptors(s.options.Interceptors),
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
		s.channels = append(s.channels, &channel{
			subscriber:      subscriber,
			HandleFunc:      handler.Handle,
			publisher:       publisher,
			MsgC:            msgC,
			ErrC:            errC,
			ShutdownTimeout: s.options.ShutdownTimeout,
			Interceptor:     chainInterceptors(s.options.Interceptors),
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
		options:  o,
		channels: nil,
		run:      atomic.Bool{},
		alive:    atomic.Bool{},
	}
}
