package herr

import "errors"

// Unwrap try to unwrap the err in [*TracedError].
func Unwrap(err error) error {
	te := &TracedError{}
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

	for {
		if fn(err) {
			return true
		}

		{
			var (
				x  interface{ Unwrap() error }
				x1 interface{ Unwrap() []error }
			)
			switch {
			case errors.As(err, &x):
				err = x.Unwrap()
				if err == nil {
					return false
				}
			case errors.As(err, &x1):
				for _, err := range x1.Unwrap() {
					if Match(err, fn) {
						return true
					}
				}

				return false
			default:
				return false
			}
		}
	}
}
