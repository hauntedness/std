package hv

//go:fix inline
func Ptr[T any](v T) *T {
	return new(v)
}

func Zero[T any]() T {
	var t T
	return t
}

func OrZero[T any](v *T) T {
	if v == nil {
		return *new(T)
	}

	return *v
}

func OrElse[T any](v *T, supply func() *T) *T {
	if v == nil {
		return supply()
	}

	return v
}
