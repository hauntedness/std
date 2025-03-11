package hs

func Eq[T comparable](a, b T) bool {
	return a == b
}

// Loc return a slice from values, it accept negative index for start and end.
//
//	Loc([1,2,3], 1, -1) --> return [2]
//
// if end is too large, it is equivlent to the len(values).
func Loc[T any](values []T, start int, end int) []T {
	length := len(values)
	if start < 0 {
		start = length + start
	}
	if end <= 0 {
		end = length + end
	}
	if end > length {
		end = length
	}
	return values[start:end]
}

func At[T any](values []T, at int) T {
	if at < 0 {
		return values[len(values)+at]
	}
	return values[at]
}
