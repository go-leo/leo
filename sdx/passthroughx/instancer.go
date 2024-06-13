package passthroughx

import (
	"github.com/go-kit/kit/sd"
)

type Instancer struct {
	instance string
}

func NewInstancer(instance string) *Instancer {
	s := &Instancer{instance: instance}
	return s
}

func (s *Instancer) Register(ch chan<- sd.Event) {
	ch <- sd.Event{Instances: []string{s.instance}}
}

func (s *Instancer) Deregister(ch chan<- sd.Event) {}

func (s *Instancer) Stop() {}
