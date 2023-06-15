package randx

// Intn returns, as an int, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n <= 0.
func Intn(n int) int {
	r := Get()
	i := r.Intn(n)
	Put(r)
	return i
}
