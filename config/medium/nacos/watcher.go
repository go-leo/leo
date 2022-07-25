package nacos

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"

	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/vo"

	"github.com/go-leo/leo/config"
	"github.com/go-leo/leo/log"
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
	client   config_client.IConfigClient
	group    string
	dataId   string
	state    uint64
	stopCtx  context.Context
	stopFunc context.CancelFunc
	log      log.Logger
}

func (watcher *Watcher) Start(ctx context.Context) (<-chan *config.Event, error) {
	if !atomic.CompareAndSwapUint64(&watcher.state, stop, start) {
		return nil, errors.New("watcher is started")
	}
	eventC := make(chan *config.Event)
	err := watcher.client.ListenConfig(vo.ConfigParam{
		DataId: watcher.dataId,
		Group:  watcher.group,
		OnChange: func(namespace, group, dataId, data string) {
			event := &EventDescription{namespace: namespace, group: group, dataId: dataId}
			eventC <- config.NewContentEvent(event, []byte(data))
		},
	})
	if err != nil {
		err := fmt.Errorf("failed to listen config, DataId: %s, Group: %s, %w", watcher.dataId, watcher.group, err)
		return nil, err
	}
	watcher.stopCtx, watcher.stopFunc = context.WithCancel(ctx)
	go func() {
		defer close(eventC)
		<-watcher.stopCtx.Done()
		err := watcher.client.CancelListenConfig(vo.ConfigParam{
			DataId: watcher.dataId,
			Group:  watcher.group,
		})
		if err != nil {
			err = fmt.Errorf("failed to cancel listen config, %w", err)
			eventC <- config.NewErrEvent(err)
		}
	}()
	return eventC, nil
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

func NewWatcher(
	client config_client.IConfigClient,
	group string,
	dataId string,
	opts ...WatcherOption,
) *Watcher {
	watcher := &Watcher{
		client: client,
		group:  group,
		dataId: dataId,
		log:    nil,
	}
	for _, opt := range opts {
		opt(watcher)
	}
	return watcher
}
