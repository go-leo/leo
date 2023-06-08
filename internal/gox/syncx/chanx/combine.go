package chanx

import "sync"

func Combine[T any](channels ...<-chan T) <-chan T {
	c := make(chan T, len(channels))
	var wg sync.WaitGroup
	for _, ch := range channels {
		wg.Add(1)
		go func(ch <-chan T) {
			defer wg.Done()
			for v := range ch {
				c <- v
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	return c
}
