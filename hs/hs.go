// package hs present some helper type and func for slices.
package hs

func As[T any](values ...T) []T {
	return values
}

func Eq[T comparable](a, b T) bool {
	return a == b
}

func EqTo[T comparable](a T) func(T) bool {
	return func(t T) bool {
		return a == t
	}
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
	if end < 0 {
		end = length + end
	} else if end > length {
		end = length
	}
	return values[start:end]
}

// Tail return the slice from start to len(values) of values, it accept negative index for start.
//
//	Loc([1,2,3], -1) --> return [2]
func Tail[T any](values []T, start int) []T {
	length := len(values)
	if start < 0 {
		start = length + start
	}
	return values[start:length]
}

func At[T any](values []T, at int) T {
	if at < 0 {
		return values[len(values)+at]
	}
	return values[at]
}
