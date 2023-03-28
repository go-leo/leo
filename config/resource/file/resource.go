package file

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/go-leo/gox/pathx/filepathx"
	"github.com/go-leo/gox/slicex"

	"codeup.aliyun.com/qimao/leo/leo/config"
	"codeup.aliyun.com/qimao/leo/leo/log"
)

var _ config.Resource = new(Resource)

var _ config.Watcher = new(Watcher)

type options struct {
	Logger    log.Logger
	Extension string
}

type Option func(o *options)

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *options) init() {

}

func Extension(Extension string) Option {
	return func(loader *options) {
		loader.Extension = Extension
	}
}

func Logger(log log.Logger) Option {
	return func(loader *options) {
		loader.Logger = log
	}
}

type Resource struct {
	options  *options
	filename string
}

func (r *Resource) Load(ctx context.Context) (*config.Source, error) {
	return r.loadSource()
}

func (r *Resource) Watch(ctx context.Context) (config.Watcher, error) {
	w := &Watcher{resource: r}
	return w, w.init(ctx)
}

func (r *Resource) loadSource() (*config.Source, error) {
	value, err := os.ReadFile(r.filename)
	if err != nil {
		return nil, err
	}
	return &config.Source{
		Name:      filepath.Base(r.filename),
		Value:     value,
		Extension: r.options.Extension,
	}, nil
}

type Watcher struct {
	resource  *Resource
	fsWatcher *fsnotify.Watcher
	eventCs   []chan<- config.Event
	closeC    chan struct{}
	mutex     sync.Mutex
}

func (watcher *Watcher) Notify(eventC chan<- config.Event) {
	watcher.mutex.Lock()
	defer watcher.mutex.Unlock()
	watcher.eventCs = slicex.AppendIfNotContains(watcher.eventCs, eventC)
}

func (watcher *Watcher) StopNotify(eventC chan<- config.Event) {
	watcher.mutex.Lock()
	defer watcher.mutex.Unlock()
	watcher.eventCs = slicex.Remove(watcher.eventCs, eventC)
}

func (watcher *Watcher) Close(ctx context.Context) error {
	err := watcher.fsWatcher.Close()
	watcher.closeC <- struct{}{}
	watcher.mutex.Lock()
	watcher.eventCs = nil
	watcher.mutex.Unlock()
	return err
}

func (watcher *Watcher) init(ctx context.Context) error {
	fsWatcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	watcher.fsWatcher = fsWatcher
	filename := watcher.resource.filename
	st, err := os.Lstat(filename)
	if err != nil {
		return err
	}

	if st.IsDir() {
		return fmt.Errorf("%q is a directory, not a file", filename)
	}
	if err := watcher.fsWatcher.Add(filepath.Dir(filename)); err != nil {
		return err
	}
	watcher.closeC = make(chan struct{})
	watcher.watch()
	return nil
}

func (watcher *Watcher) watch() {
	go func() {
		for {
			select {
			case <-watcher.closeC:
				return
			case event, ok := <-watcher.fsWatcher.Events:
				if !ok {
					return
				}
				watcher.handleFileChangeEvent(event)
			case err, ok := <-watcher.fsWatcher.Errors:
				if !ok {
					return
				}
				watcher.sendError(err)
			}
		}
	}()
}

func (watcher *Watcher) sendError(err error) {
	if err == nil {
		return
	}
	for _, eventC := range watcher.eventCs {
		watcher.resource.options.Logger.Debug("sending error event")
		eventC <- config.ErrorEvent(err)
	}
}

func (watcher *Watcher) handleFileChangeEvent(event fsnotify.Event) {
	// Ignore files we're not interested in.
	filename := watcher.resource.filename
	if filename != event.Name {
		return
	}
	source, err := watcher.resource.loadSource()
	if err != nil {
		watcher.sendError(err)
		return
	}
	for _, eventC := range watcher.eventCs {
		eventC <- config.SourceEvent(source)
	}
}

func NewResource(filename string, opts ...Option) *Resource {
	filename = filepath.Clean(filename)
	o := &options{
		Logger:    log.Discard{},
		Extension: filepathx.Extension(filename),
	}
	o.apply(opts...)
	o.init()
	for _, opt := range opts {
		opt(o)
	}
	resource := &Resource{
		options:  o,
		filename: filename,
	}
	return resource
}
