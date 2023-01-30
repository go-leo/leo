package grpc

import (
	"context"
	"errors"
	"fmt"
	"net"
	"runtime"
	"strconv"
	"sync"

	"github.com/go-leo/netx/addrx"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	"github.com/go-leo/leo/v2/registry"
)

type Service struct {
	Impl any
	Desc grpc.ServiceDesc
}

// Server grpc运行实体
type Server struct {
	host string
	port int
	// o 可选参数
	o *options
	// lis 网络端口监听接口
	lis net.Listener
	// services 包含业务Service的实现和描述信息
	services []Service
	// gRPCSrv 原生的grpc服务
	gRPCSrv *grpc.Server
	// healthSrv 健康检查服务
	healthSrv *health.Server
	id        string
	name      string
	version   string
	metaData  map[string]string
	startOnce sync.Once
	stopOnce  sync.Once
}

func NewServer(port int, services []Service, opts ...Option) (*Server, error) {
	// 监听端口
	lis, err := net.Listen("tcp", net.JoinHostPort("", strconv.Itoa(port)))
	if err != nil {
		return nil, err
	}
	// 如果上面的监听的端口为0，则会随机用一个可用的端口，所以需要回填。
	port = addrx.ExtractPort(lis.Addr())
	host, err := addrx.GlobalUnicastIPString()
	if err != nil {
		return nil, err
	}

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

	// 注册业务服务
	for _, service := range services {
		gRPCSrv.RegisterService(&service.Desc, service.Impl)
	}

	// 注册健康检查服务
	healthSrv := health.NewServer()
	for _, service := range services {
		healthSrv.SetServingStatus(service.Desc.ServiceName, grpc_health_v1.HealthCheckResponse_SERVING)
	}
	grpc_health_v1.RegisterHealthServer(gRPCSrv, healthSrv)

	// 注册反射服务
	reflection.Register(gRPCSrv)

	srv := &Server{
		host:      host,
		port:      port,
		o:         o,
		lis:       lis,
		services:  services,
		gRPCSrv:   gRPCSrv,
		startOnce: sync.Once{},
		stopOnce:  sync.Once{},
		healthSrv: healthSrv,
	}
	return srv, nil
}

func (s *Server) Start(ctx context.Context) error {
	if s.lis == nil {
		return errors.New("net listener is nil")
	}
	err := errors.New("server already started")
	s.startOnce.Do(func() {
		err = nil
		// 没有设置服务注册器，直接启动服务
		if s.o.Registrar == nil {
			err = s.startServer(ctx)
			return
		}
		// 启动服务并注册服务
		group, ctx := errgroup.WithContext(ctx)
		group.Go(func() error {
			return s.startServer(ctx)
		})
		group.Go(func() error {
			info, err := s.newServiceInfo()
			if err != nil {
				return err
			}
			runtime.Gosched()
			return s.o.Registrar.Register(ctx, info)
		})
		err = group.Wait()
	})
	return err
}

func (s *Server) Stop(ctx context.Context) error {
	err := errors.New("server already stopped")
	s.stopOnce.Do(func() {
		err = nil
		// 没有设置服务注册器，直接停止服务
		if s.o.Registrar == nil {
			err = s.stopServer(ctx)
			return
		}
		// 注销服务并停止服务
		group, ctx := errgroup.WithContext(ctx)
		group.Go(func() error {
			info, err := s.newServiceInfo()
			if err != nil {
				return err
			}
			return s.o.Registrar.Deregister(ctx, info)
		})
		group.Go(func() error {
			runtime.Gosched()
			return s.stopServer(ctx)
		})
		err = group.Wait()
	})
	return err
}

func (s *Server) String() string {
	return "grpc server"
}

func (s *Server) SetServiceInfo(getter registry.ServiceInfoGetter) {
	s.id = getter.ID()
	s.name = getter.Name()
	s.version = getter.Version()
	s.metaData = getter.MetaData()
}

func (s *Server) Services() []Service {
	return s.services
}

func (s *Server) Host() string {
	return s.host
}

func (s *Server) Port() int {
	return s.port
}

func (s *Server) startServer(ctx context.Context) error {
	if s.healthSrv != nil {
		s.healthSrv.Resume()
	}
	return s.gRPCSrv.Serve(s.lis)
}

func (s *Server) stopServer(ctx context.Context) error {
	if s.healthSrv != nil {
		s.healthSrv.Shutdown()
	}
	s.gRPCSrv.GracefulStop()
	return nil
}

func (s *Server) newServiceInfo() (*registry.ServiceInfo, error) {
	host, err := addrx.GlobalUnicastIPString()
	if err != nil {
		return nil, err
	}
	transport := registry.TransportGRPC
	id := fmt.Sprintf("%s.%s.%d", s.id, transport, s.port)
	serviceInfo := &registry.ServiceInfo{
		ID:        id,
		Name:      s.name,
		Transport: transport,
		Host:      host,
		Port:      s.port,
		Metadata:  s.metaData,
		Version:   s.version,
	}
	return serviceInfo, nil
}
