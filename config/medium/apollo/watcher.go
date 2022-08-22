package apollo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"sync/atomic"
	"time"

	"github.com/go-leo/leo/common/filex"
	"github.com/go-leo/leo/common/httpx"
	"github.com/go-leo/leo/common/stringx"
	"github.com/go-leo/leo/config"
	"github.com/go-leo/leo/log"
)

const (
	stop uint64 = iota
	start
)

const initNotificationID = -1

var _ config.Watcher = new(Watcher)

type EventDescription struct {
	appID         string
	cluster       string
	namespaceName string
}

func (ed *EventDescription) String() string {
	return fmt.Sprintf("%s.%s.%s", ed.appID, ed.cluster, ed.namespaceName)
}

type Watcher struct {
	scheme         string
	host           string
	port           string
	appID          string
	cluster        string
	namespaceName  string
	secret         string
	cli            *http.Client
	notificationID int64
	state          uint64
	stopCtx        context.Context
	stopFunc       context.CancelFunc
	log            log.Logger
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

func (watcher *Watcher) watch(eventC chan *config.Event) {
	defer close(eventC)
	for {
		select {
		case <-watcher.stopCtx.Done():
			watcher.log.Infof("stop watch apollo config, appID: %s, cluster: %s, namespaceName: %s", watcher.appID, watcher.cluster, watcher.namespaceName)
			return
		default:
			// 获取远程配置notificationID
			modify, err := watcher.isModified(watcher.stopCtx)
			if err != nil {
				time.Sleep(30 * time.Second)
				continue
			}
			if !modify {
				continue
			}
			// 不同，有变化，同步
			uri := fmt.Sprintf("%s://%s:%s/configs/%s/%s/%s", watcher.scheme, watcher.host, watcher.port, watcher.appID, watcher.cluster, watcher.namespaceName)
			ctx, cancelFunc := context.WithTimeout(watcher.stopCtx, time.Second)
			defer cancelFunc()
			contentType := filex.ExtName(watcher.namespaceName)
			if stringx.IsBlank(contentType) {
				contentType = "properties"
			}
			content, err := getConfigContent(ctx, uri, watcher.appID, watcher.secret, contentType, watcher.cli)
			if err != nil {
				continue
			}
			description := &EventDescription{
				appID:         watcher.appID,
				cluster:       watcher.cluster,
				namespaceName: watcher.namespaceName,
			}
			eventC <- config.NewContentEvent(description, []byte(content))
		}
	}
}

func (watcher *Watcher) isModified(ctx context.Context) (bool, error) {
	notifications := []*Notification{{NamespaceName: watcher.namespaceName, NotificationID: watcher.notificationID}}
	notificationsJson, err := json.Marshal(notifications)
	if err != nil {
		return false, err
	}
	uri := url.URL{
		Scheme: watcher.scheme,
		Host:   net.JoinHostPort(watcher.host, watcher.port),
		Path:   "/notifications/v2",
	}
	query := url.Values{}
	query.Set("appId", watcher.appID)
	query.Set("cluster", watcher.cluster)
	query.Set("notifications", string(notificationsJson))
	uri.RawQuery = query.Encode()
	ctx, cancel := context.WithTimeout(ctx, 120*time.Second)
	defer cancel()
	rawResp, err := requestApollo(ctx, uri.String(), watcher.appID, watcher.secret, watcher.cli)
	if err != nil {
		return false, err
	}
	defer rawResp.Body.Close()
	switch rawResp.StatusCode {
	case http.StatusOK:
		data, err := io.ReadAll(rawResp.Body)
		if err != nil {
			return false, err
		}
		var resp []*Notification
		if err = json.Unmarshal(data, &resp); err != nil {
			return false, err
		}
		watcher.notificationID = resp[0].NotificationID
		return true, nil
	case http.StatusNotModified:
		return false, nil
	}
	return false, fmt.Errorf("failed to notifications, StatusCode: %d, Status: %s", rawResp.StatusCode, rawResp.Status)
}

func (watcher *Watcher) Stop(_ context.Context) error {
	if !atomic.CompareAndSwapUint64(&watcher.state, start, stop) {
		return errors.New("watcher is stopped")
	}
	watcher.stopFunc()
	return nil
}

type WatcherOption func(watcher *Watcher)

func WithScheme(scheme string) WatcherOption {
	return func(watcher *Watcher) {
		watcher.scheme = scheme
	}
}

func WithSecret(secret string) WatcherOption {
	return func(watcher *Watcher) {
		watcher.secret = secret
	}
}

func WithLogger(log log.Logger) WatcherOption {
	return func(watcher *Watcher) {
		watcher.log = log
	}
}

func NewWatcher(host string, port string, appID string, cluster string, namespaceName string, opts ...WatcherOption) *Watcher {
	w := &Watcher{
		scheme:         "http",
		host:           host,
		port:           port,
		appID:          appID,
		cluster:        cluster,
		namespaceName:  namespaceName,
		secret:         "",
		cli:            httpx.PooledClient(),
		notificationID: initNotificationID,
		state:          0,
		stopCtx:        nil,
		stopFunc:       nil,
		log:            log.Discard{},
	}
	for _, opt := range opts {
		opt(w)
	}
	return w
}
