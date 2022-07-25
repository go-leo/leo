package slicex

func Chunk[T any](slice []T, size int) [][]T {
	length := len(slice)
	chunked := make([][]T, 0, (length+size)/size)
	for i := 0; i < length; i += size {
		if i+size < length {
			chunked = append(chunked, slice[i:i+size])
		} else {
			chunked = append(chunked, slice[i:length])
		}
	}
	return chunked
}

// ChunkIndexes Split the slice into multiple size of blocks, and return all the split point indexes.
// If the slice cannot be split into all equally long blocks, the last remaining elements will form a block.
func ChunkIndexes(length, size int) []int {
	eps := make([]int, 0, (length+size)/size)
	for i := 0; i < length; i += size {
		if i%size == 0 {
			eps = append(eps, i)
		}
	}
	eps = append(eps, length)
	return eps
}
