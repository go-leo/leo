package stream

type Publisher interface {
	Topic() string
	Publish(message Message) (any, error)
	Close() error
}
