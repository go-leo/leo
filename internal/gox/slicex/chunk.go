package slicex

func Chunk[S ~[]E, E any](s S, size int) []S {
	l := len(s)
	ss2 := make([]S, 0, (l+size)/size)
	for i := 0; i < l; i += size {
		if i+size < l {
			ss2 = append(ss2, s[i:i+size])
		} else {
			ss2 = append(ss2, s[i:l])
		}
	}
	return ss2
}
