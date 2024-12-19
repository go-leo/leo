package passthroughx

import (
	"github.com/go-kit/kit/sd"
)

type Instancer struct {
	Instance string
}

func (s Instancer) Register(ch chan<- sd.Event) {
	ch <- sd.Event{Instances: []string{s.Instance}}
}

func (s Instancer) Deregister(ch chan<- sd.Event) {}

func (s Instancer) Stop() {}
