package http

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os/signal"
	"sync"
	"sync/atomic"

	"codeup.aliyun.com/qimao/leo/leo/stream"
)

var (
	Default405Body = []byte("405 method not allowed")
	Default504Body = []byte("405 request has been canceled")
)

var _ stream.Subscriber = new(Subscriber)

type Subscriber struct {
	o          *options
	method     string
	url        string
	topic      string
	subscribed atomic.Bool
	closed     atomic.Bool
	closeC     chan struct{}
	stopC      chan struct{}
	wg         sync.WaitGroup
}

func (sub *Subscriber) Topic() string {
	return sub.topic
}

func (sub *Subscriber) Queue() string {
	return "http"
}

func (sub *Subscriber) Subscribe(ctx context.Context, msgC chan<- *stream.Message, errC chan<- error) error {
	if sub.closed.Load() {
		return stream.ErrSubscriberClosed
	}
	if !sub.subscribed.CompareAndSwap(false, true) {
		return stream.ErrSubscriberAlreadySubscribed
	}
	if sub.o.ServeMux == nil && sub.o.HttpServer == nil {
		return ErrServerIsNil
	}

	if sub.o.ServeMux == nil {
		sub.o.ServeMux = http.NewServeMux()
	}
	sub.o.ServeMux.HandleFunc(sub.url, func(resp http.ResponseWriter, req *http.Request) {
		if req.Method != sub.method {
			resp.WriteHeader(http.StatusMethodNotAllowed)
			_, _ = resp.Write(Default405Body)
		}
		select {
		case <-ctx.Done():
			resp.WriteHeader(http.StatusGatewayTimeout)
			_, _ = resp.Write(Default504Body)
			return
		case <-sub.closeC:
			resp.WriteHeader(http.StatusGatewayTimeout)
			_, _ = resp.Write(Default504Body)
			return
		default:
			sub.handleMsg(ctx, resp, req, msgC, errC)
		}
	})

	serveErrC := make(chan error)
	if sub.o.HttpServer != nil {
		sub.o.HttpServer.Handler = sub.o.ServeMux
		lis, err := net.Listen("tcp", sub.o.HttpServer.Addr)
		if err != nil {
			return err
		}
		if sub.o.HttpServer.TLSConfig != nil {
			lis = tls.NewListener(lis, sub.o.HttpServer.TLSConfig)
		}
		go func() {
			err := sub.o.HttpServer.Serve(lis)
			if errors.Is(err, http.ErrServerClosed) {
				return
			}
			if err != nil {
				serveErrC <- err
			}
		}()
	}
	defer func() {
		defer close(sub.stopC)
		sub.wg.Wait()
		if sub.o.HttpServer != nil {
			ctx, _ := signal.NotifyContext(context.Background())
			if sub.o.ShutdownTimeout > 0 {
				ctx, _ = context.WithTimeout(ctx, sub.o.ShutdownTimeout)
			}
			err := sub.o.HttpServer.Shutdown(ctx)
			if err != nil {
				errC <- err
			}
		}
	}()
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-sub.closeC:
			return nil
		case serveErr := <-serveErrC:
			return serveErr
		}
	}
}

func (sub *Subscriber) Close(ctx context.Context) error {
	if !sub.closed.CompareAndSwap(false, true) {
		return nil
	}
	close(sub.closeC)
	select {
	case <-ctx.Done():
	case <-sub.stopC:
	}
	return nil
}

func (sub *Subscriber) handleMsg(ctx context.Context, resp http.ResponseWriter, req *http.Request, msgC chan<- *stream.Message, errC chan<- error) {
	sub.wg.Add(1)
	defer sub.wg.Done()
	msg, err := sub.o.Marshaller.Unmarshal(ctx, sub.topic, req)
	if err != nil {
		errC <- fmt.Errorf("failed to unmarshal kafka message: %w", err)
		return
	}
	if sub.o.OnMessageReceived != nil {
		msg = sub.o.OnMessageReceived(msg, req)
	}

	ackC := make(chan struct{})
	stream.NotifyAck(msg, ackC, func(ctx context.Context, msg *stream.Message) (any, error) {
		if sub.o.AckResponse == nil {
			return nil, nil
		}
		return sub.o.AckResponse(resp, req, msg)
	})

	nackC := make(chan struct{})
	stream.NotifyNack(msg, nackC, func(ctx context.Context, msg *stream.Message) (any, error) {
		if sub.o.NackResponse == nil {
			return nil, nil
		}
		return sub.o.NackResponse(resp, req, msg)
	})

	select {
	case <-ctx.Done():
		return
	case <-sub.closeC:
		return
	case msgC <- msg:
	}

	select {
	case <-ctx.Done():
		return
	case <-sub.closeC:
		return
	case <-ackC:
		return
	case <-nackC:
		if sub.o.NackHandler != nil {
			sub.o.NackHandler(msg)
		}
		return
	}
}

func NewSubscriber(topic string, method, url string, opts ...Option) (*Subscriber, error) {
	o := &options{}
	o.apply(opts...)
	o.init()
	return &Subscriber{
		o:          o,
		topic:      topic,
		method:     method,
		url:        url,
		subscribed: atomic.Bool{},
		closed:     atomic.Bool{},
		closeC:     make(chan struct{}),
		stopC:      make(chan struct{}),
		wg:         sync.WaitGroup{},
	}, nil
}
