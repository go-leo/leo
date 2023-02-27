package config

import (
	"context"
)

// Watcher monitors whether the data source has changed and, if so, notifies the changed event
type Watcher interface {
	StartWatch(ctx context.Context) (<-chan Event, error)
	StopWatch(ctx context.Context) error
}

type noopWatcher struct{}

func (noopWatcher) StartWatch(ctx context.Context) (<-chan Event, error) {
	eventC := make(chan Event)
	defer close(eventC)
	return eventC, nil
}

func (noopWatcher) StopWatch(ctx context.Context) error {
	return nil
}
