package fmts

import (
	"fmt"
)

func Error(err error) error {
	return &withStack{error: err, stack: callers()}
}

// Errorf formats according to a format specifier and returns the string
// as a value that satisfies error.
// Errorf also records the stack trace at the point it was called.
func Errorf(format string, args ...any) error {
	return &withStack{error: fmt.Errorf(format, args...), stack: callers()}
}
