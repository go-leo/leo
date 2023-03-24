package ginhttp

import (
	"context"
	"crypto/tls"
	"errors"
	"net"
	"net/http"
	"runtime"
	"strconv"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"

	"github.com/go-leo/gox/contextx"
	"github.com/go-leo/gox/netx/addrx"

	"codeup.aliyun.com/qimao/leo/leo/actuator"
	"codeup.aliyun.com/qimao/leo/leo/actuator/health"
	"codeup.aliyun.com/qimao/leo/leo/registry"
)

type Server struct {
	port      int
	host      string
	options   *options
	engine    *gin.Engine
	healthSrv *HealthServer
	lis       net.Listener
}

func (server *Server) Run(ctx context.Context) error {
	// listen port
	lis, err := server.listenPort()
	if err != nil {
		return err
	}
	server.lis = lis

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

	// new tls lis
	if server.options.TLSConf != nil {
		lis = tls.NewListener(lis, server.options.TLSConf)
	}
	return lis, nil
}

func (server *Server) runServer(ctx context.Context) error {
	// new http server
	httpSrv := &http.Server{
		Handler:           server.engine,
		ReadTimeout:       server.options.ReadTimeout,
		ReadHeaderTimeout: server.options.ReadHeaderTimeout,
		WriteTimeout:      server.options.WriteTimeout,
		IdleTimeout:       server.options.IdleTimeout,
		MaxHeaderBytes:    server.options.MaxHeaderBytes,
	}

	// http server async run serve
	serveErrC := make(chan error)
	go func() {
		defer close(serveErrC)
		server.healthSrv.Resume()
		err := httpSrv.Serve(server.lis)
		if errors.Is(err, http.ErrServerClosed) {
			return
		}
		if err != nil {
			serveErrC <- err
		}
	}()
	runtime.Gosched()

	// wait until context canceled or http server failed to serve
	select {
	case serveErr := <-serveErrC:
		// failed to serve, return serve error
		return serveErr
	case <-ctx.Done():
		// context canceled, shutdown server.
		ctx, _ := contextx.WithSignal(context.Background())
		if server.options.ShutdownTimeout > 0 {
			ctx, _ = context.WithTimeout(ctx, server.options.ShutdownTimeout)
		}
		server.healthSrv.Shutdown()
		return errors.Join(httpSrv.Shutdown(ctx), <-serveErrC)
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
		Scheme("http").
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
		return errors.Join(server.options.Registrar.Deregister(ctx, instance))
	}
}

func NewServer(port int, engine *gin.Engine, opts ...Option) *Server {
	o := new(options)
	o.apply(opts...)
	o.init()
	return &Server{
		port:      port,
		host:      "",
		options:   o,
		engine:    engine,
		healthSrv: &HealthServer{running: &atomic.Bool{}},
	}
}
