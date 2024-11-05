package conf

//import "context"
//
//// Resource is a loader that can be used to load source config.
//type Resource interface {
//	Load(ctx context.Context) (Source, error)
//	Watch(ctx context.Context) (Watcher, error)
//}
//
//type Source interface {
//	Content() []byte
//}
//
//// Watcher monitors whether the data source has changed and, if so, notifies the changed event
//type Watcher interface {
//	Notify(eventC chan<- Event)
//	StopNotify(eventC chan<- Event)
//	Close(ctx context.Context) error
//}
