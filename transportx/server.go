package transportx

import (
	"context"
	"github.com/go-leo/gox/netx/addrx"
	"github.com/go-leo/leo/v3/runner"
	"net"
	"strconv"
	"time"
)

type Server interface {
	runner.StartStopper
}

type ServerFactory func(lis net.Listener, args any) Server

type options struct {
	Addr            string
	ShutdownTimeout time.Duration
}

type Option func(o *options)

func (o *options) apply(opts ...Option) *options {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func (o *options) init() *options {
	if o.Addr == "" {
		o.Addr = ":0"
	}
	return o
}

// Addr set server addr.
func Addr(addr string) Option {
	return func(o *options) {
		o.Addr = addr
	}
}

func ShutdownTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.ShutdownTimeout = timeout
	}
}

type server struct {
	lis  net.Listener
	opts *options
	host string
	port int
	srv  Server
}

func (s *server) Start(ctx context.Context) error {
	return nil
}

func (s *server) Stop(ctx context.Context) error {
	return nil
}

func (s *server) Host() string {
	return s.host
}

func (s *server) Port() int {
	return s.port
}

func NewServer(factory ServerFactory, opts ...Option) (Server, error) {
	o := new(options).apply(opts...).init()
	lis, err := net.Listen("tcp", o.Addr)
	if err != nil {
		return nil, err
	}
	host, p, err2 := funcName(lis.Addr())
	if err2 != nil {
		return nil, err2
	}
	port, err := strconv.Atoi(p)
	if err != nil {
		return nil, err
	}
	srv := factory(lis, nil)
	return newServer(srv, o, host, port), nil
}

func funcName(address net.Addr) (string, string, error) {
	host, port, err := net.SplitHostPort(address.String())
	if err != nil {
		return "", "", err
	}
	if addrx.IsGlobalUnicastIP(net.ParseIP(host)) {
		return host, port, nil
	}
	ip, err := addrx.GlobalUnicastIPString()
	if err != nil {
		return "", "", err
	}
	return ip, port, nil
}

func newServer(srv Server, o *options, host string, port int) *server {
	s := &server{
		srv:  srv,
		opts: o,
		host: host,
		port: port,
	}
	return s
}
