package close

import "errors"

type Closable interface {
	Close() error
}

// CloseError discriminate whether the error comes from close action
type CloseError struct {
	err error
}

func (c *CloseError) Error() string {
	return c.err.Error()
}

func (c *CloseError) Unwrap() error {
	return c.err
}

func AfterFunc[T Closable](t T, f func(t T) error) (err error) {
	defer func() {
		closeerr := t.Close()
		if closeerr != nil {
			err = errors.Join(&CloseError{err: closeerr}, err)
		}
	}()
	err = f(t)
	return
}

func AfterValueFunc[T Closable, V any](t T, f func(t T) (V, error)) (value V, err error) {
	defer func() {
		closeerr := t.Close()
		if closeerr != nil {
			err = errors.Join(&CloseError{err: closeerr}, err)
		}
	}()
	value, err = f(t)
	return
}
