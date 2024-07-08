package httpx

import (
	"context"
	"errors"
	kitlog "github.com/go-kit/log"
	"github.com/go-leo/gox/netx/addrx"
	"github.com/go-leo/gox/syncx/brave"
	"github.com/go-leo/leo/v3/healthx"
	"github.com/go-leo/leo/v3/logx"
	"log"
	"net"
	"net/http"
	"strconv"
)

type server struct {
	lis     net.Listener
	httpSrv *http.Server
	opts    *options
	host    string
	port    int
	checker *healthChecker
}

func (s *server) Start(ctx context.Context) error {
	s.checker.Resume()
	defer s.checker.Shutdown()
	return s.serve(ctx)
}

func (s *server) Stop(ctx context.Context) error {
	s.checker.Shutdown()
	return s.httpSrv.Shutdown(ctx)
}

func (s *server) Host() string {
	return s.host
}

func (s *server) Port() int {
	return s.port
}

func (s *server) serve(ctx context.Context) error {
	s.httpSrv.BaseContext = func(listener net.Listener) context.Context {
		return ctx
	}
	errC := brave.GoE(func() error {
		var err error
		if s.httpSrv.TLSConfig != nil {
			err = s.httpSrv.ServeTLS(s.lis, "", "")
		} else {
			err = s.httpSrv.Serve(s.lis)
		}
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	})
	go func() {
		<-ctx.Done()

	}()
	return <-errC
}

func NewServer(handler http.Handler, opts ...Option) (*server, error) {
	o := new(options).apply(opts...).init()
	l, err := net.Listen("tcp", o.Addr)
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
	return newServer(handler, o, l, host, port), nil
}

func newServer(
	handler http.Handler,
	o *options,
	lis net.Listener,
	host string,
	port int,
) *server {
	httpSrv := &http.Server{
		Addr:                         "",
		Handler:                      handler,
		DisableGeneralOptionsHandler: o.DisableGeneralOptionsHandler,
		TLSConfig:                    o.TLSConfig,
		ReadTimeout:                  o.ReadTimeout,
		ReadHeaderTimeout:            o.ReadHeaderTimeout,
		WriteTimeout:                 o.WriteTimeout,
		IdleTimeout:                  o.IdleTimeout,
		MaxHeaderBytes:               o.MaxHeaderBytes,
		TLSNextProto:                 nil,
		ConnState:                    nil,
		ErrorLog:                     log.New(kitlog.NewStdlibAdapter(logx.L()), "", 0),
		BaseContext:                  nil,
		ConnContext:                  nil,
	}
	checker := &healthChecker{}
	s := &server{httpSrv: httpSrv, checker: checker, opts: o, lis: lis, host: host, port: port}
	healthx.RegisterChecker(checker)
	return s
}
