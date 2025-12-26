package hs

import "iter"

func (v *Vec[T]) Seq() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := range v.data {
			if !yield(v.data[i]) {
				return
			}
		}
	}
}

func (v *Vec[T]) SeqPT() iter.Seq[*T] {
	return func(yield func(*T) bool) {
		for i := range v.data {
			if !yield(&v.data[i]) {
				return
			}
		}
	}
}

func (v *Vec[T]) Seq2(start, end int) iter.Seq2[int, T] {
	length := len(v.data)
	if start < 0 {
		start = length + start
	}
	if end < 0 {
		end = length + end
	} else if end > length {
		end = length
	}
	return func(yield func(int, T) bool) {
		for i := start; i < end; i++ {
			if !yield(i, v.data[i]) {
				return
			}
		}
	}
}

func (v *Vec[T]) SeqPT2(start, end int) iter.Seq2[int, *T] {
	length := len(v.data)
	if start < 0 {
		start = length + start
	}
	if end < 0 {
		end = length + end
	} else if end > length {
		end = length
	}
	return func(yield func(int, *T) bool) {
		for i := start; i < end; i++ {
			if !yield(i, &v.data[i]) {
				return
			}
		}
	}
}

func Seq2[T any](data []T, start, end int) iter.Seq2[int, T] {
	length := len(data)
	if start < 0 {
		start = length + start
	}
	if end < 0 {
		end = length + end
	} else if end > length {
		end = length
	}
	return func(yield func(int, T) bool) {
		for i := start; i < end; i++ {
			if !yield(i, data[i]) {
				return
			}
		}
	}
}

func SeqPT2[T any](data []T, start, end int) iter.Seq2[int, *T] {
	length := len(data)
	if start < 0 {
		start = length + start
	}
	if end < 0 {
		end = length + end
	} else if end > length {
		end = length
	}
	return func(yield func(int, *T) bool) {
		for i := start; i < end; i++ {
			if !yield(i, &data[i]) {
				return
			}
		}
	}
}
