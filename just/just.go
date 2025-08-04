package just

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

// TryGo spawn a new goroutine and recover panic from f and send the recovered value or nil to error chan.
func TryGo(f func()) <-chan error {
	ch := make(chan error, 1)
	go func() {
		defer func() {
			if v := recover(); v != nil {
				if e, ok := v.(error); ok {
					ch <- e
				} else {
					ch <- &PanicError{v}
				}
			} else {
				ch <- nil
			}
			close(ch)
		}()
		f()
	}()
	return ch
}

// Want want to get value with no error or else panic.
func Want[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

// Check check no error or else panic.
func Check(err error) {
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
