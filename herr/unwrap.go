package herr

import "errors"

// Unwrap try to unwrap the err in [*TracedError].
func Unwrap(err error) error {
	te, ok := err.(*TracedError)
	if ok {
		return te.error
	}

	te = &TracedError{}
	if errors.As(err, &te) {
		return te.error
	}

	return err
}
