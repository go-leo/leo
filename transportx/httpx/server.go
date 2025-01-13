package httpx

import (
	"context"
	"crypto/tls"
	"errors"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	kitlog "github.com/go-kit/log"
	"github.com/go-leo/gox/netx/addrx"
	"github.com/go-leo/leo/v3/logx"
	"github.com/go-leo/leo/v3/sdx"
	"github.com/go-leo/leo/v3/sdx/noop"
	"github.com/google/uuid"
	stdconsul "github.com/hashicorp/consul/api"
	"log"
	"net"
	"net/http"
	"time"
)

type Server struct {
	o       *serverOptions
	handler http.Handler
}

type serverOptions struct {
	Addr string

	DisableGeneralOptionsHandler bool
	TLSConfig                    *tls.Config
	ReadTimeout                  time.Duration
	ReadHeaderTimeout            time.Duration
	WriteTimeout                 time.Duration
	IdleTimeout                  time.Duration
	MaxHeaderBytes               int
	TLSNextProto                 map[string]func(*http.Server, *tls.Conn, http.Handler)
	ConnState                    func(net.Conn, http.ConnState)

	ShutdownTimeout *time.Duration

	RegistrarFactory sdx.Builder
}

type ServerOption func(o *serverOptions)

func (o *serverOptions) apply(opts ...ServerOption) *serverOptions {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func (o *serverOptions) init() *serverOptions {
	if o.RegistrarFactory == nil {
		o.RegistrarFactory = noop.Builder{}
	}
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

func NewServer(handler http.Handler, opts ...ServerOption) *Server {
	return &Server{
		handler: handler,
		o:       new(serverOptions).apply(opts...).init(),
	}
}

func (s *Server) Run(ctx context.Context) error {
	var err error
	var lis net.Listener
	lis, err = net.Listen("tcp", s.o.Addr)
	if err != nil {
		return err
	}
	host, port, err := net.SplitHostPort(lis.Addr().String())
	if err != nil {
		return err
	}
	var ip string
	if !addrx.IsGlobalUnicastIP(net.ParseIP(host)) {
		ip, err = addrx.GlobalUnicastIPString()
		if err != nil {
			return err
		}
	}

	var color string
	var registrar sd.Registrar

	client, err := stdconsul.NewClient(&stdconsul.Config{
		Address:    "localhost:8500",
		Datacenter: "dc1",
	})
	if err != nil {
		return err
	}
	_ = port
	registration := &stdconsul.AgentServiceRegistration{
		ID:      uuid.NewString(),
		Name:    "demo.http",
		Tags:    []string{color},
		Port:    0,
		Address: ip,
	}
	registrar = consul.NewRegistrar(consul.NewClient(client), registration, logx.L())

	httpSrv := &http.Server{
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
		BaseContext:                  func(listener net.Listener) context.Context { return ctx },
		ConnContext:                  func(ctx context.Context, c net.Conn) context.Context { return ctx },
	}

	var errC = make(chan error, 1)
	go func() {
		defer close(errC)
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

	go func() {
		registrar.Register()
	}()

	select {
	case err := <-errC:
		if err != nil {
			if registrar != nil {
				registrar.Deregister()
			}
			return err
		}
	case <-ctx.Done():
		registrar.Deregister()
		ctx = context.WithoutCancel(ctx)
		if s.o.ShutdownTimeout != nil {
			var cancelFunc context.CancelFunc
			ctx, cancelFunc = context.WithTimeout(ctx, *s.o.ShutdownTimeout)
			defer cancelFunc()
		}
		return errors.Join(ctx.Err(), httpSrv.Shutdown(ctx))
	}
	return nil
}
