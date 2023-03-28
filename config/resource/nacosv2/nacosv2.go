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
	w := &Watcher{
		dataId: r.dataID,
		group:  r.group,
	}
	return w, w.init(ctx)
}

type Watcher struct {
	dataId       string
	group        string
	resource     *Resource
	changeC      chan config.Event
	eventCs      []chan<- config.Event
	mutex        sync.Mutex
	configClient config_client.IConfigClient
	closeC       chan struct{}
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
	err := watcher.configClient.CancelListenConfig(vo.ConfigParam{
		DataId: watcher.dataId,
		Group:  watcher.group,
	})
	watcher.closeC <- struct{}{}
	watcher.mutex.Lock()
	watcher.eventCs = nil
	watcher.mutex.Unlock()
	return err
}

func (watcher *Watcher) init(ctx context.Context) error {
	err := watcher.configClient.ListenConfig(vo.ConfigParam{
		DataId: watcher.dataId,
		Group:  watcher.group,
		OnChange: func(namespace, group, dataId, data string) {
			watcher.changeC <- config.SourceEvent(&config.Source{
				Name:      dataId + "." + group,
				Value:     []byte(dataId),
				Extension: watcher.resource.options.Extension,
			})
		},
	})
	if err != nil {
		return err
	}
	watcher.watch()
	return nil
}

func (watcher *Watcher) watch() {
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

func (watcher *Watcher) sendError(err error) {
	if err == nil {
		return
	}
	for _, eventC := range watcher.eventCs {
		eventC <- config.ErrorEvent(err)
	}
}

func (watcher *Watcher) handleChangeEvent(event config.Event) {
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
	for _, opt := range opts {
		opt(o)
	}
	resource := &Resource{
		options:      o,
		configClient: configClient,
		dataID:       dataID,
		group:        group,
	}
	return resource, nil
}

type ConfigClientFactory interface {
	Create() (config_client.IConfigClient, error)
}

type ConfigClientFactoryFunc func() (config_client.IConfigClient, error)

func (f ConfigClientFactoryFunc) Create() (config_client.IConfigClient, error) {
	return f()
}
