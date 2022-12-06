package server

import (
	"context"
	"crypto/tls"
	"errors"
	"net"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-leo/slicex"

	"github.com/go-leo/leo/v2/runner/net/http/internal/health"
)

const HealthCheckPath = "/health/check"

type Server struct {
	o         *options
	lis       net.Listener
	httpSrv   *http.Server
	healthSrv *health.Server
	startOnce sync.Once
	stopOnce  sync.Once
}

func New(lis net.Listener, opts ...Option) *Server {
	o := new(options)
	o.apply(opts)
	o.init()
	// 设置Gin的Mode
	gin.SetMode(gin.ReleaseMode)
	// 创建Gin的Engine
	mux := gin.New()
	// 设置middlewares
	mux.Use(o.GinMiddlewares...)

	// 设置健康检查
	healthSrv := health.NewServer()
	mux.Any(HealthCheckPath, health.HandlerFunc(healthSrv))

	if slicex.IsNotEmpty(o.NoRouteHandlers) {
		mux.NoRoute(o.NoRouteHandlers...)
	}
	if slicex.IsNotEmpty(o.NoMethodHandlers) {
		mux.NoMethod(o.NoMethodHandlers...)
	}
	// 注册其他自定义的非protoc-gen-go-leo生成的路由
	for _, router := range o.Routes {
		// 如果没有指定具体method，则绑定所有可能的Method
		if len(router.Methods()) <= 0 {
			mux.Any(router.Path(), router.Handlers()...)
			continue
		}
		// 指定了method，就绑定到method上。
		for _, method := range router.Methods() {
			mux.Handle(method, router.Path(), router.Handlers()...)
		}
	}
	// 创建http.Server
	httpSrv := &http.Server{
		Handler:           mux,
		ReadTimeout:       o.ReadTimeout,
		ReadHeaderTimeout: 0,
		WriteTimeout:      o.WriteTimeout,
		IdleTimeout:       o.IdleTimeout,
		MaxHeaderBytes:    o.MaxHeaderBytes,
	}
	srv := &Server{
		o:         o,
		lis:       lis,
		healthSrv: healthSrv,
		httpSrv:   httpSrv,
	}
	return srv
}

func (s *Server) String() string {
	return "HTTP"
}

func (s *Server) Start(ctx context.Context) error {
	err := errors.New("server already started")
	s.startOnce.Do(func() {
		err = nil
		if s.lis == nil {
			err = errors.New("net listener is nil")
			return
		}
		s.healthSrv.Resume()
		if s.o.TLSConf != nil {
			err = s.httpSrv.Serve(tls.NewListener(s.lis, s.o.TLSConf))
			return
		}
		err = s.httpSrv.Serve(s.lis)
	})
	return err
}

func (s *Server) Stop(ctx context.Context) error {
	err := errors.New("server already stopped")
	s.stopOnce.Do(func() {
		err = nil
		s.healthSrv.Shutdown()
		err = s.httpSrv.Shutdown(ctx)
	})
	return err
}
