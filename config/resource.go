package config

import (
	"context"
	"errors"
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

type multiWatcher struct {
	watchers []Watcher
}

func (w *multiWatcher) Notify(eventC chan<- Event) {
	for _, watcher := range w.watchers {
		watcher.Notify(eventC)
	}
}

func (w *multiWatcher) StopNotify(eventC chan<- Event) {
	for _, watcher := range w.watchers {
		watcher.StopNotify(eventC)
	}
}

func (w *multiWatcher) Close(ctx context.Context) error {
	var errs []error
	for _, watcher := range w.watchers {
		errs = append(errs, watcher.Close(ctx))
	}
	return errors.Join(errs...)
}

// Event is a event
type Event interface {
	Get() (*Source, error)
}

type event struct {
	source *Source
	err    error
}

func (e event) Get() (*Source, error) {
	return e.source, e.err
}

func SourceEvent(source *Source) Event {
	return &event{source: source}
}

func ErrorEvent(err error) Event {
	return &event{err: err}
}

func MultiWatcher(watchers ...Watcher) Watcher {
	w := make([]Watcher, len(watchers))
	copy(w, watchers)
	return &multiWatcher{watchers: w}
}
