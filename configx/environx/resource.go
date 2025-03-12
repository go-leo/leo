package environx

import (
	"context"
	"github.com/go-leo/gox/sortx"
	"github.com/go-leo/leo/v3/configx"
	"os"
	"strings"
	"time"
)

var _ configx.Resource = (*Resource)(nil)

type Resource struct {
	configx.Env
	Prefix string
}

func (r *Resource) Load(ctx context.Context) ([]byte, error) {
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

func (r *Resource) Watch(ctx context.Context, notifyC chan<- *configx.Event) error {
	// 启动协程：在一个新的协程中执行监视逻辑。
	go func() {
		defer func() {
			// 在退出时发送停止监听事件。
			notifyC <- configx.NewErrorEvent(configx.ErrStopWatch)
		}()
		var preData []byte
		for {
			select {
			case <-ctx.Done():
				// 如果上下文 ctx 完成，则退出循环。
				return
			case <-time.After(time.Second):
				// 每隔一秒钟检查一次是否被修改。
				// 加载当前环境数据，如果加载失败则发送错误事件。
				data, err := r.Load(ctx)
				if err != nil {
					notifyC <- configx.NewErrorEvent(err)
					continue
				}

				// 比较新旧数据，如果数据发生变化，则发送数据变化事件。
				if string(data) == string(preData) {
					continue
				}
				preData = data
				notifyC <- configx.NewDataEvent(data)
			}
		}
	}()
	return nil
}
