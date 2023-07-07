package stream

import (
	"codeup.aliyun.com/qimao/leo/leo/internal/gox/slicex"
	"codeup.aliyun.com/qimao/leo/leo/log"
	"context"
	"errors"
	"fmt"
	"os/signal"
	"time"
)

type PubSubHandler interface {
	Subscriber() (Subscriber, error)
	Handle(ctx context.Context, msg *Message) ([]*Message, error)
	Publisher() (Publisher, error)
}

type Handler interface {
	Subscriber() (Subscriber, error)
	Handle(ctx context.Context, msg *Message) error
}

type handlerWrapper struct {
	Subscriber      Subscriber
	HandleFunc      func(ctx context.Context, msg *Message) ([]*Message, error)
	Publisher       Publisher
	Logger          log.Logger
	MsgC            chan *Message
	ErrC            chan error
	Interceptors    []Interceptor
	ShutdownTimeout time.Duration
}

func (handler *handlerWrapper) handle(ctx context.Context) error {
	errC := make(chan error)
	go func() {
		err := handler.Subscriber.Subscribe(ctx, handler.MsgC, handler.ErrC)
		if err != nil {
			errC <- err
		}
		close(errC)
	}()
	for {
		select {
		case err := <-errC:
			return err
		case msg := <-handler.MsgC:
			handler.handleMessage(ctx, msg)
		case <-ctx.Done():
			return handler.close()
		}
	}
}

func (handler *handlerWrapper) handleMessage(ctx context.Context, msg *Message) {
	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()

	// intercept message
	ctx, msg, err := interceptMessage(ctx, msg, handler.Interceptors...)
	if err != nil {
		handler.ErrC <- fmt.Errorf("failed to handle message, %w", err)
		handler.nackMessage(ctx, msg)
		return
	}

	// handle message
	messages, err := handler.HandleFunc(ctx, msg)
	if err != nil {
		handler.ErrC <- fmt.Errorf("failed to handle message, %w", err)
		handler.nackMessage(ctx, msg)
		return
	}
	// if publisher is nil, ack message
	if handler.Publisher == nil {
		handler.ackMessage(ctx, msg)
		return
	}

	// if messages is empty, ack message
	if slicex.IsEmpty(messages) {
		handler.ackMessage(ctx, msg)
		return
	}

	// publish message
	publish, err := handler.Publisher.Publish(ctx, messages...)
	if err != nil {
		handler.ErrC <- fmt.Errorf("failed to publish message, %w", err)
		handler.nackMessage(ctx, msg)
		return
	}
	handler.Logger.DebugF(handler.msgIDField(msg), log.MsgField(fmt.Sprintf("successfully published message, %v", publish)))
	handler.ackMessage(ctx, msg)
	return
}

func (handler *handlerWrapper) msgIDField(msg *Message) log.F {
	return log.F{K: "msg_id", V: msg.ID}
}

func (handler *handlerWrapper) ackMessage(ctx context.Context, msg *Message) {
	res, err := msg.Ack(ctx)
	if err != nil {
		if errors.Is(err, ErrMessageNacked) || errors.Is(err, ErrMessageAcked) {
			return
		}
		handler.ErrC <- fmt.Errorf("failed to ack message: %w", err)
		return
	}
	handler.Logger.DebugF(handler.msgIDField(msg), log.MsgField(fmt.Sprintf("successfully acked message: %v", res)))
	return
}

func (handler *handlerWrapper) nackMessage(ctx context.Context, msg *Message) {
	res, err := msg.Nack(ctx)
	if err != nil {
		if errors.Is(err, ErrMessageNacked) || errors.Is(err, ErrMessageAcked) {
			return
		}
		handler.ErrC <- fmt.Errorf("failed to nack message: %w", err)
		return
	}
	handler.Logger.DebugF(handler.msgIDField(msg), log.MsgField(fmt.Sprintf("successfully nacked message: %v", res)))
}

func (handler *handlerWrapper) close() error {
	ctx, cancel := signal.NotifyContext(context.Background())
	defer cancel()
	shutdownTimeout := handler.ShutdownTimeout
	if shutdownTimeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, shutdownTimeout)
		defer cancel()
	}
	var errs []error
	if err := handler.Subscriber.Close(ctx); err != nil {
		errs = append(errs, err)
	}
	if handler.Publisher == nil {
		return errors.Join(errs...)
	}
	if err := handler.Publisher.Close(ctx); err != nil {
		errs = append(errs, err)
	}
	return errors.Join(errs...)
}
