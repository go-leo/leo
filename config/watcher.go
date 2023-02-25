package config

import (
	"context"
)

// Watcher monitors whether the data source has changed and, if so, notifies the changed event
type Watcher interface {
	StartWatch(ctx context.Context) (<-chan Event, error)
	StopWatch(ctx context.Context) error
}
