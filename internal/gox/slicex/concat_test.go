package slicex

import "testing"

func TestConcat(t *testing.T) {
	t.Log(Concat([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 0}))
	t.Log(Concat([]float32{1, 2, 3, 4, 5}, []float32{6, 7, 8, 9, 0}))
	t.Log(Concat([]string{"1", "2", "3", "4", "5"}, []string{"6", "7", "8", "9", "0"}))
}
