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
type Vec[T any] []T

func Make[T any](len, cap int) *Vec[T] {
	ret := Vec[T](make([]T, len, cap))
	return &ret
}

func New[T any](data []T) *Vec[T] {
	ret := Vec[T](data)
	return &ret
}

func NewWith[T any](data ...T) *Vec[T] {
	ret := Vec[T](data)
	return &ret
}

func Repeat[T any](v T, count int) *Vec[T] {
	ret := Vec[T](slices.Repeat([]T{v}, count))
	return &ret
}

func (v *Vec[T]) Seq() iter.Seq[T] {
	return func(yield func(T) bool) {
		data := *v
		for i := range data {
			if !yield(data[i]) {
				return
			}
		}
	}
}

func (v *Vec[T]) Reduce(initial T, fn func(a, b T) T) T {
	data := *v
	for i := range data {
		initial = fn(initial, data[i])
	}
	return initial
}

func (v *Vec[T]) Data() []T {
	return *v
}

func (v *Vec[T]) Append(data ...T) *Vec[T] {
	*v = slices.Concat(*v, data)
	return v
}

// Pipe create a new Vec, the new element depends on the results of fn
func (v *Vec[T]) Pipe(fn func(T) (T, bool)) *Vec[T] {
	data := *v
	res := make([]T, 0, len(data))
	for i := range data {
		if nv, ok := fn(data[i]); ok {
			res = append(res, nv)
		}
	}
	return New(res)
}

func (v *Vec[T]) Clone() *Vec[T] {
	return New(slices.Clone(*v))
}

// Slice is just an alias to Loc.
//
// Deprecated: use Loc as it is shorter.
func (v *Vec[T]) Slice(start, end int) *Vec[T] {
	return v.Loc(start, end)
}

// Loc return the sub slice from the original vec.
func (v *Vec[T]) Loc(start, end int) *Vec[T] {
	return New(Loc(*v, start, end))
}

// Equal compare each element, and return true if all the same.
// use hs.Eq for convenience.
func (v *Vec[T]) Equal(other *Vec[T], eq func(a T, b T) bool) bool {
	return slices.EqualFunc(*v, *other, eq)
}

// Contains reports whether at least one element elem of v satisfies eq(elem, input).
// use hs.Eq for convenience.
func (v *Vec[T]) Contains(input T, eq func(elem T, input T) bool) bool {
	return slices.ContainsFunc(*v, func(elem T) bool {
		return eq(elem, input)
	})
}

// Index IndexFunc returns the first index i satisfying eq(elem, input), or -1 if none do.
// use hs.Eq for convenience.
func (v *Vec[T]) Index(input T, eq func(elem T, input T) bool) int {
	return slices.IndexFunc(*v, func(elem T) bool {
		return eq(elem, input)
	})
}

// Sort sorts the slice x in ascending order as determined by the cmp function.
func (v *Vec[T]) Sort(cmp func(a T, b T) int) *Vec[T] {
	slices.SortFunc(*v, cmp)
	return v
}

// Reverse reverses the elements of the slice in place.
func (v *Vec[T]) Reverse() *Vec[T] {
	slices.Reverse(*v)
	return v
}

func (v *Vec[T]) IsSorted(cmp func(a T, b T) int) bool {
	return slices.IsSortedFunc(*v, cmp)
}

// BinarySearch searches for target in a sorted slice and returns the earliest position where target is found.
//
// For more detail see: [slices.BinarySearch]
func (v *Vec[T]) BinarySearch(target T, cmp func(a, b T) int) (pos int, ok bool) {
	return slices.BinarySearchFunc(*v, target, cmp)
}

func (v *Vec[T]) Len() int {
	return len(*v)
}

func (v *Vec[T]) Set(i int, value T) {
	(*v)[i] = value
}

func (v *Vec[T]) Get(index int) T {
	return (*v)[index]
}

// At is similar to Get but accept negative index.
// -1 will locate the last element.
func (v *Vec[T]) At(index int) T {
	return At(*v, index)
}

func (v *Vec[T]) String() string {
	return fmt.Sprint((*[]T)(v))
}
