package management

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/go-leo/leo/v2/runner/management/router/app"
	"github.com/go-leo/leo/v2/runner/management/router/config"
	"github.com/go-leo/leo/v2/runner/management/router/env"
	"github.com/go-leo/leo/v2/runner/management/router/health"
	"github.com/go-leo/leo/v2/runner/management/router/metric"
	"github.com/go-leo/leo/v2/runner/management/router/profile"
	"github.com/go-leo/leo/v2/runner/management/router/restart"
	"github.com/go-leo/leo/v2/runner/management/router/server"
	"github.com/go-leo/leo/v2/runner/management/router/shutdown"
	"github.com/go-leo/leo/v2/runner/management/router/system"
	"github.com/go-leo/leo/v2/runner/management/router/task"
	httpserver "github.com/go-leo/leo/v2/runner/net/http/server"
	crontask "github.com/go-leo/leo/v2/runner/task/cron"
	pubsubtask "github.com/go-leo/leo/v2/runner/task/pubsub"
)

type Router struct {
	HttpMethod   string
	Path         string
	HandlerFuncs []gin.HandlerFunc
}

type options struct {
	GinMiddlewares  []gin.HandlerFunc
	Routers         []Router
	TLSConf         *tls.Config
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	MaxHeaderBytes  int
	HTTPHealthCheck *health.HttpOptions
	GRPCHealthCheck *health.GRPCOptions
	ExitSignals     []os.Signal
	RestartSignals  []os.Signal
	GRPCMapping     *server.GRPCMapping
	HTTPMapping     *server.HTTPMapping
	CronJobs        []*crontask.Job
	SubscriberJobs  []*pubsubtask.Job
}

type Option func(o *options)

func (o *options) apply(opts []Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *options) init() {
}

func GinMiddlewares(middlewares ...gin.HandlerFunc) Option {
	return func(o *options) {
		o.GinMiddlewares = middlewares
	}
}

func ReadTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.ReadTimeout = timeout
	}
}

func WriteTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.WriteTimeout = timeout
	}
}

func IdleTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.IdleTimeout = timeout
	}
}

func MaxHeaderBytes(size int) Option {
	return func(o *options) {
		o.MaxHeaderBytes = size
	}
}

func TLS(conf *tls.Config) Option {
	return func(o *options) {
		o.TLSConf = conf
	}
}

func Routers(routers ...Router) Option {
	return func(o *options) {
		o.Routers = routers
	}
}

func GRPCHealthCheck(target string, TLSConf *tls.Config, timeout time.Duration) Option {
	return func(o *options) {
		o.GRPCHealthCheck = &health.GRPCOptions{
			Target:  target,
			TLSConf: TLSConf,
			Timeout: timeout,
		}
	}
}

func HTTPHealthCheck(target string, TLSConf *tls.Config, timeout time.Duration) Option {
	return func(o *options) {
		o.HTTPHealthCheck = &health.HttpOptions{
			Target:  target,
			TLSConf: TLSConf,
			Timeout: timeout,
		}
	}
}

func GRPC(serviceDesc *grpc.ServiceDesc) Option {
	return func(o *options) {
		methodNames := make([]string, 0, len(serviceDesc.Methods))
		for _, method := range serviceDesc.Methods {
			methodNames = append(methodNames, fmt.Sprintf("/%s/%s", serviceDesc.ServiceName, method.MethodName))
		}
		o.GRPCMapping = &server.GRPCMapping{
			FullMethods: methodNames,
		}
	}
}

func HTTPRoutes(routes []httpserver.Route, richRoutes []httpserver.RichRoute) Option {
	return func(o *options) {
		routers := make([]server.HTTPRoute, 0, len(routes))
		for _, router := range routes {
			routers = append(routers, server.HTTPRoute{
				Methods: []string{router.Method()},
				Path:    router.Path(),
			})
		}
		for _, router := range richRoutes {
			routers = append(routers, server.HTTPRoute{
				Methods: router.Methods(),
				Path:    router.Path(),
			})
		}
		if o.HTTPMapping == nil {
			o.HTTPMapping = new(server.HTTPMapping)
		}
		o.HTTPMapping.HTTPRoutes = append(o.HTTPMapping.HTTPRoutes, routers...)
	}
}

func Cron(jobs []*crontask.Job) Option {
	return func(o *options) {
		o.CronJobs = jobs
	}
}

func Subscriber(jobs []*pubsubtask.Job) Option {
	return func(o *options) {
		o.SubscriberJobs = jobs
	}
}

func ShutdownSignals(signals []os.Signal) Option {
	return func(o *options) {
		o.ExitSignals = signals
	}
}

func RestartSignals(signals []os.Signal) Option {
	return func(o *options) {
		o.RestartSignals = signals
	}
}

type Server struct {
	o         *options
	lis       net.Listener
	httpSrv   *http.Server
	startOnce sync.Once
	stopOnce  sync.Once
}

func New(lis net.Listener, opts ...Option) *Server {
	o := new(options)
	o.apply(opts)
	o.init()
	gin.SetMode(gin.ReleaseMode)
	mux := gin.New()
	mux.Use(gin.Recovery())
	mux.Use(o.GinMiddlewares...)
	rg := mux.Group("/management")
	// register grpc and http mapping
	server.Route(rg, o.GRPCMapping, o.HTTPMapping)
	// register cron and sub jobs
	task.Route(rg, o.CronJobs, o.SubscriberJobs)
	// register profile
	profile.Route(rg)
	// register health
	health.Route(rg, o.HTTPHealthCheck, o.GRPCHealthCheck)
	// register system
	system.Route(rg)
	// register metrics
	metric.Route(rg)
	// register config
	config.Route(rg)
	// register env
	env.Route(rg)
	// register app
	app.Route(rg)
	// register shutdown
	shutdown.Route(rg, o.ExitSignals)
	// register restart
	restart.Route(rg, o.RestartSignals)
	for _, router := range o.Routers {
		rg.Handle(router.HttpMethod, router.Path, router.HandlerFuncs...)
	}
	httpSrv := &http.Server{
		Handler:           mux,
		ReadTimeout:       o.ReadTimeout,
		ReadHeaderTimeout: 0,
		WriteTimeout:      o.WriteTimeout,
		IdleTimeout:       o.IdleTimeout,
		MaxHeaderBytes:    o.MaxHeaderBytes,
	}
	srv := &Server{
		o:       o,
		lis:     lis,
		httpSrv: httpSrv,
	}
	return srv
}

func (s *Server) String() string {
	return "management"
}

func (s *Server) Start(ctx context.Context) error {
	err := errors.New("server already started")
	s.startOnce.Do(func() {
		if s.lis == nil {
			err = errors.New("net listener is nil")
			return
		}
		err = nil
		if s.o.TLSConf == nil {
			err = s.httpSrv.Serve(s.lis)
			return
		}
		lis := tls.NewListener(s.lis, s.o.TLSConf)
		err = s.httpSrv.Serve(lis)
	})
	return err
}

func (s *Server) Stop(ctx context.Context) error {
	err := errors.New("server already stopped")
	s.stopOnce.Do(func() {
		err = nil
		err = s.httpSrv.Shutdown(ctx)
	})
	return err
}
