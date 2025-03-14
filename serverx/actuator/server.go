package actuator

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	kitlog "github.com/go-kit/log"
	"github.com/go-leo/gox/contextx"
	"github.com/go-leo/leo/v3/logx"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Server struct {
	port    int
	handler http.Handler
	o       *options
}

type options struct {
	DisableGeneralOptionsHandler bool
	TLSConfig                    *tls.Config
	ReadTimeout                  time.Duration
	ReadHeaderTimeout            time.Duration
	WriteTimeout                 time.Duration
	IdleTimeout                  time.Duration
	MaxHeaderBytes               int
	TLSNextProto                 map[string]func(*http.Server, *tls.Conn, http.Handler)
	ConnState                    func(net.Conn, http.ConnState)
	BaseContext                  func(net.Listener) context.Context
	ConnContext                  func(ctx context.Context, c net.Conn) context.Context

	ShutdownContext func(ctx context.Context) (context.Context, context.CancelCauseFunc)
	Logger          kitlog.Logger
}

type Option func(o *options)

func (o *options) apply(opts ...Option) *options {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func (o *options) complete() *options {
	if o.Logger == nil {
		o.Logger = logx.New(os.Stdout, logx.JSON(), logx.Timestamp(), logx.Caller(0), logx.Sync())
	}
	return o
}

func DisableGeneralOptionsHandlerLS() Option {
	return func(o *options) {
		o.DisableGeneralOptionsHandler = true
	}
}

func TLSConfig(conf *tls.Config) Option {
	return func(o *options) {
		o.TLSConfig = conf
	}
}

func ReadTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.ReadTimeout = timeout
	}
}

func ReadHeaderTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.ReadHeaderTimeout = timeout
	}
}

func WriteTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.WriteTimeout = timeout
	}
}

func IdleTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.IdleTimeout = timeout
	}
}

func MaxHeaderBytes(size int) Option {
	return func(o *options) {
		o.MaxHeaderBytes = size
	}
}

func TLSNextProto(m map[string]func(*http.Server, *tls.Conn, http.Handler)) Option {
	return func(o *options) {
		o.TLSNextProto = m
	}
}

func ConnState(f func(net.Conn, http.ConnState)) Option {
	return func(o *options) {
		o.ConnState = f
	}
}

func BaseContext(f func(net.Listener) context.Context) Option {
	return func(o *options) {
		o.BaseContext = f
	}
}

func ConnContext(f func(ctx context.Context, c net.Conn) context.Context) Option {
	return func(o *options) {
		o.ConnContext = f
	}
}

func ShutdownContext(f func(ctx context.Context) (context.Context, context.CancelCauseFunc)) Option {
	return func(o *options) {
		o.ShutdownContext = f
	}
}

func Logger(logger kitlog.Logger) Option {
	return func(o *options) {
		o.Logger = logger
	}
}

func NewServer(port int, handler http.Handler, opts ...Option) *Server {
	return &Server{
		port:    port,
		handler: handler,
		o:       new(options).apply(opts...).complete(),
	}
}

func (s *Server) Run(ctx context.Context) error {
	// listen port.
	lis, err := net.Listen("tcp", net.JoinHostPort("", strconv.Itoa(s.port)))
	if err != nil {
		return err
	}
	if s.o.TLSConfig != nil {
		lis = tls.NewListener(lis, s.o.TLSConfig)
	}

	if s.o.BaseContext == nil {
		s.o.BaseContext = func(listener net.Listener) context.Context { return ctx }
	}
	if s.o.ConnContext == nil {
		s.o.ConnContext = func(ctx context.Context, c net.Conn) context.Context { return ctx }
	}
	httpSrv := &http.Server{
		Addr:                         "",
		Handler:                      s.handler,
		DisableGeneralOptionsHandler: s.o.DisableGeneralOptionsHandler,
		ReadTimeout:                  s.o.ReadTimeout,
		ReadHeaderTimeout:            s.o.ReadHeaderTimeout,
		WriteTimeout:                 s.o.WriteTimeout,
		IdleTimeout:                  s.o.IdleTimeout,
		MaxHeaderBytes:               s.o.MaxHeaderBytes,
		TLSNextProto:                 s.o.TLSNextProto,
		ConnState:                    s.o.ConnState,
		ErrorLog:                     log.New(kitlog.NewStdlibAdapter(s.o.Logger), "", 0),
		BaseContext:                  s.o.BaseContext,
		ConnContext:                  s.o.ConnContext,
	}

	// server serve.
	var serveErrC = make(chan error, 1)
	go func() {
		defer close(serveErrC)
		err := httpSrv.Serve(lis)
		if err == nil || errors.Is(err, http.ErrServerClosed) {
			return
		}
		serveErrC <- err
	}()

	select {
	case serveErr := <-serveErrC:
		return serveErr
	case <-ctx.Done():
		var errs = []error{
			fmt.Errorf("actuator server exit serve, %w", contextx.Error(ctx)),
		}
		// graceful shutdown, deregister and shutdown
		ctx = context.WithoutCancel(ctx)
		if s.o.ShutdownContext != nil {
			var cancelFunc context.CancelCauseFunc
			ctx, cancelFunc = s.o.ShutdownContext(ctx)
			defer cancelFunc(nil)
		}
		if err := httpSrv.Shutdown(ctx); err != nil {
			errs = append(errs, fmt.Errorf("actuator server shutdown, %w", err))
		}
		return errors.Join(errs...)
	}
}
