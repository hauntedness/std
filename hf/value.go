// Package hf provide functions for
package hf

// Supply is used as anonymous func to return the value.
func Supply[T any](value T) func() T {
	return func() T {
		return value
	}
}
