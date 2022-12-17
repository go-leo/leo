package server

import (
	"context"
	"errors"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type Service struct {
	Impl any
	Desc grpc.ServiceDesc
}

// Server grpc运行实体
type Server struct {
	// o 可选参数
	o *options
	// lis 网络端口监听接口
	lis net.Listener
	// service 包含业务Service的实现和描述信息
	service Service
	// gRPCSrv 原生的grpc服务
	gRPCSrv *grpc.Server
	// healthSrv 健康检查服务
	healthSrv *health.Server
	startOnce sync.Once
	stopOnce  sync.Once
}

func (s *Server) Start(_ context.Context) error {
	if s.lis == nil {
		return errors.New("net listener is nil")
	}
	err := errors.New("server already started")
	s.startOnce.Do(func() {
		err = nil
		s.healthSrv.Resume()
		err = s.gRPCSrv.Serve(s.lis)
	})
	return err
}

func (s *Server) Stop(_ context.Context) error {
	err := errors.New("server already stopped")
	s.stopOnce.Do(func() {
		err = nil
		s.healthSrv.Shutdown()
		s.gRPCSrv.GracefulStop()
	})
	return err
}

func (s *Server) String() string {
	return "gRPC"
}

func New(lis net.Listener, service Service, opts ...Option) *Server {
	o := new(options)
	o.apply(opts)
	o.init()

	var serverOptions []grpc.ServerOption
	serverOptions = append(serverOptions, o.serverOptions...)
	if o.tlsConf != nil {
		serverOptions = append(serverOptions, grpc.Creds(credentials.NewTLS(o.tlsConf)))
	}
	serverOptions = append(serverOptions, grpc.ChainUnaryInterceptor(o.unaryInterceptors...))
	gRPCSrv := grpc.NewServer(serverOptions...)

	// register health check service
	healthSrv := health.NewServer()
	healthSrv.SetServingStatus(service.Desc.ServiceName, grpc_health_v1.HealthCheckResponse_SERVING)
	grpc_health_v1.RegisterHealthServer(gRPCSrv, healthSrv)

	// register reflection service on gRPC server.
	reflection.Register(gRPCSrv)

	// register business service
	gRPCSrv.RegisterService(&service.Desc, service.Impl)

	srv := &Server{
		o:         o,
		lis:       lis,
		service:   service,
		gRPCSrv:   gRPCSrv,
		startOnce: sync.Once{},
		stopOnce:  sync.Once{},
		healthSrv: healthSrv,
	}
	return srv
}
