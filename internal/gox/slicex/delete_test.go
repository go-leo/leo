package slicex_test

import (
	"testing"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/slicex"
)

func TestDelete(t *testing.T) {
	t.Log(slicex.Delete([]int{0, 1, 2, 3, 4, 5, 6}, 0))
	t.Log(slicex.Delete([]int{0, 1, 2, 3, 4, 5, 6}, 1))
	t.Log(slicex.Delete([]int{0, 1, 2, 3, 4, 5, 6}, 2))
	t.Log(slicex.Delete([]int{0, 1, 2, 3, 4, 5, 6}, 3))
	t.Log(slicex.Delete([]int{0, 1, 2, 3, 4, 5, 6}, 4))
	t.Log(slicex.Delete([]int{0, 1, 2, 3, 4, 5, 6}, 5))
	t.Log(slicex.Delete([]int{0, 1, 2, 3, 4, 5, 6}, 6))
	// t.Log(slicex.Delete([]int{0, 1, 2, 3, 4, 5, 6}, -1))
	// t.Log(slicex.Delete([]int{0, 1, 2, 3, 4, 5, 6}, 7))
}
