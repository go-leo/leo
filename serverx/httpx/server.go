package httpx

import (
	"context"
	"crypto/tls"
	"errors"
	kitlog "github.com/go-kit/log"
	"github.com/go-leo/leo/v3/healthx"
	"github.com/go-leo/leo/v3/logx"
	"github.com/go-leo/leo/v3/sdx"
	internalsd "github.com/go-leo/leo/v3/serverx/internal/sd"
	"log"
	"net"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

type Server struct {
	o       *options
	handler http.Handler
}

type options struct {
	Port int

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

	ShutdownTimeout *time.Duration

	Builder  sdx.Builder
	Instance string
	Color    string
}

type Option func(o *options)

func (o *options) apply(opts ...Option) *options {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func (o *options) complete() *options {
	return o
}

func Port(p int) Option {
	return func(o *options) {
		o.Port = p
	}
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

func ShutdownTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.ShutdownTimeout = &timeout
	}
}

func RegistrarBuilder(builder sdx.Builder) Option {
	return func(o *options) {
		o.Builder = builder
	}
}

func Instance(instance string) Option {
	return func(o *options) {
		o.Instance = instance
	}
}

func Color(color string) Option {
	return func(o *options) {
		o.Color = color
	}
}

func NewServer(handler http.Handler, opts ...Option) *Server {
	return &Server{
		handler: handler,
		o:       new(options).apply(opts...).complete(),
	}
}

func (s *Server) Run(ctx context.Context) error {
	// listen port.
	lis, err := net.Listen("tcp", net.JoinHostPort("", strconv.Itoa(s.o.Port)))
	if err != nil {
		return err
	}
	if s.o.TLSConfig != nil {
		lis = tls.NewListener(lis, s.o.TLSConfig)
	}

	// create registrar.
	registrar, err := internalsd.NewRegistrar(ctx, lis, s.o.Builder, s.o.Instance, s.o.Color)
	if err != nil {
		return err
	}

	httpSrv := s.newHttpServer(ctx)

	// register health check service.
	checker := healthx.NewChecker("http")
	healthx.RegisterChecker(checker)

	// server serve.
	serveErrC := s.serve(httpSrv, lis, checker)

	// register
	internalsd.Register(registrar)

	select {
	case serveErr := <-serveErrC:
		if serveErr == nil {
			return nil
		}
		// if server failed to start, deregister
		internalsd.Deregister(registrar)
		return serveErr
	case <-ctx.Done():
		// graceful shutdown, deregister and shutdown
		internalsd.Deregister(registrar)
		err := s.shutdown(ctx, httpSrv, checker)
		return errors.Join(ctx.Err(), err)
	}
}

func (s *Server) serve(httpSrv *http.Server, lis net.Listener, checker healthx.Checker) chan error {
	var errC = make(chan error, 1)
	go func() {
		defer close(errC)
		checker.Resume()
		err := httpSrv.Serve(lis)
		checker.Shutdown()
		if err == nil || errors.Is(err, http.ErrServerClosed) {
			return
		}
		errC <- err
	}()
	// make sure server serve finish
	runtime.Gosched()
	return errC
}

func (s *Server) shutdown(ctx context.Context, httpSrv *http.Server, checker healthx.Checker) error {
	ctx = context.WithoutCancel(ctx)
	if s.o.ShutdownTimeout != nil {
		var cancelFunc context.CancelFunc
		ctx, cancelFunc = context.WithTimeout(ctx, *s.o.ShutdownTimeout)
		defer cancelFunc()
	}
	return httpSrv.Shutdown(ctx)
}

func (s *Server) newHttpServer(ctx context.Context) *http.Server {
	if s.o.BaseContext == nil {
		s.o.BaseContext = func(listener net.Listener) context.Context { return ctx }
	}
	if s.o.ConnContext == nil {
		s.o.ConnContext = func(ctx context.Context, c net.Conn) context.Context { return ctx }
	}
	return &http.Server{
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
		ErrorLog:                     log.New(kitlog.NewStdlibAdapter(logx.L()), "", 0),
		BaseContext:                  s.o.BaseContext,
		ConnContext:                  s.o.ConnContext,
	}
}
