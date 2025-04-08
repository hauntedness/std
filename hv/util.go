package hv

func Ptr[T any](v T) *T {
	return &v
}

func OrZero[T any](v *T) T {
	if v == nil {
		return *new(T)
	}
	return *v
}
