package httpx

import (
	"context"
	"github.com/go-leo/leo/v3/healthx"
	"sync/atomic"
)

type healthChecker struct {
	serving atomic.Bool
}

func (c *healthChecker) Check(ctx context.Context) healthx.Status {
	if c.IsServing() {
		return healthx.Serving
	}
	return healthx.NotServing
}

func (c *healthChecker) Name() string {
	return "http"
}

func (c *healthChecker) IsServing() bool {
	return c.serving.Load()
}

func (c *healthChecker) Shutdown() {
	c.serving.Store(false)
}

func (c *healthChecker) Resume() {
	c.serving.Store(true)
}
