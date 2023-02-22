package stream

type Bridge interface {
	Input() Subscriber
	Output() Publisher
	Process(msg Message) ([]Message, error)
}
