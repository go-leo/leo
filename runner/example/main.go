package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/go-leo/errorx"
	"github.com/shirou/gopsutil/v3/cpu"

	"github.com/go-leo/leo/v2/global"
	"github.com/go-leo/leo/v2/runner"
)

type caller1 struct{}

func (c caller1) String() string {
	return "caller1"
}

func (c caller1) Invoke(ctx context.Context) error {
	global.Logger().Info(c.String() + " invoke enter")
	defer global.Logger().Info(c.String() + " invoke exit")
	return nil
}

type caller2 struct{}

func (c caller2) String() string {
	return "caller2"
}

func (c caller2) Invoke(ctx context.Context) error {
	global.Logger().Info(c.String() + " invoke enter")
	defer global.Logger().Info(c.String() + " invoke exit")
	global.Logger().Info(c.String() + " will sleep 30s")
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(30 * time.Second):
	}
	return nil
}

type runner1 struct {
}

func (r runner1) String() string {
	return "runner1"
}

func (r runner1) Start(ctx context.Context) error {
	global.Logger().Info(r.String() + " Start enter")
	defer global.Logger().Info(r.String() + " Start exit")
	global.Logger().Info(r.String() + " will sleep 30s")
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(30 * time.Second):
	}
	return nil
}

func (r runner1) Stop(ctx context.Context) error {
	global.Logger().Info(r.String() + " Stop enter")
	defer global.Logger().Info(r.String() + " Stop exit")
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(4 * time.Second):
		return nil
	}
}

type runner2 struct {
	w     io.WriteCloser
	exitC chan struct{}
}

func (h *runner2) String() string { return "runner2" }

func (h *runner2) Start(ctx context.Context) error {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-h.exitC:
			ticker.Stop()
			return h.w.Close()
		case t := <-ticker.C:
			percent, err := cpu.Percent(time.Second, true)
			if err != nil {
				return err
			}
			if _, err := fmt.Fprintf(h.w, "%s cpu percent is %f\n", t.String(), percent); err != nil {
				return err
			}
		}
	}
}

func (h *runner2) Stop(ctx context.Context) error {
	close(h.exitC)
	return nil
}

func main() {
	executor := runner.NewExecutor()
	executor.AddCallable(caller1{})
	executor.AddCallable(caller2{})
	executor.AddRunnable(runner1{})
	executor.AddRunnable(&runner2{w: errorx.Quiet(os.Create("/tmp/cpup.log")), exitC: make(chan struct{})})

	ctx := context.Background()
	//ctx, cancelFunc := context.WithTimeout(ctx, 6*time.Second)
	//defer cancelFunc()
	err := executor.Execute(ctx)
	if err != nil {
		panic(err)
	}
}
