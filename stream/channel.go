package stream

import (
	"context"
	"errors"
	"fmt"
	"os/signal"
	"time"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/slicex"
)

type Channel interface {
	Subscriber() Subscriber
	Publisher() Publisher
	handleFunc() func(ctx context.Context, msg *Message) ([]*Message, error)
}

var _ Channel = (*channel)(nil)

type channel struct {
	subscriber      Subscriber
	publisher       Publisher
	HandleFunc      func(ctx context.Context, msg *Message) ([]*Message, error)
	MsgC            chan *Message
	ErrC            chan error
	Interceptor     Interceptor
	ShutdownTimeout time.Duration
}

func (channel *channel) Subscriber() Subscriber {
	return channel.subscriber
}

func (channel *channel) Publisher() Publisher {
	return channel.publisher
}

func (channel *channel) handleFunc() func(ctx context.Context, msg *Message) ([]*Message, error) {
	return channel.HandleFunc
}

func (channel *channel) pipe(ctx context.Context) error {
	errC := make(chan error)
	go func() {
		err := channel.subscriber.Subscribe(ctx, channel.MsgC, channel.ErrC)
		if err != nil {
			errC <- err
		}
		close(errC)
	}()
	for {
		select {
		case err := <-errC:
			return err
		case msg := <-channel.MsgC:
			channel.handleMessage(ctx, msg)
		case <-ctx.Done():
			return channel.close()
		}
	}
}

func (channel *channel) handleMessage(ctx context.Context, msg *Message) {
	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()

	if channel.Interceptor != nil {
		if err := channel.Interceptor(ctx, msg, channel, invoke); err != nil {
			channel.ErrC <- err
			channel.nackMessage(ctx, msg)
			return
		}
		channel.ackMessage(ctx, msg)
		return
	}

	if err := invoke(ctx, msg, channel); err != nil {
		channel.ErrC <- err
		channel.nackMessage(ctx, msg)
		return
	}
	channel.ackMessage(ctx, msg)
}

func invoke(ctx context.Context, msg *Message, channel Channel) error {
	// handle message
	messages, err := channel.handleFunc()(ctx, msg)
	if err != nil {
		return err
	}
	// if publisher is nil, ack message
	if channel.Publisher() == nil {
		return nil
	}

	// if messages is empty, ack message
	if slicex.IsEmpty(messages) {
		return nil
	}

	// publish message
	if _, err := channel.Publisher().Publish(ctx, messages...); err != nil {
		return err
	}
	return nil
}

func (channel *channel) ackMessage(ctx context.Context, msg *Message) {
	_, err := msg.Ack(ctx)
	if err != nil {
		if errors.Is(err, ErrMessageNacked) || errors.Is(err, ErrMessageAcked) {
			return
		}
		channel.ErrC <- fmt.Errorf("failed to ack message: %w", err)
		return
	}
}

func (channel *channel) nackMessage(ctx context.Context, msg *Message) {
	_, err := msg.Nack(ctx)
	if err != nil {
		if errors.Is(err, ErrMessageNacked) || errors.Is(err, ErrMessageAcked) {
			return
		}
		channel.ErrC <- fmt.Errorf("failed to nack message: %w", err)
		return
	}
}

func (channel *channel) close() error {
	ctx, cancel := signal.NotifyContext(context.Background())
	defer cancel()
	shutdownTimeout := channel.ShutdownTimeout
	if shutdownTimeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, shutdownTimeout)
		defer cancel()
	}
	var errs []error
	if err := channel.subscriber.Close(ctx); err != nil {
		errs = append(errs, err)
	}
	if channel.publisher == nil {
		return errors.Join(errs...)
	}
	if err := channel.publisher.Close(ctx); err != nil {
		errs = append(errs, err)
	}
	return errors.Join(errs...)
}
