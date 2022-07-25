package config

import (
	"context"
	"fmt"
)

// Loader load data from some medium.
type Loader interface {
	// ContentType get config content type
	ContentType() string
	// Load load config raw data from some where, like local file or remote server
	Load() error
	// RawData get raw data
	RawData() []byte
}

// Watcher monitors whether the data source has changed and, if so, notifies the changed event
type Watcher interface {
	// Start watch
	Start(ctx context.Context) (<-chan *Event, error)
	// Stop watch
	Stop(ctx context.Context) error
}

type EventKey interface {
	KeyName() string
}

// Event
type Event struct {
	// 错误
	err         error
	description fmt.Stringer
	data        []byte
}

func (e *Event) Data() []byte {
	return e.data
}

func (e *Event) String() string {
	return e.description.String() + "/n" + string(e.data)
}

func (e *Event) Err() error {
	return e.err
}

func NewErrEvent(err error) *Event {
	return &Event{err: err}
}

func NewContentEvent(description fmt.Stringer, data []byte) *Event {
	return &Event{description: description, data: data}
}
