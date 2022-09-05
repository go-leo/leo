package slicex

import "fmt"

func Stringfy[S ~[]E, E any](s S) []string {
	result := make([]string, 0, len(s))
	for _, e := range s {
		result = append(result, fmt.Sprint(e))
	}
	return result
}
