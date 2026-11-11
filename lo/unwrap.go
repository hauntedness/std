package lo

import "errors"

// Unwrap try to unwrap the err in [*Error].
func Unwrap(err error) error {
	te := &ErrorChain{}
	if errors.As(err, &te) {
		return te.error
	}

	return err
}

func As[T error](err error) (T, bool) {
	return errors.AsType[T](err)
}

func Is(err error, target error) bool {
	return errors.Is(err, target)
}

// Match check whether the error or any error in its tree matches your specified function.
func Match(err error, fn func(err error) bool) bool {
	if err == nil || fn == nil {
		return false
	}

	type Unwrapable = interface {
		error
		Unwrap() error
	}

	type Unwrapable2 = interface {
		error
		Unwrap() []error
	}

	for {
		if fn(err) {
			return true
		}

		if x, ok := errors.AsType[Unwrapable](err); ok {
			err = x.Unwrap()
			if err == nil {
				return false
			}
		} else if x1, ok := errors.AsType[Unwrapable2](err); ok {
			for _, err := range x1.Unwrap() {
				if Match(err, fn) {
					return true
				}
			}

			return false
		} else {
			return false
		}
	}
}
