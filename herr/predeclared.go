package herr

import "errors"

// ErrInvalid indicates that a value or operation is invalid.
//
// Functions and methods should not return this error but should instead return a wrap of this one.
var ErrInvalid = errors.New("value is invalid")

// ErrUnexpected indicates that the error is unexpected but happened. It's useful for debugging.
//
// Functions and methods should not return this error but should instead return a wrap of this one.
var ErrUnexpected = errors.New("unexpected error")
