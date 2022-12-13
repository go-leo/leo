package nacosv2

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"

	"github.com/go-leo/leo/config"
	"github.com/go-leo/leo/log"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

const (
	stop uint64 = iota
	start
)

var _ config.Watcher = new(Watcher)

type EventDescription struct {
	namespace string
	group     string
	dataId    string
}

func (ed *EventDescription) String() string {
	return fmt.Sprintf("%s.%s.%s", ed.namespace, ed.group, ed.dataId)
}

type Watcher struct {
	group    string
	dataId   string
	state    uint64
	stopCtx  context.Context
	stopFunc context.CancelFunc
	log      log.Logger
	client   config_client.IConfigClient
}

func (watcher *Watcher) Start(ctx context.Context) (<-chan *config.Event, error) {
	if !atomic.CompareAndSwapUint64(&watcher.state, stop, start) {
		return nil, errors.New("watcher is started")
	}
	watcher.stopCtx, watcher.stopFunc = context.WithCancel(ctx)
	eventC := make(chan *config.Event)
	_ = watcher.client.ListenConfig(vo.ConfigParam{
		DataId: watcher.dataId,
		Group:  watcher.group,
		OnChange: func(namespace, group, dataId, data string) {
			watcher.log.Info("config changed group:" + group + ", dataId:" + dataId + ", content:" + data)
		},
	})
	return eventC, nil
}

func (watcher *Watcher) Stop(_ context.Context) error {
	if !atomic.CompareAndSwapUint64(&watcher.state, start, stop) {
		return errors.New("watcher is stopped")
	}
	watcher.stopFunc()
	err := watcher.client.CancelListenConfig(vo.ConfigParam{
		DataId: watcher.dataId,
		Group:  watcher.group,
	})
	if err != nil {
		return err
	}
	return nil
}

type WatcherOption func(watcher *Watcher)

func WithLogger(log log.Logger) WatcherOption {
	return func(watcher *Watcher) {
		watcher.log = log
	}
}

func NewWatcher(
	client config_client.IConfigClient,
	group string,
	dataId string,
	opts ...WatcherOption,
) *Watcher {
	watcher := &Watcher{
		group:  group,
		dataId: dataId,
		log:    &log.Discard{},
		client: client,
	}
	for _, opt := range opts {
		opt(watcher)
	}
	return watcher
}
