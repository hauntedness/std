package hs

import "slices"

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

// Map apply fn on each element of values []T and return the transformed []R.
func Map[T any, R any](values []T, fn func(T) R) []R {
	r := make([]R, len(values))
	for i := range values {
		r[i] = fn(values[i])
	}
	return r
}

// Pipe is filter + map from []T to []R.
//
// Pipe does NOT removes unused capacity, call [slices.Clip] by yourself if needed.
func Pipe[T any, R any](values []T, fn func(T) (R, bool)) []R {
	r := make([]R, 0, len(values))
	for i := range values {
		if v, ok := fn(values[i]); ok {
			r = append(r, v)
		}
	}
	return r
}

// PipeVec is filter + map from *Vec[T] to *Vec[R].
//
// PipeVec does NOT removes unused capacity, call [Clip] by yourself if needed.
func PipeVec[T any, R any](values *Vec[T], fn func(T) (R, bool)) *Vec[R] {
	r := make([]R, 0, len(values.data))
	for i := range values.data {
		if v, ok := fn(values.data[i]); ok {
			r = append(r, v)
		}
	}
	return &Vec[R]{data: r}
}

// Contains Contains reports whether value is present in values.
func Contains[T comparable](value T, values ...T) bool {
	return slices.Contains(values, value)
}

func Distinct[T comparable](v []T) []T {
	if len(v) == 0 {
		return nil
	}
	// Use a map to track seen elements
	seen := make(map[T]struct{})
	res := make([]T, 0, len(v))

	for _, val := range v {
		if _, ok := seen[val]; !ok {
			seen[val] = struct{}{}
			res = append(res, val)
		}
	}

	return res
}
