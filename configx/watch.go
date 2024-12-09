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
func Watch[Config proto.Message](ctx context.Context, opts ...Option) (<-chan Config, <-chan error, func()) {
	// 初始化配置选项。
	opt := newOptions()
	opt.Apply(opts...)

	// 创建监听器
	// 遍历 opt.Resources，为每个资源创建一个监听器。
	// 收集错误列表 errs、通知通道 notifyCs 和停止函数 stops。
	var errs []error
	var notifyCs []chan *Event
	var stops []func()
	for _, watcher := range opt.Resources {
		notifyC := make(chan *Event, opt.BufferSize)
		stop, err := watcher.Watch(ctx, notifyC)
		if err != nil {
			// 如果在创建监听器时发生错误，记录错误并继续处理其他资源。
			errs = append(errs, err)
			continue
		}
		stops = append(stops, stop)
		notifyCs = append(notifyCs, notifyC)
	}

	// 处理错误：
	// 如果有任何错误，停止所有已启动的监听器，关闭通知通道，并返回错误通道。
	if len(errs) > 0 {
		for _, stop := range stops {
			brave.Go(stop)
		}
		for _, notifyC := range notifyCs {
			chanx.AsyncDiscard(notifyC)
		}
		return nil, chanx.Emit(ctx, errs...), nil
	}

	// 创建停止通道 stopC。
	// 用户调用stop函数后，将停止所有监听
	stopC := make(chan context.Context)
	for _, stop := range stops {
		brave.Go(func(stop func()) func() {
			return func() {
				select {
				case <-ctx.Done():
				case <-stopC:
					stop()
				}
			}
		}(stop))
	}

	// 启动多个协程，每个协程负责处理一个通知通道的事件。
	confC := make(chan Config, opt.BufferSize)
	errC := make(chan error, opt.BufferSize)
	var wg sync.WaitGroup
	for _, notifyC := range notifyCs {
		wg.Add(1)
		brave.Go(func(notifyC chan *Event) func() {
			return func() {
				defer wg.Done()
				for event := range notifyC {
					// 错误事件
					if ee, ok := event.AsErrorEvent(); ok {
						if errors.Is(ee.Err, ErrStopWatch) {
							// 停止监听
							close(notifyC)
							return
						}
						_ = chanx.TrySend(errC, ee.Err)
						continue
					}
					// 数据事件
					// 如果已经被stop或者上下文被取消，则停止监听。
					// 否则，加载配置。发送到通道中
					select {
					case <-stopC:
					case <-ctx.Done():
					default:
						conf, err := Load[Config](ctx, opts...)
						if err != nil {
							_ = chanx.TrySend(errC, err)
							continue
						}
						_ = chanx.TrySend(confC, conf)
						continue
					}
				}
			}
		}(notifyC))
	}

	go func() {
		wg.Wait()
		errC <- ErrStopWatch
		close(confC)
		close(errC)
	}()

	// 返回配置通道、错误通道和一个停止监听的函数，该函数会关闭停止通道。
	return confC, errC, func() { close(stopC) }
}
