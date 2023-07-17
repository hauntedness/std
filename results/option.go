package results

import (
	"errors"
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

func Ok[T any](value T) Result[T] {
	return Result[T]{
		value: value,
		err:   nil,
	}
}

func Err[T any](err error) Result[T] {
	if err == nil {
		panic(ErrResultMustHasError)
	}
	var value T
	return Result[T]{
		value: value,
		err:   err,
	}
}

var (
	ErrResultMustHasError   = errors.New("result must have err")
	ErrResultMustHasNoError = errors.New("result must have no err")
)

func (r Result[T]) Get() T {
	if r.err != nil {
		panic(ErrResultMustHasNoError)
	}
	return r.value
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
