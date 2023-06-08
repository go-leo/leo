package chanx

func Any[T any](channels ...<-chan T) T {
	cancelC := make(chan struct{})
	valueC := make(chan T)
	for _, ch := range channels {
		go func(ch <-chan T) {
			select {
			case <-cancelC:
				return
			case v := <-ch:
				valueC <- v
				close(cancelC)
			}
		}(ch)
	}
	return <-valueC
}
