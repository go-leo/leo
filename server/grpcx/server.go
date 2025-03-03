package grpcx

import (
	"context"
	"errors"
	"github.com/go-kit/kit/sd"
	"github.com/go-leo/leo/v3/sdx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"net"
	"net/url"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Server struct {
	o *serverOptions

	mux      sync.RWMutex
	services map[*grpc.ServiceDesc]any
}

type serverOptions struct {
	Port int

	ServerOptions []grpc.ServerOption

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

func ServerOptions(opts ...grpc.ServerOption) ServerOption {
	return func(o *serverOptions) {
		o.ServerOptions = append(o.ServerOptions, opts...)
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

func NewServer(opts ...ServerOption) *Server {
	return &Server{
		o:        new(serverOptions).apply(opts...).init(),
		mux:      sync.RWMutex{},
		services: make(map[*grpc.ServiceDesc]any),
	}
}

func (s *Server) RegisterService(sd *grpc.ServiceDesc, ss any) {
	s.mux.Lock()
	s.services[sd] = ss
	s.mux.Unlock()
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

	grpcSrv := grpc.NewServer(s.o.ServerOptions...)

	// register services.
	s.mux.RLock()
	services := s.services
	s.mux.RUnlock()
	for desc, impl := range services {
		grpcSrv.RegisterService(desc, impl)
	}

	// register health check service.
	healthSrv := health.NewServer()
	for desc := range services {
		healthSrv.SetServingStatus(desc.ServiceName, grpc_health_v1.HealthCheckResponse_SERVING)
	}
	grpc_health_v1.RegisterHealthServer(grpcSrv, healthSrv)

	// register reflection service.
	reflection.Register(grpcSrv)

	errC := s.serve(grpcSrv, lis, healthSrv)

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
		err := s.shutdown(ctx, grpcSrv, healthSrv)
		return errors.Join(ctx.Err(), err)
	}
	return nil
}

func (s *Server) serve(grpcSrv *grpc.Server, lis net.Listener, healthSrv *health.Server) chan error {
	var errC = make(chan error, 1)
	go func() {
		defer close(errC)
		healthSrv.Resume()
		err := grpcSrv.Serve(lis)
		if err == nil {
			return
		}
		errC <- err
	}()
	runtime.Gosched()
	return errC
}

func (s *Server) shutdown(ctx context.Context, grpcSrv *grpc.Server, healthSrv *health.Server) error {
	ctx = context.WithoutCancel(ctx)
	if s.o.ShutdownTimeout != nil {
		var cancelFunc context.CancelFunc
		ctx, cancelFunc = context.WithTimeout(ctx, *s.o.ShutdownTimeout)
		defer cancelFunc()
	}
	done := make(chan struct{})
	go func() {
		defer close(done)
		healthSrv.Shutdown()
		grpcSrv.GracefulStop()
	}()
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
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
