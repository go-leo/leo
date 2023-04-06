package nacosv2

import (
	"context"
	"sync"

	"github.com/go-leo/gox/slicex"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"

	"codeup.aliyun.com/qimao/leo/leo/config"
	"codeup.aliyun.com/qimao/leo/leo/log"
)

var _ config.Resource = new(Resource)

var _ config.Watcher = new(watcher)

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
	return func(o *options) {
		o.Extension = Extension
	}
}

func Logger(log log.Logger) Option {
	return func(o *options) {
		o.Logger = log
	}
}

type Resource struct {
	options      *options
	configClient config_client.IConfigClient
	dataID       string
	group        string
}

func (r *Resource) Load(ctx context.Context) (*config.Source, error) {
	content, err := r.configClient.GetConfig(vo.ConfigParam{
		DataId:   r.dataID,
		Group:    r.group,
		OnChange: nil,
	})
	if err != nil {
		return nil, err
	}
	return &config.Source{
		Name:      r.dataID + "." + r.group,
		Value:     []byte(content),
		Extension: r.options.Extension,
	}, nil
}

func (r *Resource) Watch(ctx context.Context) (config.Watcher, error) {
	w := &watcher{resource: r}
	return w, w.init(ctx)
}

type watcher struct {
	resource *Resource
	changeC  chan config.Event
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
	err := watcher.resource.configClient.CancelListenConfig(vo.ConfigParam{
		DataId: watcher.resource.dataID,
		Group:  watcher.resource.group,
	})
	watcher.closeC <- struct{}{}
	watcher.mutex.Lock()
	watcher.eventCs = nil
	watcher.mutex.Unlock()
	return err
}

func (watcher *watcher) init(ctx context.Context) error {
	err := watcher.resource.configClient.ListenConfig(vo.ConfigParam{
		DataId: watcher.resource.dataID,
		Group:  watcher.resource.group,
		OnChange: func(namespace, group, dataId, data string) {
			watcher.changeC <- config.SourceEvent(&config.Source{
				Name:      dataId + "." + group,
				Value:     []byte(data),
				Extension: watcher.resource.options.Extension,
			})
		},
	})
	if err != nil {
		return err
	}
	watcher.changeC = make(chan config.Event)
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
			case event, ok := <-watcher.changeC:
				if !ok {
					return
				}
				watcher.handleChangeEvent(event)
			}
		}
	}()
}

func (watcher *watcher) sendError(err error) {
	if err == nil {
		return
	}
	for _, eventC := range watcher.eventCs {
		eventC <- config.ErrorEvent(err)
	}
}

func (watcher *watcher) handleChangeEvent(event config.Event) {
	if event == nil {
		return
	}
	for _, eventC := range watcher.eventCs {
		eventC <- event
	}
}

func NewResource(dataID, group string, factory ConfigClientFactoryFunc, opts ...Option) (*Resource, error) {
	configClient, err := factory.Create()
	if err != nil {
		return nil, err
	}
	o := &options{
		Logger: log.Discard{},
	}
	o.apply(opts...)
	o.init()
	resource := &Resource{
		options:      o,
		configClient: configClient,
		dataID:       dataID,
		group:        group,
	}
	return resource, nil
}
