package slicex

import (
	"strconv"
)

// JoinInt concatenates the slice of int to create a single string. The separator
// string sep is placed between elements in the resulting string.
func JoinInt(elems []int, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return strconv.Itoa(elems[0])
	default:
		var bs []byte
		bs = strconv.AppendInt(bs, int64(elems[0]), 10)
		for _, elem := range elems[1:] {
			bs = append(bs, sep...)
			bs = strconv.AppendInt(bs, int64(elem), 10)
		}
		return string(bs)
	}
}

// JoinInt32 concatenates the slice of int32 to create a single string. The separator
// string sep is placed between elements in the resulting string.
func JoinInt32(elems []int32, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return strconv.Itoa(int(elems[0]))
	default:
		var bs []byte
		bs = strconv.AppendInt(bs, int64(elems[0]), 10)
		for _, elem := range elems[1:] {
			bs = append(bs, sep...)
			bs = strconv.AppendInt(bs, int64(elem), 10)
		}
		return string(bs)
	}
}

// JoinInt64 concatenates the slice of int64 to create a single string. The separator
// string sep is placed between elements in the resulting string
func JoinInt64(elems []int64, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return strconv.Itoa(int(elems[0]))
	default:
		var bs []byte
		bs = strconv.AppendInt(bs, elems[0], 10)
		for _, elem := range elems[1:] {
			bs = append(bs, sep...)
			bs = strconv.AppendInt(bs, elem, 10)
		}
		return string(bs)
	}
}
