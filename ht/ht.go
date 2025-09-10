// package ht provide some predeclared types alias or constructor functions for convience
package ht

type Unit = struct{}

type Dict = map[string]any

func Ptr[T any](v T) *T {
	return &v
}

func Zero[T any]() T {
	return *(new(T))
}
