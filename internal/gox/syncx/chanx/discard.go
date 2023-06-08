package chanx

func Discard[T any](c <-chan T) {
	for range c {
	}
}

func AsyncDiscard[T any](c <-chan T) {
	go Discard(c)
}
