package actuator

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"strconv"

	"codeup.aliyun.com/qimao/leo/leo/actuator/internal/handler"
)

type Server struct {
	o       *options
	port    int
	httpSrv *http.Server
}

func New(port int, opts ...Option) *Server {
	o := new(options)
	o.apply(opts...)
	o.init()
	return &Server{o: o, port: port}
}

func (s *Server) Start(ctx context.Context) error {
	address := net.JoinHostPort("", strconv.Itoa(s.port))
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	o := s.o
	mux := http.NewServeMux()
	var handlers []Handler

	if o.ConsoleCommander != nil {
		handlers = append(handlers, &handler.ConsoleCommanderHandler{ConsoleCommander: o.ConsoleCommander})
	}

	if o.ResourceServer != nil {
		handlers = append(handlers, &handler.ResourceServerHandler{ResourceServer: o.ResourceServer})
	}

	if o.RPCProvider != nil {
		handlers = append(handlers, &handler.RPCProviderHandler{RPCProvider: o.RPCProvider})
	}

	if o.Scheduler != nil {
		handlers = append(handlers, &handler.SchedulerHandler{Scheduler: o.Scheduler})
	}

	for o.SteamRouter != nil {
		handlers = append(handlers, &handler.SteamRouterHandler{SteamRouter: o.SteamRouter})
	}

	if o.ViewController != nil {
		handlers = append(handlers, &handler.ViewControllerHandler{ViewController: o.ViewController})
	}

	for _, healthChecker := range o.HealthCheckers {
		handlers = append(handlers, &handler.HealthCheckerHandler{HealthChecker: healthChecker})
	}

	if o.PProfEnabled {
		handlers = append(handlers, &handler.PProfIndexHandler{})
		handlers = append(handlers, &handler.PProfCmdlineHandler{})
		handlers = append(handlers, &handler.PProfProfileHandler{})
		handlers = append(handlers, &handler.PProfSymbolHandler{})
		handlers = append(handlers, &handler.PProfTraceHandler{})
	}

	for _, h := range o.Handlers {
		handlers = append(handlers, h)
	}

	for _, h := range handlers {
		mux.HandleFunc(h.Pattern(), h.Handle)
	}

	s.httpSrv = &http.Server{Handler: mux}

	if s.o.TLSConf != nil {
		return s.httpSrv.Serve(tls.NewListener(lis, s.o.TLSConf))
	}
	return s.httpSrv.Serve(lis)
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpSrv.Shutdown(ctx)
}
