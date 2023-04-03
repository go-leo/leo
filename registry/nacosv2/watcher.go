package nacosv2

import (
	"context"
	"log"
	"sync"

	"github.com/go-leo/gox/slicex"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"

	"codeup.aliyun.com/qimao/leo/leo/registry"
)

type Watcher struct {
	instance       registry.ServiceInstance
	discovery      *Discovery
	instanceCs     []chan<- []registry.ServiceInstance
	mutex          sync.Mutex
	subscribeParam *vo.SubscribeParam
	closeC         chan struct{}
}

func (watcher *Watcher) init(ctx context.Context) error {
	watcher.subscribeParam = &vo.SubscribeParam{
		ServiceName:       watcher.instance.Name(),
		Clusters:          watcher.discovery.nacosOptions.Clusters,
		GroupName:         watcher.instance.Scheme(),
		SubscribeCallback: watcher.subscribeCallback(ctx),
	}
	return watcher.subscribe(ctx)
}

func (watcher *Watcher) Notify(instanceC chan<- []registry.ServiceInstance) {
	watcher.mutex.Lock()
	defer watcher.mutex.Unlock()
	watcher.instanceCs = slicex.AppendIfNotContains(watcher.instanceCs, instanceC)
}

func (watcher *Watcher) StopNotify(instanceC chan<- []registry.ServiceInstance) {
	watcher.mutex.Lock()
	defer watcher.mutex.Unlock()
	watcher.instanceCs = slicex.Remove(watcher.instanceCs, instanceC)
}

func (watcher *Watcher) Close(ctx context.Context) error {
	err := watcher.unsubscribe()
	watcher.closeC <- struct{}{}
	watcher.mutex.Lock()
	watcher.instanceCs = nil
	watcher.mutex.Unlock()
	return err
}

func (watcher *Watcher) subscribe(ctx context.Context) error {
	return watcher.discovery.namingClient.Subscribe(watcher.subscribeParam)
}

func (watcher *Watcher) unsubscribe() error {
	return watcher.discovery.namingClient.Unsubscribe(watcher.subscribeParam)
}

func (watcher *Watcher) subscribeCallback(ctx context.Context) func(services []model.Instance, err error) {
	return func(services []model.Instance, err error) {
		watcher.notify(ctx)
	}
}

func (watcher *Watcher) notify(ctx context.Context) {
	instances, err := watcher.discovery.GetInstances(ctx, watcher.instance)
	if err != nil {
		log.Println("failed to get instances:", err)
		return
	}
	watcher.mutex.Lock()
	for _, instanceC := range watcher.instanceCs {
		select {
		case instanceC <- instances:
		case <-watcher.closeC:
			return
		default:
		}
	}
	watcher.mutex.Unlock()
}
