package httpx

import (
	"context"
	"errors"
	kitlog "github.com/go-kit/log"
	"github.com/go-leo/gox/netx/addrx"
	"github.com/go-leo/leo/v3/healthx"
	"github.com/go-leo/leo/v3/logx"
	"log"
	"net"
	"net/http"
	"strconv"
)

type Server struct {
	lis     net.Listener
	httpSrv *http.Server
	opts    *options
	host    string
	port    int
	checker *Checker
}

func (s *Server) Start(ctx context.Context) error {
	err := s.serve(ctx)
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}

func (s *Server) Stop(ctx context.Context) error {
	s.checker.Shutdown()
	return s.httpSrv.Shutdown(ctx)
}

func (s *Server) Host() string {
	return s.host
}

func (s *Server) Port() int {
	return s.port
}

func (s *Server) serve(ctx context.Context) error {
	s.httpSrv.BaseContext = func(listener net.Listener) context.Context {
		return ctx
	}
	s.checker.Resume()
	defer s.checker.Shutdown()
	if s.httpSrv.TLSConfig != nil {
		return s.httpSrv.ServeTLS(s.lis, "", "")
	} else {
		return s.httpSrv.Serve(s.lis)
	}
}

func NewServer(handler http.Handler, opts ...Option) (*Server, error) {
	options := new(options).apply(opts...).init()
	l, err := net.Listen("tcp", options.Addr)
	if err != nil {
		return nil, err
	}
	host, p, err := net.SplitHostPort(l.Addr().String())
	if err != nil {
		return nil, err
	}
	if !addrx.IsGlobalUnicastIP(net.ParseIP(host)) {
		ip, err := addrx.GlobalUnicastIPString()
		if err != nil {
			return nil, err
		}
		host = ip
	}
	port, err := strconv.Atoi(p)
	if err != nil {
		return nil, err
	}
	return newServer(handler, options, l, host, port), nil
}

func newServer(handler http.Handler, opts *options, lis net.Listener, host string, port int) *Server {
	httpSrv := &http.Server{
		Addr:                         "",
		Handler:                      handler,
		DisableGeneralOptionsHandler: opts.DisableGeneralOptionsHandler,
		TLSConfig:                    opts.TLSConfig,
		ReadTimeout:                  opts.ReadTimeout,
		ReadHeaderTimeout:            opts.ReadHeaderTimeout,
		WriteTimeout:                 opts.WriteTimeout,
		IdleTimeout:                  opts.IdleTimeout,
		MaxHeaderBytes:               opts.MaxHeaderBytes,
		TLSNextProto:                 nil,
		ConnState:                    nil,
		ErrorLog:                     log.New(kitlog.NewStdlibAdapter(logx.L()), "", 0),
		BaseContext:                  nil,
		ConnContext:                  nil,
	}
	checker := &Checker{}
	s := &Server{httpSrv: httpSrv, checker: checker, opts: opts, lis: lis, host: host, port: port}
	healthx.RegisterChecker(checker)
	return s
}
