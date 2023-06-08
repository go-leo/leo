package slicex

import "golang.org/x/exp/slices"

func Remove[S ~[]E, E comparable](s S, v E) S {
	if IsEmpty(s) {
		return slices.Clone(s)
	}
	return Delete(s, slices.Index(s, v))
}

func RemoveFunc[S ~[]E, E any](s S, f func(E) bool) S {
	if IsEmpty(s) {
		return slices.Clone(s)
	}
	return Delete(s, slices.IndexFunc(s, f))
}

func RemoveAll[S ~[]E, E comparable](s S, vs ...E) S {
	if IsEmpty(s) || IsEmpty(vs) {
		return slices.Clone(s)
	}
	occurrences := make(map[E]int)
	total := 0
	for _, v := range vs {
		total++
		count, ok := occurrences[v]
		if !ok {
			occurrences[v] = 1
		} else {
			occurrences[v] = count + 1
		}
	}
	toRemove := make([]int, 0, total)
	for i := 0; i < len(s); i++ {
		key := s[i]
		count, ok := occurrences[key]
		if ok {
			count--
			if count == 0 {
				delete(occurrences, key)
			}
			toRemove = append(toRemove, i)
		}
	}
	return DeleteAll(s, toRemove...)
}

func RemoveAllFunc[S ~[]E, E comparable](s S, f func(E) bool) S {
	if IsEmpty(s) {
		return slices.Clone(s)
	}
	return DeleteAll(s, IndexesFunc(s, f)...)
}
