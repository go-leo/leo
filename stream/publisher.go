package stream

type Publisher interface {
	Topic() string
	Publish(messages ...Message) (any, error)
	Close() error
}
