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

	"github.com/go-leo/stringx"

	"github.com/go-leo/leo/runner/net/http/header"
	"github.com/go-leo/leo/runner/net/http/internal/health"
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
	middlewares := append([]gin.HandlerFunc{header.GinMiddleware()}, o.GinMiddlewares...)
	mux.Use(middlewares...)
	// 设置健康检查
	healthSrv := health.NewServer()
	mux.Any(HealthCheckPath, health.HandlerFunc(healthSrv))
	if slicex.IsNotEmpty(o.NoRouteHandlers) {
		mux.NoRoute(o.NoRouteHandlers...)
	}
	if slicex.IsNotEmpty(o.NoMethodHandlers) {
		mux.NoMethod(o.NoMethodHandlers...)
	}
	// 基于服务描述ServiceDesc，将HTTPMethod、Path、handler等注册请求处理并包装Handler方法
	if o.GRPCClient != nil && o.ServiceDesc != nil {
		for _, methodDesc := range o.ServiceDesc.Methods {
			// 绑定Handler
			mux.Handle(methodDesc.HTTPMethod, methodDesc.Path, HandlerFunc(o.GRPCClient, o.ServiceDesc, methodDesc))
		}
	}
	// 注册其他自定义的非protoc-gen-go-leo生成的路由
	for _, router := range o.Routers {
		// 兼容
		if len(router.HTTPMethods) <= 0 || stringx.IsNotBlank(router.HTTPMethod) {
			router.HTTPMethods = append(router.HTTPMethods, router.HTTPMethod)
		}
		// 如果没有指定具体method，则绑定所有可能的Method
		if len(router.HTTPMethods) <= 0 {
			mux.Any(router.Path, router.HandlerFuncs...)
			continue
		}
		// 指定了method，就绑定到method上。
		for _, method := range router.HTTPMethods {
			mux.Handle(method, router.Path, router.HandlerFuncs...)
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
