package config

import (
	"context"
)

// Resource is a loader that can be used to load source config.
type Resource interface {
	Load(ctx context.Context) (*Source, error)
	Watch(ctx context.Context) (Watcher, error)
}

// Watcher monitors whether the data source has changed and, if so, notifies the changed event
type Watcher interface {
	Notify(eventC chan<- Event)
	StopNotify(eventC chan<- Event)
	Close(ctx context.Context) error
}

// Event is a event
type Event interface {
	Data() []byte
	Err() error
	Description() string
}

func DataEvent(data []byte, desc string) Event {
	return &event{data: data, desc: desc}
}

func ErrorEvent(err error, desc string) Event {
	return &event{err: err, desc: desc}
}

type event struct {
	data []byte
	err  error
	desc string
}

func (e event) Data() []byte {
	return e.data
}

func (e event) Err() error {
	return e.err
}

func (e event) Description() string {
	return e.desc
}
