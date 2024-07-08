package httpx

import (
	"context"
	"crypto/tls"
	"errors"
	kitlog "github.com/go-kit/log"
	"github.com/go-leo/leo/v3/logx"
	"github.com/go-leo/leo/v3/runner"
	"github.com/go-leo/leo/v3/transportx"
	"log"
	"net"
	"net/http"
	"time"
)

var _ runner.StartStopper = (*server)(nil)

type server struct {
	*http.Server
	lis net.Listener
	o   *serverOptions
}

func (s *server) Start(ctx context.Context) error {
	s.Server.BaseContext = func(listener net.Listener) context.Context {
		return ctx
	}
	s.Server.ConnContext = func(ctx context.Context, c net.Conn) context.Context {
		return ctx
	}
	var err error
	if s.Server.TLSConfig != nil {
		err = s.Server.ServeTLS(s.lis, "", "")
	} else {
		err = s.Server.Serve(s.lis)
	}
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}

func (s *server) Stop(ctx context.Context) error {
	ctx = context.WithoutCancel(ctx)
	if s.o.ShutdownTimeout == nil {
		return s.Server.Shutdown(ctx)
	}
	ctx, cancelFunc := context.WithTimeout(ctx, *s.o.ShutdownTimeout)
	defer cancelFunc()
	return s.Server.Shutdown(ctx)
}

type serverOptions struct {
	DisableGeneralOptionsHandler bool
	TLSConfig                    *tls.Config
	ReadTimeout                  time.Duration
	ReadHeaderTimeout            time.Duration
	WriteTimeout                 time.Duration
	IdleTimeout                  time.Duration
	MaxHeaderBytes               int
	ShutdownTimeout              *time.Duration
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

func ShutdownTimeout(timeout time.Duration) ServerOption {
	return func(o *serverOptions) {
		o.ShutdownTimeout = &timeout
	}
}

func ServerFactory(handler http.Handler, opts ...ServerOption) transportx.ServerFactory {
	o := new(serverOptions).apply(opts...).init()
	httpSrv := &http.Server{
		Handler:                      handler,
		DisableGeneralOptionsHandler: o.DisableGeneralOptionsHandler,
		TLSConfig:                    o.TLSConfig,
		ReadTimeout:                  o.ReadTimeout,
		ReadHeaderTimeout:            o.ReadHeaderTimeout,
		WriteTimeout:                 o.WriteTimeout,
		IdleTimeout:                  o.IdleTimeout,
		MaxHeaderBytes:               o.MaxHeaderBytes,
		ErrorLog:                     log.New(kitlog.NewStdlibAdapter(logx.L()), "", 0),
	}
	return func(lis net.Listener, args any) transportx.Server {
		return &server{Server: httpSrv, lis: lis}
	}
}
