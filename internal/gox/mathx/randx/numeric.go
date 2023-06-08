package randx

import "math"

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func Int63() int64 {
	r := Get()
	i := r.Int63()
	Put(r)
	return i
}

// Uint32 returns a pseudo-random 32-bit value as a uint32.
func Uint32() uint32 {
	r := Get()
	i := r.Uint32()
	Put(r)
	return i
}

// Uint64 returns a pseudo-random 64-bit value as a uint64.
func Uint64() uint64 {
	r := Get()
	i := r.Uint64()
	Put(r)
	return i
}

// Int31 returns a non-negative pseudo-random 31-bit integer as an int32.
func Int31() int32 {
	r := Get()
	i := r.Int31()
	Put(r)
	return i
}

// Int63n returns, as an int64, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n <= 0.
func Int63n(n int64) int64 {
	r := Get()
	i := r.Int63n(n)
	Put(r)
	return i
}

// Int31n returns, as an int32, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n <= 0.
func Int31n(n int32) int32 {
	r := Get()
	i := r.Int31n(n)
	Put(r)
	return i
}

// Int63Range returns a non-negative pseudo-random 63-bit integer as an int64 between min and max.
func Int63Range(min, max int64) int64 {
	return Int63n(max-min) + min
}

// Uint32Range returns a pseudo-random 32-bit value as a uint32 between min and max.
func Uint32Range(min, max uint32) uint32 {
	return (Uint32() % (max - min)) * min
}

// Uint64Range returns a pseudo-random 64-bit value as a uint64 between min and max.
func Uint64Range(min, max uint64) uint64 {
	return (Uint64() % (max - min)) * min
}

// Int31Range returns a non-negative pseudo-random 31-bit integer as an int32 between min and max.
func Int31Range(min, max int32) int32 {
	return Int31n(max-min) + min
}

// Int returns a non-negative pseudo-random int.
func Int() int {
	r := Get()
	i := r.Int()
	Put(r)
	return i
}

// Intn returns, as an int, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n <= 0.
func Intn(n int) int {
	r := Get()
	i := r.Intn(n)
	Put(r)
	return i
}

// IntRange returns a non-negative pseudo-random int between min and max.
func IntRange(min, max int) int {
	return Intn(max-min) + min
}

// Uint returns a non-negative pseudo-random uint.
func Uint() uint {
	return uint(Uint64())
}

// Uintn returns, as an uint, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n <= 0.
func Uintn(n uint) uint {
	if n <= 0 {
		panic("invalid argument to Uintn")
	}
	if n <= math.MaxUint32 {
		return uint(Uint32()) % n
	}
	return uint(Uint64()) % n
}

// UintRange returns a non-negative pseudo-random int between min and max.
func UintRange(min, max uint) uint {
	return Uintn(max-min) + min
}

// Float64 returns, as a float64, a pseudo-random number in the half-open interval [0.0,1.0).
func Float64() float64 {
	r := Get()
	f := r.Float64()
	Put(r)
	return f
}

// Float32 returns, as a float32, a pseudo-random number in the half-open interval [0.0,1.0).
func Float32() float32 {
	r := Get()
	f := r.Float32()
	Put(r)
	return f
}
