package slicex

import "testing"

func TestRemove(t *testing.T) {
	ints := Remove([]int{1, 2, 3, 4, 5}, 2)
	t.Log(ints)
}
