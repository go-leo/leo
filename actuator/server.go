package actuator

import (
	"context"
	"crypto/tls"
	"errors"
	"net"
	"net/http"
	"path"
	"strconv"
)

type Server struct {
	options *options
	port    int
}

func New(port int, opts ...Option) *Server {
	o := new(options)
	o.apply(opts...)
	o.init()
	return &Server{options: o, port: port}
}

func (server *Server) Run(ctx context.Context) error {
	// listen port
	address := net.JoinHostPort("", strconv.Itoa(server.port))
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	if server.options.TLSConf != nil {
		lis = tls.NewListener(lis, server.options.TLSConf)
	}

	// handle http
	mux := http.NewServeMux()
	for _, h := range server.options.Handlers {
		mux.HandleFunc(path.Join(server.options.PathPrefix, h.Pattern()), h.Handle)
	}

	// new go std http server
	httpSrv := &http.Server{
		Handler:        mux,
		ReadTimeout:    server.options.ReadTimeout,
		WriteTimeout:   server.options.WriteTimeout,
		IdleTimeout:    server.options.IdleTimeout,
		MaxHeaderBytes: server.options.MaxHeaderBytes,
	}

	// async run http serve
	serveErrC := make(chan error)
	go func(errC chan<- error, lis net.Listener) {
		defer close(errC)
		err := httpSrv.Serve(lis)
		if errors.Is(err, http.ErrServerClosed) {
			return
		}
		if err != nil {
			serveErrC <- err
		}
	}(serveErrC, lis)

	// wait until context canceled or failed to serve
	select {
	case serveErr := <-serveErrC:
		// failed to serve. return serve error
		return serveErr
	case <-ctx.Done():
		// context canceled
		// shutdown server.
		ctx := context.Background()
		var cancelFunc = func() {}
		if server.options.CloseTimeout > 0 {
			ctx, cancelFunc = context.WithTimeout(ctx, server.options.CloseTimeout)
		}
		defer cancelFunc()
		shutdownErr := httpSrv.Shutdown(ctx)
		// wait Serve shutdown
		serveErr := <-serveErrC
		// return shutdown and serve error if has error
		return errors.Join(shutdownErr, serveErr)
	}
}
