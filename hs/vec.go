package hs

import (
	"fmt"
	"iter"
	"slices"
)

// Vec is simpler slice, mostly you don't need this.
//
//	Use *Vec instead of Vec
//
// The zero value is hard to use thus...
type Vec[T any] struct {
	data []T
}

func Make[T any](len, cap int) *Vec[T] {
	return &Vec[T]{data: make([]T, len, cap)}
}

func New[T any](data []T) *Vec[T] {
	return &Vec[T]{data: data}
}

func NewWith[T any](data ...T) *Vec[T] {
	return &Vec[T]{data: data}
}

func Repeat[T any](v T, count int) *Vec[T] {
	return &Vec[T]{data: slices.Repeat([]T{v}, count)}
}

func (v *Vec[T]) Seq() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := range v.data {
			if !yield(v.data[i]) {
				return
			}
		}
	}
}

func (v *Vec[T]) Reduce(initial T, fn func(a, b T) T) T {
	for i := range v.data {
		initial = fn(initial, v.data[i])
	}
	return initial
}

func (v *Vec[T]) Data() []T {
	return v.data
}

func (v *Vec[T]) Append(data ...T) *Vec[T] {
	v.data = slices.Concat(v.data, data)
	return v
}

// Pipe create a new Vec, the new element depends on the results of fn
func (v *Vec[T]) Pipe(fn func(T) (T, bool)) *Vec[T] {
	res := make([]T, 0, len(v.data))
	for i := range v.data {
		if nv, ok := fn(v.data[i]); ok {
			res = append(res, nv)
		}
	}
	return New(res)
}

func (v *Vec[T]) Clone() *Vec[T] {
	return New(slices.Clone(v.data))
}

// Slice is just an alias to Loc.
//
// Deprecated: use Loc as it is shorter.
func (v *Vec[T]) Slice(start, end int) *Vec[T] {
	return v.Loc(start, end)
}

// Loc return the sub slice from the original vec.
func (v *Vec[T]) Loc(start, end int) *Vec[T] {
	return New(Loc(v.data, start, end))
}

// Equal compare each element, and return true if all the same.
// use hs.Eq for convenience.
func (v *Vec[T]) Equal(other *Vec[T], eq func(a T, b T) bool) bool {
	return slices.EqualFunc(v.data, other.data, eq)
}

// Contains reports whether at least one element elem of v satisfies eq(elem, input).
// use hs.Eq for convenience.
func (v *Vec[T]) Contains(fn func(elem T) bool) bool {
	return slices.ContainsFunc(v.data, fn)
}

// Index IndexFunc returns the first index i satisfying eq(elem, input), or -1 if none do.
// use hs.Eq for convenience.
func (v *Vec[T]) Index(fn func(elem T) bool) int {
	return slices.IndexFunc(v.data, fn)
}

// Sort sorts the slice x in ascending order as determined by the cmp function.
func (v *Vec[T]) Sort(cmp func(a T, b T) int) *Vec[T] {
	slices.SortFunc(v.data, cmp)
	return v
}

// Reverse reverses the elements of the slice in place.
func (v *Vec[T]) Reverse() *Vec[T] {
	slices.Reverse(v.data)
	return v
}

func (v *Vec[T]) IsSorted(cmp func(a T, b T) int) bool {
	return slices.IsSortedFunc(v.data, cmp)
}

// BinarySearch searches for target in a sorted slice and returns the earliest position where target is found.
//
// For more detail see: [slices.BinarySearch]
func (v *Vec[T]) BinarySearch(target T, cmp func(a, b T) int) (pos int, ok bool) {
	return slices.BinarySearchFunc(v.data, target, cmp)
}

func (v *Vec[T]) Len() int {
	return len(v.data)
}

func (v *Vec[T]) Set(i int, value T) {
	v.data[i] = value
}

func (v *Vec[T]) Get(index int) T {
	return v.data[index]
}

// At is similar to Get but accept negative index.
// -1 will locate the last element.
func (v *Vec[T]) At(index int) T {
	return At(v.data, index)
}

func (v *Vec[T]) String() string {
	return fmt.Sprint(v.data)
}
