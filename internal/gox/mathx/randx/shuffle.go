package randx

func Shuffle[S []E, E any](s S) S {
	r := Get()
	r.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
	Put(r)
	return s
}
