package httpx

import (
	"context"
	"github.com/go-leo/leo/v3/healthx"
	"sync/atomic"
)

type Checker struct {
	serving atomic.Bool
}

func (c *Checker) Check(ctx context.Context) healthx.Status {
	if c.IsServing() {
		return healthx.Serving
	}
	return healthx.NotServing
}

func (c *Checker) Name() string {
	return "http"
}

func (c *Checker) IsServing() bool {
	return c.serving.Load()
}

func (c *Checker) Shutdown() {
	c.serving.Store(false)
}

func (c *Checker) Resume() {
	c.serving.Store(true)
}
