package chanx

func AppendSendChannel[T any](c []<-chan T, channels ...chan T) []<-chan T {
	for _, ch := range channels {
		c = append(c, ch)
	}
	return c
}

func AppendReceiveChannel[T any](c []chan<- T, channels ...chan T) []chan<- T {
	for _, ch := range channels {
		c = append(c, ch)
	}
	return c
}
