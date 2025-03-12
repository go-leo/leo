package filex

import (
	"context"
	"github.com/fsnotify/fsnotify"
	"github.com/go-leo/leo/v3/configx"
	"os"
	"path/filepath"
	"strings"
)

var _ configx.Resource = (*Resource)(nil)

type Resource struct {
	Formatter configx.Formatter
	Filename  string
}

func (r *Resource) Format() string {
	if r.Formatter == nil {
		return strings.TrimPrefix(filepath.Ext(r.Filename), ".")
	}
	return r.Formatter.Format()
}

func (r *Resource) Load(ctx context.Context) ([]byte, error) {
	return os.ReadFile(r.Filename)
}

func (r *Resource) Watch(ctx context.Context, notifyC chan<- *configx.Event) error {
	// 初始化文件监视器：创建一个 fsnotify.Watcher 实例，用于监听文件变化。
	fsWatcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	// 添加文件目录到监视器：将文件所在的目录添加到监视器中。
	if err := fsWatcher.Add(filepath.Dir(r.Filename)); err != nil {
		return err
	}
	// 启动协程：在一个新的协程中执行监视逻辑。
	go func() {
		defer func() {
			// 在退出时关闭监听，并发送停止监听事件。
			if err := fsWatcher.Close(); err != nil {
				notifyC <- configx.NewErrorEvent(err)
			}
			notifyC <- configx.NewErrorEvent(configx.ErrStopWatch)
		}()
		var preData []byte
		for {
			select {
			case <-ctx.Done():
				// 如果上下文 ctx 完成，则退出循环。
				return
			case event, ok := <-fsWatcher.Events:
				if !ok {
					return
				}
				if filepath.Clean(event.Name) != r.Filename {
					continue
				}
				// 监听文件变化事件，如果是写操作，则加载文件内容并与前一次内容进行比较，如果不同则发送数据事件。
				if !event.Has(fsnotify.Write) && !event.Has(fsnotify.Create) {
					continue
				}
				data, err := r.Load(ctx)
				if err != nil {
					notifyC <- configx.NewErrorEvent(err)
					continue
				}
				if string(data) == string(preData) {
					continue
				}
				preData = data
				notifyC <- configx.NewDataEvent(data)
			case err, ok := <-fsWatcher.Errors:
				if !ok {
					return
				}
				// 监听文件监视器的错误事件，并发送错误事件。
				notifyC <- configx.NewErrorEvent(err)
			}
		}
	}()
	return nil
}
