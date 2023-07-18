package results

import (
	"errors"
)

var (
	ErrResultNilError = errors.New("error is nil")
	ErrResultHasError = errors.New("result has error")
)

type Result[T any] struct {
	value T
	err   error
}

func From[T any](value T, err error) Result[T] {
	return Result[T]{
		value: value,
		err:   err,
	}
}

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func Ok[T any](value T) Result[T] {
	return Result[T]{
		value: value,
		err:   nil,
	}
}

func Err[T any](err error) Result[T] {
	if err == nil {
		panic(ErrResultNilError)
	}
	var value T
	return Result[T]{
		value: value,
		err:   err,
	}
}

func (r Result[T]) Get() T {
	if r.err != nil {
		panic(ErrResultHasError)
	}
	return r.value
}

func (r Result[T]) Unpack() (T, error) {
	return r.value, r.err
}

func (r Result[T]) Err() error {
	return r.err
}

func (r Result[T]) OrElse(other T) T {
	if r.err != nil {
		return other
	}
	return r.value
}

func (r Result[T]) OrEmpty() T {
	if r.err != nil {
		var other T
		return other
	}
	return r.value
}

func (r Result[T]) IsOk() bool {
	return r.err == nil
}

func (r Result[T]) IsErr() bool {
	return r.err != nil
}
