package configx

import (
	"context"
	"google.golang.org/protobuf/types/known/structpb"
)

// Formatter 定义了格式化器接口，用于标识不同的数据格式，例如 JSON、YAML 等。
type Formatter interface {
	// Format 返回表示格式的字符串。
	Format() string
}

// Parser 定义了一个解析器接口。用于判断是否支持某种格式和解析加载器/监视器提供的字节切片数据，并将其转换为 structpb.Struct 类型。
type Parser interface {
	// Support 判断是否支持某种格式。
	Support(format Formatter) bool
	// Parse 解析字节切片数据，并返回 structpb.Struct 类型和可能的错误
	Parse(data []byte) (*structpb.Struct, error)
}

// Merger 定义了一个合并器接口。用于将多个 structpb.Struct 对象合并成一个。
type Merger interface {
	// Merge 合并多个 structpb.Struct 对象，并返回合并后的 structpb.Struct 类型
	Merge(values ...*structpb.Struct) *structpb.Struct
}

// Loader 定义了一个加载器接口，用于从某种数据源（如文件、网络等）加载配置数据，并返回字节切片和可能的错误。
type Loader interface {
	Formatter
	// Load 返回字节切片和可能的错误
	Load(ctx context.Context) ([]byte, error)
}

// Watcher 定义了一个观察者接口，用于监控配置数据的变化，返回一个通道（用于接收变化通知）、一个停止监听函数和可能的错误。
type Watcher interface {
	Formatter
	// Watch 监控配置数据的变化，接受一个通道（用于接收事件通知）, 返回立即发生的错误
	// 如果 Context 被取消，则停止监听。
	// 停止监听必须要发送 ErrStopWatch 错误事件。
	Watch(ctx context.Context, notifyC chan<- *Event) error
}

// Resource 定义了一个资源接口，继承了 Loader 和 Watcher 接口，表示一个既能加载配置数据又能监控其变化的资源。
type Resource interface {
	Loader
	Watcher
}
