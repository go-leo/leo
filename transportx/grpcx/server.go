package grpcx

import (
	"context"
	"github.com/go-leo/leo/v3/runner"
	"google.golang.org/grpc"
	"net"
)

var _ runner.StartStopper = (*server)(nil)

type server struct {
	*grpc.Server
	lis net.Listener
}

func (s *server) Start(ctx context.Context) error {
	return s.Server.Serve(s.lis)
}

func (s *server) Stop(ctx context.Context) error {
	s.Server.GracefulStop()
	return nil
}
