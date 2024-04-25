package closer

import "errors"

type Closable interface {
	Close() error
}

func AfterFunc[T Closable](t Closable, f func(t Closable) error) (err error) {
	defer func() {
		closeerr := t.Close()
		if closeerr != nil {
			err = errors.Join(closeerr, err)
		}
	}()
	err = f(t)
	return
}

func AfterValueFunc[T Closable, V any](t T, f func(t T) (V, error)) (value V, err error) {
	defer func() {
		closeerr := t.Close()
		if closeerr != nil {
			err = errors.Join(closeerr, err)
		}
	}()
	value, err = f(t)
	return
}
