package text

import (
	"context"
	"path/filepath"
	"sync"
	"time"

	"codeup.aliyun.com/qimao/leo/leo/config"
	"codeup.aliyun.com/qimao/leo/leo/internal/gox/slicex"
)

var _ config.Resource = new(Resource)
var _ config.Watcher = new(watcher)

type Resource struct {
	text      *Text
	name      string
	extension string
}

func (r *Resource) Load(ctx context.Context) (*config.Source, error) {
	return r.loadSource()
}

func (r *Resource) Watch(ctx context.Context) (config.Watcher, error) {
	source, err := r.loadSource()
	if err != nil {
		return nil, err
	}
	w := &watcher{resource: r, source: source}
	return w, w.init(ctx)
}

func (r *Resource) loadSource() (*config.Source, error) {
	return config.NewSource(filepath.Base(r.name), []byte(r.text.Get()), r.extension), nil
}

type watcher struct {
	resource *Resource
	source   *config.Source
	eventCs  []chan<- config.Event
	closeC   chan struct{}
	mutex    sync.RWMutex
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

func (watcher *watcher) eventChannels() []chan<- config.Event {
	watcher.mutex.RLock()
	defer watcher.mutex.RUnlock()
	return watcher.eventCs
}

func (watcher *watcher) init(ctx context.Context) error {
	watcher.closeC = make(chan struct{})
	watcher.watch()
	watcher.resource.text.AddObserver(ObserverFunc(func(newText, oldText string) {
		if newText == oldText {
			return
		}
		source := config.NewSource(filepath.Base(watcher.resource.name), []byte(newText), watcher.resource.extension)
		watcher.source = source
		for _, eventC := range watcher.eventChannels() {
			eventC <- config.SourceEvent(source)
		}
	}))
	return nil
}

func (watcher *watcher) watch() {
	go func() {
		for {
			select {
			case <-watcher.closeC:
				return
			case <-time.After(time.Second):
				source, err := watcher.resource.loadSource()
				if err != nil {
					watcher.sendError(err)
					continue
				}
				if string(source.Value()) == string(watcher.source.Value()) {
					continue
				}
				watcher.source = source
				for _, eventC := range watcher.eventChannels() {
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
	for _, eventC := range watcher.eventChannels() {
		eventC <- config.ErrorEvent(err)
	}
}

func NewResource(text *Text, name string, ext string) *Resource {
	resource := &Resource{
		text:      text,
		name:      name,
		extension: ext,
	}
	return resource
}
