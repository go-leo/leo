package grpcx

import (
	"context"
	"errors"
	"github.com/go-leo/gox/mapx"
	"github.com/go-leo/leo/v3/healthx"
	"github.com/go-leo/leo/v3/sdx"
	internalsd "github.com/go-leo/leo/v3/serverx/internal/sd"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"net"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Server struct {
	o        *options
	services sync.Map
}

type options struct {
	// Port is the port to listen on.
	Port int
	// ServerOptions
	ServerOptions []grpc.ServerOption
	// Builder creates a new sd.Registrar.
	Builder sdx.Builder
	// Instance is the instance to register.
	Instance string
	// Color is the color of the instance.
	Color string
	// ShutdownTimeout is the timeout for graceful shutdown.
	ShutdownTimeout *time.Duration
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

func ServerOptions(opts ...grpc.ServerOption) Option {
	return func(o *options) {
		o.ServerOptions = append(o.ServerOptions, opts...)
	}
}

func Builder(builder sdx.Builder) Option {
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

func ShutdownTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.ShutdownTimeout = &timeout
	}
}

func NewServer(opts ...Option) *Server {
	return &Server{
		o: new(options).apply(opts...).complete(),
	}
}

// RegisterService implements grpc.ServiceRegistrar.
func (s *Server) RegisterService(desc *grpc.ServiceDesc, ss any) {
	s.services.Store(desc, ss)
}

func (s *Server) Run(ctx context.Context) error {
	// listen port.
	lis, err := net.Listen("tcp", net.JoinHostPort("", strconv.Itoa(s.o.Port)))
	if err != nil {
		return err
	}

	// create registrar.
	registrar, err := internalsd.NewRegistrar(ctx, lis, s.o.Builder, s.o.Instance, s.o.Color)
	if err != nil {
		return err
	}

	// create grpc server.
	grpcSrv := grpc.NewServer(s.o.ServerOptions...)

	// register services.
	services := mapx.FromRanger[map[*grpc.ServiceDesc]any](&s.services)
	for desc, ss := range services {
		grpcSrv.RegisterService(desc, ss)
	}

	// register health check service.
	checker := healthx.NewChecker("grpc")
	grpc_health_v1.RegisterHealthServer(grpcSrv, checker)
	healthx.RegisterChecker(checker)

	// register reflection service.
	reflection.Register(grpcSrv)

	// server serve.
	serveErrC := s.serve(grpcSrv, lis, checker)

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
		err := s.shutdown(ctx, grpcSrv, checker)
		return errors.Join(ctx.Err(), err)
	}
}

func (s *Server) serve(grpcSrv *grpc.Server, lis net.Listener, checker healthx.Checker) chan error {
	var errC = make(chan error, 1)
	go func() {
		defer close(errC)
		checker.Resume()
		err := grpcSrv.Serve(lis)
		checker.Shutdown()
		if err == nil {
			return
		}
		errC <- err
	}()
	// make sure server serve finish
	runtime.Gosched()
	return errC
}

func (s *Server) shutdown(ctx context.Context, grpcSrv *grpc.Server, checker healthx.Checker) error {
	ctx = context.WithoutCancel(ctx)
	if s.o.ShutdownTimeout != nil {
		var cancelFunc context.CancelFunc
		ctx, cancelFunc = context.WithTimeout(ctx, *s.o.ShutdownTimeout)
		defer cancelFunc()
	}
	done := make(chan struct{})
	go func() {
		defer close(done)
		checker.Shutdown()
		grpcSrv.GracefulStop()
	}()
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
