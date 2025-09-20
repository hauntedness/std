package hv

func Ptr[T any](v T) *T {
	return &v
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
