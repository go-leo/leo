package lgrpc

import (
	"context"
	"errors"
	"net"
	"runtime"
	"strconv"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	grpchealth "google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	"github.com/go-leo/gox/contextx"
	"github.com/go-leo/gox/netx/addrx"

	"codeup.aliyun.com/qimao/leo/leo/actuator"
	"codeup.aliyun.com/qimao/leo/leo/actuator/health"
	"codeup.aliyun.com/qimao/leo/leo/registry"
)

// Server 服务
type Server struct {
	port               int
	host               string
	options            *options
	gRPCSrv            *grpc.Server
	healthSrv          *grpchealth.Server
	lis                net.Listener
	registerServiceMap map[*grpc.ServiceDesc]any
}

func (server *Server) Run(ctx context.Context) error {
	// listen port
	lis, err := server.listenPort()
	if err != nil {
		return err
	}
	server.lis = lis

	stopGroup, newCtx := errgroup.WithContext(ctx)
	stopGroup.Go(func() error { return server.runServer(newCtx) })
	runtime.Gosched()
	stopGroup.Go(func() error { return server.runRegistrar(newCtx) })
	runtime.Gosched()
	return stopGroup.Wait()
}

func (server *Server) ActuatorHandler() actuator.Handler {
	return &actuatorHandler{server: server}
}

func (server *Server) HealthChecker() health.Checker {
	return &healthChecker{server: server}
}

func (server *Server) RegisterService(sd *grpc.ServiceDesc, ss any) {
	server.registerServiceMap[sd] = ss
}

func (server *Server) GetServiceInfo() map[string]grpc.ServiceInfo {
	if server.gRPCSrv == nil {
		return nil
	}
	return server.gRPCSrv.GetServiceInfo()
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
	for sd, ss := range server.registerServiceMap {
		server.gRPCSrv.RegisterService(sd, ss)
	}
	// grpc server async run serve
	serveErrC := make(chan error)
	go func() {
		defer close(serveErrC)
		server.healthSrv.Resume()
		err := server.gRPCSrv.Serve(server.lis)
		if errors.Is(err, grpc.ErrServerStopped) {
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

	instance := registry.Builder().
		ID(server.options.ID).
		Name(server.options.Name).
		IP(server.host).
		Port(server.port).
		Metadata(server.options.MetaData).
		Scheme("grpc").
		Build()

	// registrar register instance
	err := server.options.Registrar.Register(ctx, instance)
	if err != nil {
		// failed to register. return register error
		return err
	}

	// wait until context canceled or registrar failed to register
	select {
	case <-ctx.Done():
		// context canceled, deregister server.
		ctx, _ := contextx.WithSignal(context.Background())
		if server.options.ShutdownTimeout > 0 {
			ctx, _ = context.WithTimeout(ctx, server.options.ShutdownTimeout)
		}
		// return register and deregister error if has error
		return server.options.Registrar.Deregister(ctx, instance)
	}
}

func NewServer(port int, opts ...Option) *Server {
	o := new(options)
	o.apply(opts)
	o.init()
	srv := &Server{
		options:            o,
		port:               port,
		registerServiceMap: make(map[*grpc.ServiceDesc]any),
	}
	return srv
}
