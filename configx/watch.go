package configx

import (
	"context"
	"errors"
	"github.com/go-leo/gox/syncx/brave"
	"github.com/go-leo/gox/syncx/chanx"
	"google.golang.org/protobuf/proto"
	"sync"
)

// Watch 函数用于监听指定资源的配置变化，并通过通道发送这些配置。
func Watch[Config proto.Message](ctx context.Context, opts ...Option) (<-chan Config, error) {
	// 初始化配置选项。
	opt := newOptions()
	opt.Apply(opts...)

	// 创建监听器
	// 遍历 opt.Resources，为每个资源创建一个监听器。创建监听配置变化通道 notifyCs。
	var notifyCs []chan *Event
	var errs []error
	watchCtx, cancelFunc := context.WithCancel(ctx)
	for _, watcher := range opt.Resources {
		notifyC := make(chan *Event, opt.BufferSize)
		// 创建一个带有取消功能的上下文，并调用监听器的Watch方法来启动监听。
		if err := watcher.Watch(watchCtx, notifyC); err != nil {
			errs = append(errs, err)
			break
		}
		notifyCs = append(notifyCs, notifyC)
	}
	// 如果有一个监听错误，则返回错误。
	if len(errs) > 0 {
		// 关闭所有监听器
		cancelFunc()
		return nil, errors.Join(errs...)
	}

	// 启动多个goroutine，每个goroutine负责处理一个通知通道的事件。
	confC := make(chan Config, opt.BufferSize)
	var wg sync.WaitGroup
	for _, notifyC := range notifyCs {
		wg.Add(1)
		brave.Go(func(notifyC chan *Event) func() {
			return func() {
				defer wg.Done()
				for event := range notifyC {
					// 错误事件
					if errEvent, ok := event.AsErrorEvent(); ok {
						// 如果是停止监听信号，则停止监听并退出goroutine。
						if errors.Is(errEvent.Err, ErrStopWatch) {
							close(notifyC)
							return
						}
						// 其他错误信号。
						opt.ErrorCallback(errEvent.Err)
						continue
					}
					// 数据事件, 重新加载配置。发送到通道中
					conf, err := Load[Config](ctx, opts...)
					if err != nil {
						opt.ErrorCallback(err)
						continue
					}
					_ = chanx.TrySend(confC, conf)
					continue
				}
			}
		}(notifyC))
	}

	// 等待所有都监听完毕
	go func() {
		wg.Wait()
		cancelFunc()
		opt.ErrorCallback(ErrStopWatch)
		close(confC)
	}()

	// 返回配置通道、错误通道和一个停止监听的函数，该函数会关闭停止通道。
	return confC, nil
}
