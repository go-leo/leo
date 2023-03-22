package lgrpc

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"reflect"
	"runtime"
	"strconv"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	grpchealth "google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	"codeup.aliyun.com/qimao/leo/leo/actuator"
	"codeup.aliyun.com/qimao/leo/leo/actuator/health"
	"codeup.aliyun.com/qimao/leo/leo/internal/contextx"
	"codeup.aliyun.com/qimao/leo/leo/internal/netx/addrx"
	"codeup.aliyun.com/qimao/leo/leo/internal/reflectx"
	"codeup.aliyun.com/qimao/leo/leo/registry"
)

// Binder 服务绑定
type Binder struct {
	// ServerImpl 服务的实现
	ServerImpl any
	// RegisterFunc 注册方法
	RegisterFunc any
}

// Server 服务
type Server struct {
	port      int
	host      string
	options   *options
	gRPCSrv   *grpc.Server
	healthSrv *grpchealth.Server
	binders   []Binder
}

func (server *Server) Run(ctx context.Context) error {
	stopGroup, ctx := errgroup.WithContext(ctx)
	stopGroup.Go(func() error { return server.runServer(ctx) })
	runtime.Gosched()
	stopGroup.Go(func() error { return server.runRegistrar(ctx) })
	runtime.Gosched()
	return stopGroup.Wait()
}

func (server *Server) ActuatorHandler() actuator.Handler {
	return &actuatorHandler{server: server}
}

func (server *Server) HealthChecker() health.Checker {
	return &healthChecker{server: server}
}

func (server *Server) listenPort() (net.Listener, error) {
	// get global unicast ip
	host, err := addrx.GlobalUnicastIPString()
	if err != nil {
		return nil, err
	}
	server.host = host

	// listen port
	address := net.JoinHostPort("", strconv.Itoa(server.port))
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	// write back port
	server.port = addrx.ExtractPort(lis.Addr())
	return lis, nil
}

func (server *Server) runServer(ctx context.Context) error {
	// listen port
	lis, err := server.listenPort()
	if err != nil {
		return err
	}

	var serverOptions []grpc.ServerOption
	serverOptions = append(serverOptions, server.options.ServerOptions...)
	serverOptions = append(serverOptions, grpc.ChainUnaryInterceptor(server.options.UnaryInterceptors...))
	serverOptions = append(serverOptions, grpc.ChainStreamInterceptor(server.options.StreamInterceptors...))
	server.gRPCSrv = grpc.NewServer(serverOptions...)

	// register health check service
	server.healthSrv = grpchealth.NewServer()
	server.healthSrv.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
	grpc_health_v1.RegisterHealthServer(server.gRPCSrv, server.healthSrv)

	// register reflection service on gRPC server.
	reflection.Register(server.gRPCSrv)

	// register business service
	gRPCSrvVal := reflect.ValueOf(server.gRPCSrv)
	for _, register := range server.binders {
		registerFunc := reflectx.Indirect(register.RegisterFunc)
		val := reflect.ValueOf(registerFunc)
		if val.Kind() != reflect.Func {
			return fmt.Errorf("registerfunc is not a func kind")
		}
		_ = val.Call([]reflect.Value{gRPCSrvVal, reflect.ValueOf(register.ServerImpl)})
	}

	// grpc server async run serve
	serveErrC := make(chan error)
	go func() {
		defer close(serveErrC)
		server.healthSrv.Resume()
		err := server.gRPCSrv.Serve(lis)
		if errors.Is(err, http.ErrServerClosed) {
			return
		}
		if err != nil {
			serveErrC <- err
		}
	}()
	runtime.Gosched()

	// wait until context canceled or grpc server failed to serve
	select {
	case serveErr := <-serveErrC:
		// failed to serve, return serve error
		return serveErr
	case <-ctx.Done():
		// context canceled, shutdown server.
		server.healthSrv.Shutdown()
		server.gRPCSrv.GracefulStop()
		return nil
	}
}

func (server *Server) runRegistrar(ctx context.Context) error {
	if server.options.Registrar == nil {
		return nil
	}

	scheme := "grpc"
	instance := registry.NewServiceInstance(
		server.options.ID,
		server.options.Name,
		server.host,
		server.port,
		server.options.MetaData,
		scheme,
	)

	// registrar async run register
	registerErrC := make(chan error)
	go func() {
		defer close(registerErrC)
		err := server.options.Registrar.Register(ctx, instance)
		if errors.Is(err, registry.ErrServiceDeregistered) {
			return
		}
		if err != nil {
			registerErrC <- err
		}
	}()
	runtime.Gosched()

	// wait until context canceled or registrar failed to register
	select {
	case registerErr := <-registerErrC:
		// failed to register. return register error
		return registerErr
	case <-ctx.Done():
		// context canceled, deregister server.
		ctx, _ := contextx.WithSignal(context.Background())
		if server.options.ShutdownTimeout > 0 {
			ctx, _ = context.WithTimeout(ctx, server.options.ShutdownTimeout)
		}
		// return register and deregister error if has error
		return errors.Join(server.options.Registrar.Deregister(ctx, instance), <-registerErrC)
	}
}

func NewServer(port int, binders []Binder, opts ...Option) *Server {
	o := new(options)
	o.apply(opts)
	o.init()
	srv := &Server{
		options: o,
		port:    port,
		binders: binders,
	}
	return srv
}
