package sortx_test

import (
	"testing"

	"github.com/go-leo/leo/common/sortx"
)

func TestIntsDesc(t *testing.T) {
	arr := []int{4, 2, 6, 3, 8, 9, 1}
	sortx.IntsDesc(arr)
	t.Log(arr)
}

func TestIntsAsc(t *testing.T) {
	arr := []int{4, 2, 6, 3, 8, 9, 1}
	sortx.IntsAsc(arr)
	t.Log(arr)
}
