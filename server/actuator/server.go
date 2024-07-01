package actuator

import (
	"context"
	"errors"
	kitlog "github.com/go-kit/log"
	"github.com/go-leo/leo/v3/logx"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
	"strconv"
)

type Server struct {
	httpSrv *http.Server
	opts    *options
	lis     net.Listener
}

func (s *Server) Start(ctx context.Context) error {
	err := s.serve(ctx)
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}

func (s *Server) Stop(ctx context.Context) error {
	if s.opts.ShutdownTimeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, s.opts.ShutdownTimeout)
		defer cancel()
	}
	return s.httpSrv.Shutdown(ctx)
}

func (s *Server) serve(ctx context.Context) error {
	s.httpSrv.Handler = s.newHandler()
	s.httpSrv.BaseContext = func(listener net.Listener) context.Context { return ctx }
	if s.httpSrv.TLSConfig != nil {
		return s.httpSrv.ServeTLS(s.lis, "", "")
	} else {
		return s.httpSrv.Serve(s.lis)
	}
}

func (s *Server) newHandler() *mux.Router {
	router := mux.NewRouter()
	router.Use(s.opts.Middlewares...)
	for _, h := range getHandlers() {
		router.PathPrefix(s.opts.PathPrefix).Path(h.Pattern()).HandlerFunc(h.ServeHTTP)
	}
	for _, h := range s.opts.Handlers {
		router.PathPrefix(s.opts.PathPrefix).Path(h.Pattern()).HandlerFunc(h.ServeHTTP)
	}
	return router
}

func NewServer(port int, opts ...Option) (*Server, error) {
	address := net.JoinHostPort("", strconv.Itoa(port))
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	o := new(options).apply(opts...).init()
	return newServer(lis, o), nil
}

func newServer(lis net.Listener, options *options) *Server {
	httpSrv := &http.Server{
		Addr:                         "",
		Handler:                      nil,
		DisableGeneralOptionsHandler: options.DisableGeneralOptionsHandler,
		TLSConfig:                    options.TLSConfig,
		ReadTimeout:                  options.ReadTimeout,
		ReadHeaderTimeout:            options.ReadHeaderTimeout,
		WriteTimeout:                 options.WriteTimeout,
		IdleTimeout:                  options.IdleTimeout,
		MaxHeaderBytes:               options.MaxHeaderBytes,
		TLSNextProto:                 nil,
		ConnState:                    nil,
		ErrorLog:                     log.New(kitlog.NewStdlibAdapter(logx.L()), "", 0),
		BaseContext:                  nil,
		ConnContext:                  nil,
	}
	return &Server{httpSrv: httpSrv, opts: options, lis: lis}
}
