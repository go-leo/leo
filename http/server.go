package http

import (
	"context"
	"crypto/tls"
	"errors"
	"net"
	"net/http"
	"sync"

	"github.com/go-leo/netx/httpx"
)

type Server struct {
	o         *options
	lis       net.Listener
	httpSrv   *http.Server
	healthSrv *healthServer
	startOnce sync.Once
	stopOnce  sync.Once
}

func NewServer(lis net.Listener, handler http.Handler, opts ...Option) *Server {
	o := new(options)
	o.apply(opts...)
	o.init()

	// 设置健康检查
	healthSrv := newHealthServer(SERVING)
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

	// 创建http.Server
	httpSrv := &http.Server{
		Handler:           compositeHandler,
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

	return &Server{
		o:         o,
		lis:       lis,
		httpSrv:   httpSrv,
		healthSrv: healthSrv,
		startOnce: sync.Once{},
		stopOnce:  sync.Once{},
	}
}

func (s *Server) String() string {
	return "HTTP"
}

func (s *Server) Start(ctx context.Context) error {
	if s.lis == nil {
		return errors.New("net listener is nil")
	}
	err := errors.New("server already started")
	s.startOnce.Do(func() {
		s.healthSrv.Resume(ctx)
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
		s.healthSrv.Shutdown(ctx)
		err = s.httpSrv.Shutdown(ctx)
	})
	return err
}
