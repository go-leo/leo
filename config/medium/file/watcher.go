package file

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"sync/atomic"

	"github.com/fsnotify/fsnotify"

	"github.com/go-leo/leo/config"
	"github.com/go-leo/leo/log"
)

const (
	stop uint64 = iota
	start
)

var _ config.Watcher = new(Watcher)

type Watcher struct {
	filename string
	state    uint64
	stopCtx  context.Context
	stopFunc context.CancelFunc
	log      log.Logger
}

func (watcher *Watcher) Start(ctx context.Context) (<-chan *config.Event, error) {
	if !atomic.CompareAndSwapUint64(&watcher.state, stop, start) {
		return nil, errors.New("watcher is started")
	}
	fsWatcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	if err = fsWatcher.Add(watcher.filename); err != nil {
		return nil, err
	}
	watcher.stopCtx, watcher.stopFunc = context.WithCancel(ctx)
	eventC := make(chan *config.Event)
	go watcher.watch(fsWatcher, eventC)
	return eventC, nil
}

func (watcher *Watcher) watch(fsWatcher *fsnotify.Watcher, eventC chan *config.Event) {
	defer close(eventC)
	defer func() {
		if err := fsWatcher.Close(); err != nil {
			eventC <- config.NewErrEvent(err)
		}
	}()
	for {
		select {
		case <-watcher.stopCtx.Done():
			watcher.log.Info("stop watch file:", watcher.filename)
			return
		case event, ok := <-fsWatcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				data, err := ioutil.ReadFile(watcher.filename)
				if err != nil {
					watcher.log.Error(fmt.Errorf("failed read file %s, %w", watcher.filename, err))
					return
				}
				eventC <- config.NewContentEvent(event, data)
			}
		case err, ok := <-fsWatcher.Errors:
			if !ok {
				return
			}
			eventC <- config.NewErrEvent(err)
		}
	}
}

func (watcher *Watcher) Stop(_ context.Context) error {
	if !atomic.CompareAndSwapUint64(&watcher.state, start, stop) {
		return errors.New("watcher is stopped")
	}
	watcher.stopFunc()
	return nil
}

type WatcherOption func(watcher *Watcher)

func WithLogger(log log.Logger) LoaderOption {
	return func(loader *Loader) {
		loader.log = log
	}
}

func NewWatcher(filename string, opts ...WatcherOption) *Watcher {
	w := &Watcher{
		filename: filename,
		state:    stop,
		stopCtx:  nil,
		stopFunc: nil,
		log:      log.Discard{},
	}
	for _, opt := range opts {
		opt(w)
	}
	return w
}
