package configx

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/structpb"
)

// ErrStopWatch 定义了一个停止监听的错误事件。
var ErrStopWatch = errors.New("config: stop watch")

// EventKind 定义了一个事件类型接口，用于标识配置数据的变化。
type EventKind interface {
	isEventKind()
}

// ErrorEvent 定义了一个错误事件，用于表示监听配置数据时发生的错误。
type ErrorEvent struct{ Err error }

func (ErrorEvent) isEventKind() {}

// DataEvent 定义了一个数据变化事件，用于表示监听配置数据时发生的数据变化。
type DataEvent struct{ Data *structpb.Struct }

func (DataEvent) isEventKind() {}

// Event 表示一个动态类型的值，它可以是 ErrorEvent 或 DataEvent。
type Event struct{ Kind EventKind }

func (e *Event) GetKind() EventKind {
	if e != nil {
		return e.Kind
	}
	return nil
}

func (e *Event) AsErrorEvent() (*ErrorEvent, bool) {
	if ee, ok := e.GetKind().(*ErrorEvent); ok {
		return ee, true
	}
	return nil, false
}

func (e *Event) AsDataEvent() (*DataEvent, bool) {
	if de, ok := e.GetKind().(*DataEvent); ok {
		return de, true
	}
	return nil, false
}

// Formatter 定义了格式化器接口，用于标识不同的数据格式，例如 JSON、YAML 等。
type Formatter interface {
	// Format 返回表示格式的字符串。
	Format() string
}

// FormatSupporter 定义了一个格式支持者接口，用于判断是否支持某种格式。
type FormatSupporter interface {
	// Support 判断是否支持某种格式。
	Support(format Formatter) bool
}

// Parser 定义了一个解析器接口。用于解析加载器提供的字节切片数据，并将其转换为 structpb.Struct 类型。
type Parser interface {
	FormatSupporter
	// Parse 解析字节切片数据，并返回 structpb.Struct 类型和可能的错误
	Parse(source []byte) (*structpb.Struct, error)
}

// Merger 定义了一个合并器接口。用于将多个 structpb.Struct 对象合并成一个。
type Merger interface {
	// Merge 合并多个 structpb.Struct 对象，并返回合并后的 structpb.Struct 类型
	Merge(values ...*structpb.Struct) *structpb.Struct
}

// Loader 定义了一个加载器接口，用于从某种数据源（如文件、网络等）加载配置数据，并返回字节切片和可能的错误。
type Loader interface {
	// Load 返回字节切片和可能的错误
	Load(ctx context.Context) ([]byte, error)
}

// Watcher 定义了一个观察者接口，用于监控配置数据的变化，返回一个通道（用于接收变化通知）、一个停止监听函数和可能的错误。
type Watcher interface {
	Formatter
	// Watch 监控配置数据的变化，接受一个通道（用于接收事件通知）, 返回一个停止监听函数和立即发生的错误
	// 如果 Context 被取消，必须停止监听。
	// 停止监听必须要发送 ErrStopWatch 错误事件。
	Watch(ctx context.Context, notifyC chan<- *Event) (func(), error)
}

// Resource 定义了一个资源接口，继承了 Loader 和 Watcher 接口，表示一个既能加载配置数据又能监控其变化的资源。
type Resource interface {
	Loader
	Watcher
}
