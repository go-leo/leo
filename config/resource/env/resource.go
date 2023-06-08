package env

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/pathx/filepathx"
	"codeup.aliyun.com/qimao/leo/leo/internal/gox/slicex"
	"codeup.aliyun.com/qimao/leo/leo/internal/gox/stringx"
	"github.com/fsnotify/fsnotify"

	"codeup.aliyun.com/qimao/leo/leo/config"
	"codeup.aliyun.com/qimao/leo/leo/log"
)

var _ config.Resource = new(Resource)
var _ config.Watcher = new(dotEnvWatcher)
var _ config.Watcher = new(environWatcher)

type options struct {
	Prefix         string
	Extension      string
	DotEnvFilename string
	Logger         log.Logger
}

type Option func(o *options)

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *options) init() {
	if stringx.IsNotBlank(o.DotEnvFilename) && stringx.IsBlank(o.Extension) {
		o.Extension = filepathx.Extension(o.DotEnvFilename)
	}
	if stringx.IsBlank(o.DotEnvFilename) && stringx.IsBlank(o.Extension) {
		o.Extension = "env"
	}
}

func Prefix(prefix string) Option {
	return func(o *options) {
		o.Prefix = prefix
	}
}

func Extension(ext string) Option {
	return func(o *options) {
		o.Extension = ext
	}
}

func DotEnvFilename(filename string) Option {
	return func(o *options) {
		o.DotEnvFilename = filename
	}
}

func Logger(log log.Logger) Option {
	return func(o *options) {
		o.Logger = log
	}
}

type Resource struct {
	options *options
}

func (r *Resource) Load(ctx context.Context) (*config.Source, error) {
	return r.loadSource()
}

func (r *Resource) Watch(ctx context.Context) (config.Watcher, error) {
	if stringx.IsBlank(r.options.DotEnvFilename) {
		source, err := r.loadFromEnv()
		if err != nil {
			return nil, err
		}
		w := &environWatcher{resource: r, source: source}
		return w, w.init(ctx)
	}
	w := &dotEnvWatcher{resource: r}
	return w, w.init(ctx)
}

func (r *Resource) loadSource() (*config.Source, error) {
	if stringx.IsBlank(r.options.DotEnvFilename) {
		return r.loadFromEnv()
	}
	return r.loadFromDotEnv()
}

func (r *Resource) loadFromEnv() (*config.Source, error) {
	environs := os.Environ()
	prefix := r.options.Prefix
	var filterEnvirons []string
	for _, environ := range environs {
		if strings.HasPrefix(environ, prefix) {
			filterEnvirons = append(filterEnvirons, environ)
		}
	}
	return config.NewSource(
		"environ",
		[]byte(strings.Join(filterEnvirons, "\n")),
		r.options.Extension,
	), nil
}

func (r *Resource) loadFromDotEnv() (*config.Source, error) {
	value, err := os.ReadFile(r.options.DotEnvFilename)
	if err != nil {
		return nil, err
	}
	return config.NewSource(
		filepath.Base(r.options.DotEnvFilename),
		value,
		r.options.Extension,
	), nil
}

type dotEnvWatcher struct {
	resource  *Resource
	fsWatcher *fsnotify.Watcher
	eventCs   []chan<- config.Event
	closeC    chan struct{}
	mutex     sync.Mutex
}

func (watcher *dotEnvWatcher) Notify(eventC chan<- config.Event) {
	watcher.mutex.Lock()
	defer watcher.mutex.Unlock()
	watcher.eventCs = slicex.AppendIfNotContains(watcher.eventCs, eventC)
}

func (watcher *dotEnvWatcher) StopNotify(eventC chan<- config.Event) {
	watcher.mutex.Lock()
	defer watcher.mutex.Unlock()
	watcher.eventCs = slicex.Remove(watcher.eventCs, eventC)
}

func (watcher *dotEnvWatcher) Close(ctx context.Context) error {
	err := watcher.fsWatcher.Close()
	watcher.closeC <- struct{}{}
	watcher.mutex.Lock()
	watcher.eventCs = nil
	watcher.mutex.Unlock()
	return err
}

func (watcher *dotEnvWatcher) init(ctx context.Context) error {
	fsWatcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	watcher.fsWatcher = fsWatcher
	filename := watcher.resource.options.DotEnvFilename
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

func (watcher *dotEnvWatcher) watch() {
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

func (watcher *dotEnvWatcher) sendError(err error) {
	if err == nil {
		return
	}
	for _, eventC := range watcher.eventCs {
		watcher.resource.options.Logger.Debug("sending error event")
		eventC <- config.ErrorEvent(err)
	}
}

func (watcher *dotEnvWatcher) handleFileChangeEvent(event fsnotify.Event) {
	// Ignore files we're not interested in.
	filename := watcher.resource.options.DotEnvFilename
	if filename != event.Name {
		return
	}
	source, err := watcher.resource.loadFromDotEnv()
	if err != nil {
		watcher.sendError(err)
		return
	}
	for _, eventC := range watcher.eventCs {
		eventC <- config.SourceEvent(source)
	}
}

type environWatcher struct {
	resource *Resource
	source   *config.Source
	eventCs  []chan<- config.Event
	closeC   chan struct{}
	mutex    sync.Mutex
}

func (watcher *environWatcher) Notify(eventC chan<- config.Event) {
	watcher.mutex.Lock()
	defer watcher.mutex.Unlock()
	watcher.eventCs = slicex.AppendIfNotContains(watcher.eventCs, eventC)
}

func (watcher *environWatcher) StopNotify(eventC chan<- config.Event) {
	watcher.mutex.Lock()
	defer watcher.mutex.Unlock()
	watcher.eventCs = slicex.Remove(watcher.eventCs, eventC)
}

func (watcher *environWatcher) Close(ctx context.Context) error {
	watcher.closeC <- struct{}{}
	watcher.mutex.Lock()
	watcher.eventCs = nil
	watcher.mutex.Unlock()
	return nil
}

func (watcher *environWatcher) init(ctx context.Context) error {
	watcher.closeC = make(chan struct{})
	watcher.watch()
	return nil
}

func (watcher *environWatcher) watch() {
	go func() {
		for {
			select {
			case <-watcher.closeC:
				return
			case <-time.After(time.Second):
				source, err := watcher.resource.loadFromEnv()
				if err != nil {
					watcher.sendError(err)
					continue
				}
				if string(source.Value()) == string(watcher.source.Value()) {
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

func (watcher *environWatcher) sendError(err error) {
	if err == nil {
		return
	}
	for _, eventC := range watcher.eventCs {
		watcher.resource.options.Logger.Debug("sending error event")
		eventC <- config.ErrorEvent(err)
	}
}

func NewResource(opts ...Option) *Resource {
	o := &options{
		Logger: log.L(),
	}
	o.apply(opts...)
	o.init()
	resource := &Resource{
		options: o,
	}
	return resource
}
