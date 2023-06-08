package chanx

func All[T any](channels ...<-chan T) []T {
	values := make([]T, 0, len(channels))
	for _, ch := range channels {
		values = append(values, <-ch)
	}
	return values
}
