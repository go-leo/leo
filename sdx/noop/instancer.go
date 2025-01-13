package noop

import (
	"github.com/go-kit/kit/sd"
)

// Instancer is an empty implementation of the sd.Instancer interface.
type Instancer struct{}

// Register implements Instancer.
func (Instancer) Register(ch chan<- sd.Event) {}

// Deregister implements Instancer.
func (Instancer) Deregister(ch chan<- sd.Event) {}

// Stop implements Instancer.
func (i Instancer) Stop() {}
