package httpx

import (
	"context"
	"crypto/tls"
	"errors"
	"github.com/go-kit/kit/sd"
	kitlog "github.com/go-kit/log"
	"github.com/go-leo/leo/v3/logx"
	"github.com/go-leo/leo/v3/sdx"
	"github.com/go-leo/leo/v3/transportx/internal"
	"log"
	"net"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"time"
)

type Server struct {
	o       *serverOptions
	handler http.Handler
}

type serverOptions struct {
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

	RegistrarBuilder sdx.Builder
	Instance         string
	Color            string
}

type ServerOption func(o *serverOptions)

func (o *serverOptions) apply(opts ...ServerOption) *serverOptions {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func (o *serverOptions) init() *serverOptions {
	return o
}

func Port(p int) ServerOption {
	return func(o *serverOptions) {
		o.Port = p
	}
}

func DisableGeneralOptionsHandlerLS() ServerOption {
	return func(o *serverOptions) {
		o.DisableGeneralOptionsHandler = true
	}
}

func TLSConfig(conf *tls.Config) ServerOption {
	return func(o *serverOptions) {
		o.TLSConfig = conf
	}
}

func ReadTimeout(timeout time.Duration) ServerOption {
	return func(o *serverOptions) {
		o.ReadTimeout = timeout
	}
}

func ReadHeaderTimeout(timeout time.Duration) ServerOption {
	return func(o *serverOptions) {
		o.ReadHeaderTimeout = timeout
	}
}

func WriteTimeout(timeout time.Duration) ServerOption {
	return func(o *serverOptions) {
		o.WriteTimeout = timeout
	}
}

func IdleTimeout(timeout time.Duration) ServerOption {
	return func(o *serverOptions) {
		o.IdleTimeout = timeout
	}
}

func MaxHeaderBytes(size int) ServerOption {
	return func(o *serverOptions) {
		o.MaxHeaderBytes = size
	}
}

func TLSNextProto(m map[string]func(*http.Server, *tls.Conn, http.Handler)) ServerOption {
	return func(o *serverOptions) {
		o.TLSNextProto = m
	}
}

func ConnState(f func(net.Conn, http.ConnState)) ServerOption {
	return func(o *serverOptions) {
		o.ConnState = f
	}
}

func BaseContext(f func(net.Listener) context.Context) ServerOption {
	return func(o *serverOptions) {
		o.BaseContext = f
	}
}

func ConnContext(f func(ctx context.Context, c net.Conn) context.Context) ServerOption {
	return func(o *serverOptions) {
		o.ConnContext = f
	}
}

func ShutdownTimeout(timeout time.Duration) ServerOption {
	return func(o *serverOptions) {
		o.ShutdownTimeout = &timeout
	}
}

func RegistrarBuilder(builder sdx.Builder) ServerOption {
	return func(o *serverOptions) {
		o.RegistrarBuilder = builder
	}
}

func Instance(instance string) ServerOption {
	return func(o *serverOptions) {
		o.Instance = instance
	}
}

func Color(color string) ServerOption {
	return func(o *serverOptions) {
		o.Color = color
	}
}

func NewServer(handler http.Handler, opts ...ServerOption) *Server {
	return &Server{
		handler: handler,
		o:       new(serverOptions).apply(opts...).init(),
	}
}

func (s *Server) Run(ctx context.Context) error {
	var err error
	var lis net.Listener
	lis, err = net.Listen("tcp", net.JoinHostPort("", strconv.Itoa(s.o.Port)))
	if err != nil {
		return err
	}

	registrar, err := s.newRegister(ctx, lis)
	if err != nil {
		return err
	}

	httpSrv := s.newHttpServer(ctx)

	errC := s.serve(httpSrv, lis)

	s.register(registrar)

	select {
	case err := <-errC:
		if err != nil {
			// if server failed to start, deregister
			s.deregister(registrar)
			return err
		}
	case <-ctx.Done():
		// graceful shutdown, deregister and shutdown
		s.deregister(registrar)
		err := s.shutdown(ctx, httpSrv)
		return errors.Join(ctx.Err(), err)
	}
	return nil
}

func (s *Server) serve(httpSrv *http.Server, lis net.Listener) chan error {
	var errC = make(chan error, 1)
	go func() {
		defer close(errC)
		var err error
		if httpSrv.TLSConfig != nil {
			err = httpSrv.ServeTLS(lis, "", "")
		} else {
			err = httpSrv.Serve(lis)
		}
		if err == nil {
			return
		}
		if errors.Is(err, http.ErrServerClosed) {
			return
		}
		errC <- err
	}()
	runtime.Gosched()
	return errC
}

func (s *Server) shutdown(ctx context.Context, httpSrv *http.Server) error {
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
		Handler:                      s.handler,
		DisableGeneralOptionsHandler: s.o.DisableGeneralOptionsHandler,
		TLSConfig:                    s.o.TLSConfig,
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

func (s *Server) newRegister(ctx context.Context, lis net.Listener) (sd.Registrar, error) {
	if s.o.RegistrarBuilder == nil {
		return nil, nil
	}
	instanceUrl, err := url.Parse(s.o.Instance)
	if err != nil {
		return nil, err
	}
	ip, port, err := internal.GlobalUnicastAddr(lis.Addr())
	if err != nil {
		return nil, err
	}
	return s.o.RegistrarBuilder.BuildRegistrar(ctx, instanceUrl, ip, port, s.o.Color)
}

func (s *Server) register(registrar sd.Registrar) {
	go func() {
		runtime.Gosched()
		if registrar != nil {
			registrar.Register()
		}
	}()
	runtime.Gosched()
}

func (s *Server) deregister(registrar sd.Registrar) {
	if registrar != nil {
		registrar.Deregister()
	}
}
