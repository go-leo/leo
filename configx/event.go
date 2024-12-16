package configx

import "errors"

// ErrStopWatch 定义了一个停止监听的错误事件。
var ErrStopWatch = errors.New("config: stop watch")

// EventKind 定义了一个事件类型接口，用于标识配置数据的变化。
// see: ErrorEvent and DataEvent
type EventKind interface {
	isEventKind()
}

// ErrorEvent 定义了一个错误事件，用于表示监听配置数据时发生的错误。
type ErrorEvent struct{ Err error }

func (ErrorEvent) isEventKind() {}

// DataEvent 定义了一个数据变化事件，用于表示监听配置数据时发生的数据变化。
type DataEvent struct{ Data []byte }

func (DataEvent) isEventKind() {}

// Event 表示一个动态类型的值，它可以是 ErrorEvent 或 DataEvent。
type Event struct{ kind EventKind }

func (e *Event) GetKind() EventKind {
	if e != nil {
		return e.kind
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

func NewErrorEvent(err error) *Event {
	return &Event{kind: &ErrorEvent{Err: err}}
}

func NewDataEvent(data []byte) *Event {
	return &Event{kind: &DataEvent{Data: data}}
}
