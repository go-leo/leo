package text

import (
	"context"
	"sync"
	"time"

	"github.com/go-leo/gox/slicex"

	"codeup.aliyun.com/qimao/leo/leo/config"
	"codeup.aliyun.com/qimao/leo/leo/log"
)

var _ config.Resource = new(Resource)
var _ config.Watcher = new(watcher)

type options struct {
	Logger log.Logger
}

type Option func(o *options)

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *options) init() {
}

func Logger(log log.Logger) Option {
	return func(o *options) {
		o.Logger = log
	}
}

type Resource struct {
	options   *options
	text      *[]byte
	name      string
	extension string
}

func (r *Resource) Load(ctx context.Context) (*config.Source, error) {
	return r.loadSource(ctx)
}

func (r *Resource) Watch(ctx context.Context) (config.Watcher, error) {
	source, _ := r.loadSource(ctx)
	w := &watcher{resource: r, source: source}
	return w, w.init(ctx)
}

func (r *Resource) loadSource(ctx context.Context) (*config.Source, error) {
	return &config.Source{
		Name:      r.name,
		Value:     *r.text,
		Extension: r.extension,
	}, nil
}

type watcher struct {
	resource *Resource
	source   *config.Source
	eventCs  []chan<- config.Event
	closeC   chan struct{}
	mutex    sync.Mutex
}

func (watcher *watcher) Notify(eventC chan<- config.Event) {
	watcher.mutex.Lock()
	defer watcher.mutex.Unlock()
	watcher.eventCs = slicex.AppendIfNotContains(watcher.eventCs, eventC)
}

func (watcher *watcher) StopNotify(eventC chan<- config.Event) {
	watcher.mutex.Lock()
	defer watcher.mutex.Unlock()
	watcher.eventCs = slicex.Remove(watcher.eventCs, eventC)
}

func (watcher *watcher) Close(ctx context.Context) error {
	watcher.closeC <- struct{}{}
	watcher.mutex.Lock()
	watcher.eventCs = nil
	watcher.mutex.Unlock()
	return nil
}

func (watcher *watcher) init(ctx context.Context) error {
	watcher.closeC = make(chan struct{})
	watcher.watch()
	return nil
}

func (watcher *watcher) watch() {
	go func() {
		for {
			select {
			case <-watcher.closeC:
				return
			case <-time.After(time.Second):
				source, err := watcher.resource.loadSource(context.Background())
				if err != nil {
					watcher.sendError(err)
					continue
				}
				if string(source.Value) == string(watcher.source.Value) {
					continue
				}
				watcher.source = source
				for _, eventC := range watcher.eventCs {
					eventC <- config.SourceEvent(source)
				}
			}
		}
	}()
}

func (watcher *watcher) sendError(err error) {
	if err == nil {
		return
	}
	for _, eventC := range watcher.eventCs {
		watcher.resource.options.Logger.Debug("sending error event")
		eventC <- config.ErrorEvent(err)
	}
}

func NewResource(text *[]byte, name string, extension string, opts ...Option) *Resource {
	o := &options{
		Logger: log.Discard{},
	}
	o.apply(opts...)
	o.init()
	resource := &Resource{
		options:   o,
		text:      text,
		name:      name,
		extension: extension,
	}
	return resource
}
