package nacos

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sync/atomic"
	"time"

	"github.com/go-leo/netx/httpx"

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
	scheme    string
	host      string
	port      string
	namespace string
	group     string
	dataId    string
	state     uint64
	stopCtx   context.Context
	stopFunc  context.CancelFunc
	log       log.Logger
	client    *http.Client
}

func (watcher *Watcher) Start(ctx context.Context) (<-chan *config.Event, error) {
	if !atomic.CompareAndSwapUint64(&watcher.state, stop, start) {
		return nil, errors.New("watcher is started")
	}
	watcher.stopCtx, watcher.stopFunc = context.WithCancel(ctx)
	eventC := make(chan *config.Event)
	go watcher.watch(eventC)
	return eventC, nil
}

func (watcher *Watcher) Stop(_ context.Context) error {
	if !atomic.CompareAndSwapUint64(&watcher.state, start, stop) {
		return errors.New("watcher is stopped")
	}
	watcher.stopFunc()
	return nil
}

func (watcher *Watcher) watch(eventC chan *config.Event) {
	defer close(eventC)
	for {
		select {
		case <-watcher.stopCtx.Done():
			watcher.log.Infof("stop watch nacos config")
			return
		default:
			// 获取远程配置notificationID
			modify, err := watcher.isModified(watcher.stopCtx)
			if err != nil {
				watcher.log.Error(err)
				time.Sleep(30 * time.Second)
				continue
			}
			if !modify {
				continue
			}
			// 不同，有变化，同步
			watcher.log.Infof("get config DataId: %s, Group: %s", watcher.dataId, watcher.group)
			ctx, _ := context.WithTimeout(watcher.stopCtx, time.Second) // nolint
			uri := fmt.Sprintf("%s://%s:%s/nacos/v1/cs/configs", watcher.scheme, watcher.host, watcher.port)
			request, err := new(httpx.RequestBuilder).
				Get().
				URLString(uri).
				Query("tenant", watcher.namespace).
				Query("group", watcher.group).
				Query("dataId", watcher.dataId).
				Build(ctx)
			if err != nil {
				watcher.log.Error(err)
				time.Sleep(30 * time.Second)
				continue
			}
			helper := httpx.NewResponseHelper(watcher.client.Do(request))
			if err := helper.Err(); err != nil {
				watcher.log.Error(err)
				time.Sleep(30 * time.Second)
				continue
			}
			content, err := helper.BytesBody()
			if err != nil {
				watcher.log.Error(err)
				time.Sleep(30 * time.Second)
				continue
			}
			event := &EventDescription{namespace: watcher.namespace, group: watcher.group, dataId: watcher.dataId}
			eventC <- config.NewContentEvent(event, content)
		}
	}
}

func (watcher *Watcher) isModified(ctx context.Context) (bool, error) {
	uri := ""
	form := url.Values{}
	form.Set("Listening-Configs", "")
	request, err := new(httpx.RequestBuilder).
		Post().
		URLString(uri).
		Header("Long-Pulling-Timeout", "30000").
		FormBody(form).
		Build(ctx)
	if err != nil {
		return false, err
	}
	helper := httpx.NewResponseHelper(watcher.client.Do(request))
	if err := helper.Err(); err != nil {
		return false, err
	}
	if helper.StatusCode() != 200 {
		return false, fmt.Errorf("failed call nocos, status: %d, %s", helper.StatusCode(), http.StatusText(helper.StatusCode()))
	}
	body, err := helper.BytesBody()
	if err != nil {
		return false, err
	}
	// 如果配置无变化：会返回空串
	if len(body) <= 0 {
		return false, nil
	}
	// 如果配置变化, dataId%02group%02tenant%01
	return true, nil
}

type WatcherOption func(watcher *Watcher)

func WithScheme(scheme string) WatcherOption {
	return func(watcher *Watcher) {
		watcher.scheme = scheme
	}
}

func WithLogger(log log.Logger) WatcherOption {
	return func(watcher *Watcher) {
		watcher.log = log
	}
}

func NewWatcher(
	host string,
	port string,
	namespace string,
	group string,
	dataId string,
	opts ...WatcherOption,
) *Watcher {
	watcher := &Watcher{
		scheme:    "http",
		host:      host,
		port:      port,
		namespace: namespace,
		group:     group,
		dataId:    dataId,
		state:     0,
		stopCtx:   nil,
		stopFunc:  nil,
		log:       &log.Discard{},
		client:    httpx.PooledClient(),
	}
	for _, opt := range opts {
		opt(watcher)
	}
	return watcher
}
