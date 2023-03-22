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
	"github.com/go-leo/gox/operator"

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

	// new tls listener
	if server.options.TLSConf != nil {
		lis = tls.NewListener(lis, server.options.TLSConf)
	}
	return lis, nil
}

func (server *Server) runServer(ctx context.Context) error {
	// listen port
	lis, err := server.listenPort()
	if err != nil {
		return err
	}

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
		err = httpSrv.Serve(lis)
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

	scheme := operator.Ternary(server.options.TLSConf != nil, "https", "http")
	instance := registry.NewServiceInstance(
		server.options.ID,
		server.options.Name,
		server.host,
		server.port,
		server.options.MetaData,
		scheme)

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
