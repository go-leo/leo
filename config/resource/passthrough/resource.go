package passthrough

import (
	"context"
	"sync"

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
	options *options
	source  *config.Source
}

func (r *Resource) Load(ctx context.Context) (*config.Source, error) {
	return r.source, nil
}

func (r *Resource) Watch(ctx context.Context) (config.Watcher, error) {
	w := &watcher{resource: r}
	return w, w.init(ctx)
}

type watcher struct {
	resource *Resource
	eventCs  []chan<- config.Event
	mutex    sync.Mutex
}

func (watcher *watcher) Update(source *config.Source) {
	for _, eventC := range watcher.eventCs {
		eventC <- config.SourceEvent(source)
	}
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
	watcher.mutex.Lock()
	watcher.resource.source.DeleteObserver(watcher)
	watcher.eventCs = nil
	watcher.mutex.Unlock()
	return nil
}

func (watcher *watcher) init(ctx context.Context) error {
	watcher.resource.source.AddObserver(watcher)
	return nil
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

func NewResource(source *config.Source, opts ...Option) *Resource {
	o := &options{
		Logger: log.L(),
	}
	o.apply(opts...)
	o.init()
	resource := &Resource{
		options: o,
		source:  source,
	}
	return resource
}
