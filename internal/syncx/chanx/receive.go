package chanx

func ReceiveUtilClosed[T any](c <-chan T) []T {
	var ts []T
	for t := range c {
		ts = append(ts, t)
	}
	return ts
}
