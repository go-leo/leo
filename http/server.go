package http

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
	"runtime"
	"strconv"
	"sync"

	"github.com/go-leo/netx/addrx"
	"github.com/go-leo/netx/httpx"
	"github.com/go-leo/stringx"
	"golang.org/x/sync/errgroup"

	"github.com/go-leo/leo/v2/registry"
)

type Server struct {
	port      int
	o         *options
	lis       net.Listener
	httpSrv   *http.Server
	healthSrv *healthServer
	id        string
	name      string
	version   string
	metaData  map[string]string
	startOnce sync.Once
	stopOnce  sync.Once
}

func NewServer(port int, handler http.Handler, opts ...Option) (*Server, error) {
	// 监听端口
	lis, err := net.Listen("tcp", net.JoinHostPort("", strconv.Itoa(port)))
	if err != nil {
		return nil, err
	}
	// 如果上面的监听的端口为0，则会随机用一个可用的端口，所以需要回填。
	port = addrx.ExtractPort(lis.Addr())

	o := new(options)
	o.apply(opts...)
	o.init()

	// 创建http.Server
	httpSrv := newHttpServer(o)

	// 判断是否支持健康检查
	var healthSrv *healthServer
	if stringx.IsBlank(o.HealthCheckPath) {
		httpSrv.Handler = handler
	} else {
		// 设置健康检查
		healthSrv = newHealthServer(SERVING)
		httpSrv.Handler = newCompositeHandler(o, healthSrv, handler)
	}

	return &Server{
		port:      port,
		o:         o,
		lis:       lis,
		httpSrv:   httpSrv,
		healthSrv: healthSrv,
		startOnce: sync.Once{},
		stopOnce:  sync.Once{},
	}, nil
}

func (s *Server) String() string {
	return fmt.Sprintf("%s server", s.transport())
}

func (s *Server) Start(ctx context.Context) error {
	if s.lis == nil {
		return errors.New("net listener is nil")
	}
	err := errors.New("server already started")
	s.startOnce.Do(func() {
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

func (s *Server) ID() string {
	return s.id
}

func (s *Server) Name() string {
	return s.name
}

func (s *Server) Version() string {
	return s.version
}

func (s *Server) MetaData() map[string]string {
	return s.metaData
}

func (s *Server) SetID(id string) {
	s.id = id
}

func (s *Server) SetName(name string) {
	s.name = name
}

func (s *Server) SetVersion(version string) {
	s.version = version
}

func (s *Server) SetMetaData(metaData map[string]string) {
	s.metaData = metaData
}

func (s *Server) Scheme() string {
	if s.o.TLSConf != nil {
		return "https"
	}
	return "http"
}

func (s *Server) Port() int {
	return s.port
}

func (s *Server) startServer(ctx context.Context) error {
	if s.httpSrv != nil {
		s.healthSrv.Resume(ctx)
	}
	if s.o.TLSConf != nil {
		return s.httpSrv.Serve(tls.NewListener(s.lis, s.o.TLSConf))
	}
	return s.httpSrv.Serve(s.lis)
}

func (s *Server) stopServer(ctx context.Context) error {
	if s.httpSrv != nil {
		s.healthSrv.Shutdown(ctx)
	}
	return s.httpSrv.Shutdown(ctx)
}

func (s *Server) newServiceInfo() (*registry.ServiceInfo, error) {
	host, err := addrx.GlobalUnicastIPString()
	if err != nil {
		return nil, err
	}
	transport := s.transport()
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

func (s *Server) transport() string {
	transport := registry.TransportHTTP
	if s.o.TLSConf != nil {
		transport = registry.TransportHTTPS
	}
	return transport
}

func newHttpServer(o *options) *http.Server {
	httpSrv := &http.Server{
		ReadTimeout:       o.ReadTimeout,
		ReadHeaderTimeout: o.ReadHeaderTimeout,
		WriteTimeout:      o.WriteTimeout,
		IdleTimeout:       o.IdleTimeout,
		MaxHeaderBytes:    o.MaxHeaderBytes,
		TLSNextProto:      o.TLSNextProto,
		ConnState:         o.ConnState,
		ErrorLog:          o.ErrorLog,
		BaseContext:       o.BaseContext,
		ConnContext:       o.ConnContext,
	}
	return httpSrv
}

func newCompositeHandler(o *options, healthSrv *healthServer, handler http.Handler) *httpx.CompositeHandler {
	mux := http.NewServeMux()
	mux.HandleFunc(o.HealthCheckPath, healthSrv.ServeHTTP)
	compositeHandler := new(httpx.CompositeHandler)
	compositeHandler.AddHandler(mux, func(request *http.Request) bool {
		if request.URL.Path == o.HealthCheckPath {
			return true
		}
		return false
	})
	compositeHandler.AddHandler(handler, func(_ *http.Request) bool { return true })
	return compositeHandler
}
