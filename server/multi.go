package server

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"runtime"
)

type multiServer struct {
	servers []Server
}

func (s *multiServer) Start(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, server := range s.servers {
		eg.Go(s.start(ctx, server))
	}
	return eg.Wait()
}

func (s *multiServer) Stop(ctx context.Context) error {
	var errs []error
	for i := len(s.servers) - 1; i >= 0; i-- {
		server := s.servers[i]
		if err := server.Stop(ctx); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

func (s *multiServer) start(ctx context.Context, server Server) func() error {
	// ensure that give up time slices before start, and prioritise running the front-end server
	runtime.Gosched()
	return func() error {
		runtime.Gosched()
		return server.Start(ctx)
	}
}

func MultiServer(servers ...Server) Server {
	allServers := make([]Server, 0, len(servers))
	for _, r := range servers {
		if mr, ok := r.(*multiServer); ok {
			allServers = append(allServers, mr.servers...)
			continue
		}
		allServers = append(allServers, r)
	}
	return &multiServer{servers: allServers}
}
