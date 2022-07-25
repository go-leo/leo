package randomx

func Intn(n int) int {
	return r.Intn(n)
}

func PickItemsInt32(size int, n ...int32) []int32 {
	if size <= 0 {
		return []int32{}
	}
	copy := append([]int32{}, n...)
	if size >= len(n) {
		return copy
	}
	indexSet := make(map[int]struct{})
	for i := 0; i < size; i++ {
		index := Intn(len(n))
		if _, ok := indexSet[index]; !ok {
			indexSet[index] = struct{}{}
			continue
		}
		for {
			index = Intn(len(n))
			if _, ok := indexSet[index]; !ok {
				indexSet[index] = struct{}{}
				break
			}
		}

	}

	return copy[:size]
}
