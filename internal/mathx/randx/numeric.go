package randx

import "math/rand"

// Int implements rand.Int on the randx global source.
func Int() int {
	mu.Lock()
	defer mu.Unlock()
	rand.Int()
	return r.Int()
}

// Int63n implements rand.Int63n on the randx global source.
func Int63n(n int64) int64 {
	mu.Lock()
	defer mu.Unlock()
	return r.Int63n(n)
}

// Intn implements rand.Intn on the randx global source.
func Intn(n int) int {
	mu.Lock()
	defer mu.Unlock()
	return r.Intn(n)
}

// Int31n implements rand.Int31n on the randx global source.
func Int31n(n int32) int32 {
	mu.Lock()
	defer mu.Unlock()
	return r.Int31n(n)
}

// Float64 implements rand.Float64 on the randx global source.
func Float64() float64 {
	mu.Lock()
	defer mu.Unlock()
	return r.Float64()
}

// Uint64 implements rand.Uint64 on the randx global source.
func Uint64() uint64 {
	mu.Lock()
	defer mu.Unlock()
	return r.Uint64()
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
