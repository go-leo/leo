package slicex

import "codeup.aliyun.com/qimao/leo/leo/internal/gox/mathx"

func Shift[S ~[]E, E any](array S, offset int) {
	if len(array) <= 0 {
		return
	}
	Shifta(array, 0, len(array), offset)
}

func Shifta[S ~[]E, E any](array S, startIndexInclusive, endIndexExclusive, offset int) {
	if len(array) <= 0 || startIndexInclusive >= len(array)-1 || endIndexExclusive <= 0 {
		return
	}
	if startIndexInclusive < 0 {
		startIndexInclusive = 0
	}
	if endIndexExclusive >= len(array) {
		endIndexExclusive = len(array)
	}
	n := endIndexExclusive - startIndexInclusive
	if n <= 1 {
		return
	}
	offset %= n
	if offset < 0 {
		offset += n
	}
	// For algorithm explanations and proof of O(n) time complexity and O(1) space complexity
	// see https://beradrian.wordpress.com/2015/04/07/shift-an-array-in-on-in-place/
	for n > 1 && offset > 0 {
		nOffset := n - offset

		if offset > nOffset {
			Swap(array, startIndexInclusive, startIndexInclusive+n-nOffset, nOffset)
			n = offset
			offset -= nOffset
		} else if offset < nOffset {
			Swap(array, startIndexInclusive, startIndexInclusive+nOffset, offset)
			startIndexInclusive += offset
			n = nOffset
		} else {
			Swap(array, startIndexInclusive, startIndexInclusive+nOffset, offset)
			break
		}
	}

}

// Swap swaps a series of elements in the given array.
// array the array to swap.
// offset1 the index of the first element in the series to swap.
// offset2 the index of the second element in the series to swap.
// length the number of elements to swap starting with the given indices.
func Swap[S ~[]E, E any](array S, offset1, offset2, length int) {
	if IsEmpty(array) || offset1 >= len(array) || offset2 >= len(array) {
		return
	}
	if offset1 < 0 {
		offset1 = 0
	}
	if offset2 < 0 {
		offset2 = 0
	}
	length = mathx.Min(mathx.Min(length, len(array)-offset1), len(array)-offset2)
	for i := 0; i < length; i++ {
		aux := array[offset1]
		array[offset1] = array[offset2]
		array[offset2] = aux
		offset1++
		offset2++
	}
}
