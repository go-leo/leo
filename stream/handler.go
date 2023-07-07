package stream

import (
	"codeup.aliyun.com/qimao/leo/leo/internal/gox/slicex"
	"codeup.aliyun.com/qimao/leo/leo/log"
	"context"
	"errors"
	"fmt"
	"os/signal"
	"sync/atomic"
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
	subscriber      Subscriber
	handleFunc      func(ctx context.Context, msg *Message) ([]*Message, error)
	publisher       Publisher
	logger          log.Logger
	msgC            chan *Message
	errC            chan error
	shutdownTimeout time.Duration
	running         atomic.Bool
}

func (handler *handlerWrapper) handle(ctx context.Context) error {
	errC := make(chan error)
	go func() {
		err := handler.subscriber.Subscribe(ctx, handler.msgC, handler.errC)
		if err != nil {
			errC <- err
		}
		close(errC)
	}()
	for {
		select {
		case err := <-errC:
			return err
		case msg := <-handler.msgC:
			handler.handleMessage(ctx, msg)
		case <-ctx.Done():
			return handler.close()
		}
	}
}

func (handler *handlerWrapper) handleMessage(ctx context.Context, msg *Message) {
	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()
	// handle message
	messages, err := handler.handleFunc(ctx, msg)
	if err != nil {
		handler.errC <- fmt.Errorf("failed to handle message, %w", err)
		handler.nackMessage(ctx, msg)
		return
	}
	// if publisher is nil, ack message
	if handler.publisher == nil {
		handler.ackMessage(ctx, msg)
		return
	}

	// if messages is empty, ack message
	if slicex.IsEmpty(messages) {
		handler.ackMessage(ctx, msg)
		return
	}

	// publish message
	publish, err := handler.publisher.Publish(ctx, messages...)
	if err != nil {
		handler.errC <- fmt.Errorf("failed to publish message, %w", err)
		handler.nackMessage(ctx, msg)
		return
	}
	handler.logger.DebugF(handler.msgIDField(msg), log.MsgField(fmt.Sprintf("successfully published message, %v", publish)))
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
		handler.errC <- fmt.Errorf("failed to ack message: %w", err)
		return
	}
	handler.logger.DebugF(handler.msgIDField(msg), log.MsgField(fmt.Sprintf("successfully acked message: %v", res)))
	return
}

func (handler *handlerWrapper) nackMessage(ctx context.Context, msg *Message) {
	res, err := msg.Nack(ctx)
	if err != nil {
		if errors.Is(err, ErrMessageNacked) || errors.Is(err, ErrMessageAcked) {
			return
		}
		handler.errC <- fmt.Errorf("failed to nack message: %w", err)
		return
	}
	handler.logger.DebugF(handler.msgIDField(msg), log.MsgField(fmt.Sprintf("successfully nacked message: %v", res)))
}

func (handler *handlerWrapper) close() error {
	ctx, cancel := signal.NotifyContext(context.Background())
	defer cancel()
	shutdownTimeout := handler.shutdownTimeout
	if shutdownTimeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, shutdownTimeout)
		defer cancel()
	}
	var errs []error
	if err := handler.subscriber.Close(ctx); err != nil {
		errs = append(errs, err)
	}
	if handler.publisher == nil {
		return errors.Join(errs...)
	}
	if err := handler.publisher.Close(ctx); err != nil {
		errs = append(errs, err)
	}
	return errors.Join(errs...)
}
