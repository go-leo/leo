package management

import (
	"context"
	"crypto/tls"
	"errors"
	"net"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/go-leo/leo/v2/cron"
	leogrpc "github.com/go-leo/leo/v2/grpc"
	leohttp "github.com/go-leo/leo/v2/http"
	"github.com/go-leo/leo/v2/management/router/app"
	"github.com/go-leo/leo/v2/management/router/config"
	mgmtcron "github.com/go-leo/leo/v2/management/router/cron"
	"github.com/go-leo/leo/v2/management/router/env"
	mgmtgrpc "github.com/go-leo/leo/v2/management/router/grpc"
	mgmthttp "github.com/go-leo/leo/v2/management/router/http"
	"github.com/go-leo/leo/v2/management/router/metric"
	"github.com/go-leo/leo/v2/management/router/profile"
	"github.com/go-leo/leo/v2/management/router/restart"
	"github.com/go-leo/leo/v2/management/router/shutdown"
	"github.com/go-leo/leo/v2/management/router/system"
	"github.com/go-leo/leo/v2/pubsub"
)

type Router struct {
	HttpMethod   string
	Path         string
	HandlerFuncs []gin.HandlerFunc
}

type options struct {
	GinMiddlewares      []gin.HandlerFunc
	Routers             []Router
	TLSConf             *tls.Config
	ReadTimeout         time.Duration
	WriteTimeout        time.Duration
	IdleTimeout         time.Duration
	MaxHeaderBytes      int
	ExitSignals         []os.Signal
	RestartSignals      []os.Signal
	CronTask            *cron.Task
	PubSubTask          *pubsub.Task
	HTTPSrv             *leohttp.Server
	HTTPHealthCheckOpts *mgmthttp.HealthCheckOptions
	GRPCSrv             *leogrpc.Server
	GRPCHealthCheckOpts *mgmtgrpc.HealthCheckOptions
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

func HTTP(srv *leohttp.Server, HttpHealthCheckOpts *mgmthttp.HealthCheckOptions) Option {
	return func(o *options) {
		o.HTTPSrv = srv
		o.HTTPHealthCheckOpts = HttpHealthCheckOpts
	}
}

func GRPC(srv *leogrpc.Server, grpcHealthCheckOpts *mgmtgrpc.HealthCheckOptions) Option {
	return func(o *options) {
		o.GRPCSrv = srv
		o.GRPCHealthCheckOpts = grpcHealthCheckOpts

	}
}

func Cron(task *cron.Task) Option {
	return func(o *options) {
		o.CronTask = task
	}
}

func PubSub(task *pubsub.Task) Option {
	return func(o *options) {
		o.PubSubTask = task
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

func NewServer(port int, opts ...Option) (*Server, error) {
	if port <= 0 {
		return nil, errors.New("management port is zero")
	}
	// 监听端口
	lis, err := net.Listen("tcp", net.JoinHostPort("", strconv.Itoa(port)))
	if err != nil {
		return nil, err
	}

	o := new(options)
	o.apply(opts)
	o.init()

	gin.SetMode(gin.ReleaseMode)
	mux := gin.New()
	mux.Use(gin.Recovery())
	mux.Use(o.GinMiddlewares...)
	rg := mux.Group("/management")
	// register http server
	mgmthttp.Route(rg, o.HTTPSrv, o.HTTPHealthCheckOpts)
	// register grpc server
	mgmtgrpc.Route(rg, o.GRPCSrv, o.GRPCHealthCheckOpts)
	// register cron task
	mgmtcron.Route(rg, o.CronTask)
	// register profile
	profile.Route(rg)
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
	return srv, nil
}

func (s *Server) String() string {
	return "management"
}

func (s *Server) Start(ctx context.Context) error {
	if s.lis == nil {
		return errors.New("net listener is nil")
	}
	err := errors.New("server already started")
	s.startOnce.Do(func() {
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
