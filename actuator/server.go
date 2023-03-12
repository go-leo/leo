package actuator

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"path"
	"strconv"

	"github.com/go-leo/gox/syncx/brave"
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
	serveErrC := brave.GoE(
		func() error {
			if server.options.TLSConf != nil {
				return httpSrv.Serve(tls.NewListener(lis, server.options.TLSConf))
			}
			return httpSrv.Serve(lis)
		},
		func(p any) error { return fmt.Errorf("panic triggered: %+v", p) },
	)

	// wait until context canceled or failed to serve
	select {
	case <-ctx.Done():
		// context canceled, shutdown server.
		ctx := context.Background()
		var cancelFunc = func() {}
		if server.options.CloseTimeout > 0 {
			ctx, cancelFunc = context.WithTimeout(ctx, server.options.CloseTimeout)
		}
		defer cancelFunc()
		return httpSrv.Shutdown(ctx)
	case err := <-serveErrC:
		// failed to serve.
		return err
	}
}
