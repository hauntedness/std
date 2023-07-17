package options

import (
	"errors"
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

var ErrNotPresent = errors.New("value not present")

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
