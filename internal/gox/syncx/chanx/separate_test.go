package chanx

import (
	"fmt"
	"testing"
	"time"
)

func TestSeparateChannels(t *testing.T) {
	c := make(chan int, 10)
	channels := Separate(c, 4)
	for _, ch := range channels {
		go func(ch <-chan int) {
			for v := range ch {
				fmt.Println(v)
			}
		}(ch)
	}

	go func() {
		i := 0
		for {
			c <- i
			i++
			time.Sleep(time.Millisecond)
		}
	}()

	select {}
}
