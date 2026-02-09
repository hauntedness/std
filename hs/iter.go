package hs

import "iter"

// Seq returns an iterator over the values of the vector.
func (v *Vec[T]) Seq() iter.Seq[T] {
	return func(yield func(T) bool) {
		data := v.data
		for i := range data {
			if !yield(data[i]) {
				return
			}
		}
	}
}

// Seqr returns an iterator over pointers to the values in the vector.
func (v *Vec[T]) Seqr() iter.Seq[*T] {
	return func(yield func(*T) bool) {
		data := v.data
		for i := range data {
			if !yield(&data[i]) {
				return
			}
		}
	}
}

// Seq2 returns an iterator that yields both the index and the value.
func (v *Vec[T]) Seq2() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, value := range v.data {
			if !yield(i, value) {
				return
			}
		}
	}
}

// Seqr2 returns an iterator that yields the index and a pointer to the value.
func (v *Vec[T]) Seqr2() iter.Seq2[int, *T] {
	return func(yield func(int, *T) bool) {
		data := v.data
		for i := range data {
			if !yield(i, &data[i]) {
				return
			}
		}
	}
}

// Seq returns an iterator that yields each value from data.
func Seq[T any](data []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := range data {
			if !yield(data[i]) {
				return
			}
		}
	}
}

// Seq2 returns an iterator that yields both the index and the value from data.
func Seq2[T any](data []T) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i := range data {
			if !yield(i, data[i]) {
				return
			}
		}
	}
}

// Seqr2 returns an iterator that yields the index and a pointer to the value from data.
func Seqr2[T any](data []T) iter.Seq2[int, *T] {
	return func(yield func(int, *T) bool) {
		for i := range data {
			if !yield(i, &data[i]) {
				return
			}
		}
	}
}
