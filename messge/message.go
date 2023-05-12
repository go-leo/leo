package message

import "time"

// Message is a Message. Message is emitted by Publisher and received by Subscriber.
type Message struct {
	ID      string
	Time    time.Time
	Payload []byte
	Header  Header
}
