package herr

import (
	"fmt"
)

// Try recover panic from f and save the recovered value to error.
func Try[T any](f func() T) (t T, err error) {
	defer func() {
		if v := recover(); v != nil {
			if e, ok := v.(error); ok {
				err = e
			} else {
				err = &PanicError{v}
			}
		}
	}()

	t = f()

	return t, err
}

// TryDo is similar to [Try] but for function without 0 result.
func TryDo(f func()) (err error) {
	defer func() {
		if v := recover(); v != nil {
			if e, ok := v.(error); ok {
				err = e
			} else {
				err = &PanicError{v}
			}
		}
	}()

	f()

	return err
}

// Panicf panic with formated error instead of interface{}.
func Panicf(format string, a ...any) {
	panic(fmt.Errorf(format, a...))
}

// Must Must get value when no error or else panic.
func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}

	return value
}

// MustNil check no error or else panic.
func MustNil(err error) {
	if err != nil {
		panic(err)
	}
}

type PanicError struct {
	Value any
}

func (p *PanicError) Error() string {
	return fmt.Sprintf("panic: %v", p.Value)
}
