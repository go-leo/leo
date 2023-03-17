package operator

func Ternary[T any](condition bool, exprIfTrue, exprIfFalse T) T {
	if condition {
		return exprIfTrue
	}
	return exprIfFalse
}
