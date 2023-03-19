package chanx

func Separate[T any](channel <-chan T, size int) []<-chan T {
	channels := make([]<-chan T, 0, size)
	for i := 0; i < size; i++ {
		ch := make(chan T, cap(channels))
		channels = append(channels, ch)
		go func(ch chan T) {
			for v := range channel {
				ch <- v
			}
		}(ch)
	}
	return channels
}
