package options

import (
	"errors"
)

var (
	ErrNotPresent = errors.New("value not present")
	ErrNotOk      = errors.New("value not ok")
)

type Option[T any] struct {
	value     T
	isPresent bool
}

func From[T any](value T, isPresent bool) Option[T] {
	return Option[T]{
		value:     value,
		isPresent: isPresent,
	}
}

func FromPointer[T any](pointer *T) Option[T] {
	if pointer == nil {
		return Empty[T]()
	}
	return Option[T]{
		value:     *pointer,
		isPresent: true,
	}
}

func OrElse[T any](ok bool, value T, other T) T {
	if ok {
		return value
	}
	return other
}

func Must[T any](value T, ok bool) T {
	if ok {
		return value
	}
	panic(ErrNotOk)
}

func Some[T any](value T) Option[T] {
	return Option[T]{
		value:     value,
		isPresent: true,
	}
}

func Empty[T any]() Option[T] {
	var value T
	return Option[T]{
		value:     value,
		isPresent: false,
	}
}

func (o Option[T]) Get() T {
	if !o.isPresent {
		panic(ErrNotPresent)
	}
	return o.value
}

func (o Option[T]) OrElse(other T) T {
	if !o.isPresent {
		return other
	}
	return o.value
}

func (o Option[T]) OrEmpty() T {
	if !o.isPresent {
		var other T
		return other
	}
	return o.value
}

func (o Option[T]) IsPresent() bool {
	return o.isPresent
}
