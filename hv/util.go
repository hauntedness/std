package hv

func Pointr[T any](v T) *T {
	return &v
}

func OrEmpty[T any](v *T) T {
	if v == nil {
		return *new(T)
	}
	return *v
}

func OrElse[T any](v *T, fallback T) T {
	if v == nil {
		return fallback
	}
	return *v
}

func OrZero[T any](value T, ok bool) T {
	if ok {
		return value
	}
	return *new(T)
}
