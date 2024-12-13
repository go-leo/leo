package configx

import (
	"context"
	"github.com/fsnotify/fsnotify"
	"github.com/go-leo/gox/sortx"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var _ Resource = (*Environ)(nil)

type Environ struct {
	Env
	Prefix string
}

func (r *Environ) Load(ctx context.Context) ([]byte, error) {
	// 获取当前进程的所有环境变量。
	environs := os.Environ()

	//遍历这些环境变量，筛选出以 r.Prefix 开头的变量。
	var prefixedEnvirons []string
	for _, environ := range environs {
		if strings.HasPrefix(environ, r.Prefix) {
			prefixedEnvirons = append(prefixedEnvirons, environ)
		}
	}

	//将筛选后的环境变量按字典序排序。
	//将排序后的环境变量以换行符连接成一个字符串，并转换为字节切片返回。
	return []byte(strings.Join(sortx.Asc(prefixedEnvirons), "\n")), nil
}

func (r *Environ) Watch(ctx context.Context, notifyC chan<- *Event) (func(), error) {
	// 创建停止通道,用于停止监视。
	stopC := make(chan struct{})

	// 启动协程：在一个新的协程中执行监视逻辑。
	go func() {
		defer func() {
			// 在退出时发送停止监听事件。
			notifyC <- NewErrorEvent(ErrStopWatch)
		}()
		var preData []byte
		for {
			select {
			case <-ctx.Done():
				// 如果上下文 ctx 完成，则退出循环。
				return
			case <-stopC:
				// 如果收到停止信号，则退出循环。
				return
			case <-time.After(time.Second):
				// 每隔一秒钟检查一次是否被修改。
				// 加载当前环境数据，如果加载失败则发送错误事件。
				data, err := r.Load(ctx)
				if err != nil {
					notifyC <- NewErrorEvent(err)
					continue
				}

				// 比较新旧数据，如果数据发生变化，则发送数据变化事件。
				if string(data) == string(preData) {
					continue
				}
				preData = data
				notifyC <- NewDataEvent(data)
			}
		}
	}()

	// 返回停止函数：返回一个关闭 stopC 通道的函数，用于外部调用停止监视。
	return func() { close(stopC) }, nil
}

var _ Resource = (*File)(nil)

type File struct {
	Formatter Formatter
	Filename  string
}

func (r *File) Format() string {
	return r.Formatter.Format()
}

func (r *File) Load(ctx context.Context) ([]byte, error) {
	return os.ReadFile(r.Filename)
}

func (r *File) Watch(ctx context.Context, notifyC chan<- *Event) (func(), error) {
	// 初始化文件监视器：创建一个 fsnotify.Watcher 实例，用于监听文件变化。
	fsWatcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	// 添加文件目录到监视器：将文件所在的目录添加到监视器中。
	if err := fsWatcher.Add(filepath.Dir(r.Filename)); err != nil {
		return nil, err
	}
	// 创建停止通道,用于停止监视。
	stopC := make(chan struct{})

	// 启动协程：在一个新的协程中执行监视逻辑。
	go func() {
		defer func() {
			// 在退出时关闭监听，并发送停止监听事件。
			if err := fsWatcher.Close(); err != nil {
				notifyC <- NewErrorEvent(err)
			}
			notifyC <- NewErrorEvent(ErrStopWatch)
		}()
		var preData []byte
		for {
			select {
			case <-ctx.Done():
				// 如果上下文 ctx 完成，则退出循环。
				return
			case <-stopC:
				// 如果收到停止信号，则退出循环。
				return
			case event, ok := <-fsWatcher.Events:
				if !ok {
					return
				}
				// 监听文件变化事件，如果是写操作，则加载文件内容并与前一次内容进行比较，如果不同则发送数据事件。
				if !event.Has(fsnotify.Write) {
					continue
				}
				data, err := r.Load(ctx)
				if err != nil {
					notifyC <- NewErrorEvent(err)
					continue
				}
				if string(data) == string(preData) {
					continue
				}
				preData = data
				notifyC <- NewDataEvent(data)
			case err, ok := <-fsWatcher.Errors:
				if !ok {
					return
				}
				// 监听文件监视器的错误事件，并发送错误事件。
				notifyC <- NewErrorEvent(err)
			}
		}
	}()
	return func() { close(stopC) }, nil
}
