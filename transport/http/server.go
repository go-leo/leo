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
	"sync/atomic"

	"codeup.aliyun.com/qimao/leo/leo/actuator"
	"codeup.aliyun.com/qimao/leo/leo/actuator/health"
	"github.com/gin-gonic/gin"
	"github.com/go-leo/gox/contextx"
	"github.com/go-leo/gox/netx/addrx"
	"github.com/go-leo/gox/operator"
	"github.com/go-leo/gox/syncx/brave"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	port    int
	host    string
	o       *options
	httpSrv *http.Server
	engine  *gin.Engine
	running *atomic.Bool
}

func (s *Server) Run(ctx context.Context) error {
	startErrC := brave.GoE(func() error {
		return s.start(ctx)
	}, func(p any) error {
		return fmt.Errorf("panic triggered, %+v", p)
	})
	s.running.Store(true)
	var startErr error
	select {
	case <-ctx.Done():
	case startErr = <-startErrC:
	}
	s.running.Store(false)
	stopErr := s.stop(ctx)
	return errors.Join(startErr, stopErr)
}

func (s *Server) ID() string {
	return s.o.ID
}

func (s *Server) Name() string {
	return s.Name()
}

func (s *Server) Kind() string {
	return operator.Ternary(s.o.TLSConf == nil, "http", "https")
}

func (s *Server) Host() string {
	return s.host
}

func (s *Server) Port() int {
	return s.port
}

func (s *Server) Metadata() map[string]string {
	return s.o.MetaData
}

func (s *Server) Address() string {
	return net.JoinHostPort(s.host, strconv.Itoa(s.port))
}

func (s *Server) ActuatorHandler() actuator.Handler {
	return &actuatorHandler{server: s}
}

func (s *Server) HealthChecker() health.Checker {
	return &healthChecker{server: s}
}

func (s *Server) start(ctx context.Context) error {
	startGroup, ctx := errgroup.WithContext(ctx)
	startGroup.Go(s.startServer())
	runtime.Gosched()
	startGroup.Go(s.register(ctx))
	runtime.Gosched()
	return startGroup.Wait()
}

func (s *Server) stop(ctx context.Context) error {
	stopGroup, ctx := errgroup.WithContext(ctx)
	stopGroup.Go(s.deregister(ctx))
	runtime.Gosched()
	stopGroup.Go(s.stopServer())
	runtime.Gosched()
	return stopGroup.Wait()
}

func (s *Server) startServer() func() error {
	return func() error {
		if s.httpSrv == nil {
			return nil
		}
		// 监听端口
		lis, err := net.Listen("tcp", net.JoinHostPort("", strconv.Itoa(s.port)))
		if err != nil {
			return err
		}
		port := addrx.ExtractPort(lis.Addr())
		s.port = port
		host, err := addrx.GlobalUnicastIPString()
		if err != nil {
			return err
		}
		s.host = host

		if s.o.TLSConf != nil {
			return s.httpSrv.Serve(tls.NewListener(lis, s.o.TLSConf))
		}
		return s.httpSrv.Serve(lis)
	}
}

func (s *Server) stopServer() func() error {
	return func() error {
		if s.httpSrv == nil {
			return nil
		}
		ctx, _ := contextx.WithSignal(context.Background())
		if s.o.ShutdownTimeout >= 0 {
			ctx, _ = context.WithTimeout(context.Background(), s.o.ShutdownTimeout)
		}
		return s.httpSrv.Shutdown(ctx)
	}
}

func (s *Server) register(ctx context.Context) func() error {
	return func() error {
		if s.o.Registrar == nil {
			return nil
		}
		runtime.Gosched()
		return s.o.Registrar.Register(ctx, s)
	}
}

func (s *Server) deregister(ctx context.Context) func() error {
	return func() error {
		if s.o.Registrar == nil {
			return nil
		}
		runtime.Gosched()
		return s.o.Registrar.Register(ctx, s)
	}
}

func (s *Server) isRunning() bool {
	return s.running.Load()
}

func NewServer(port int, engine *gin.Engine, opts ...Option) *Server {
	o := new(options)
	o.apply(opts...)
	o.init()

	httpSrv := &http.Server{
		Handler:           engine,
		ReadTimeout:       o.ReadTimeout,
		ReadHeaderTimeout: o.ReadHeaderTimeout,
		WriteTimeout:      o.WriteTimeout,
		IdleTimeout:       o.IdleTimeout,
		MaxHeaderBytes:    o.MaxHeaderBytes,
	}

	srv := &Server{
		port:    port,
		host:    "",
		o:       o,
		httpSrv: httpSrv,
		engine:  engine,
		running: &atomic.Bool{},
	}

	return srv
}
