package fmts

import (
	"errors"
	"fmt"
)

// ErrAnonymouse represent errors that doesn't have source. see [Errorf].
var ErrAnonymouse = errors.New("error")

// Error wraps an error into a StructError, recording the stack trace and a default message.
func Error(err error) error {
	return &StructError{error: err, stack: callers(), msg: "no message"}
}

// Errors formats according to a format specifier and returns the string
// as a value that satisfies error.
// Errors also records the stack trace at the point it was called.
func Errors(format string, args ...any) error {
	return &StructError{error: ErrAnonymouse, msg: fmt.Sprintf(format, args...), stack: callers()}
}

// Errorf wraps an error into a StructError, appending a formatted message.
// If the error is already a StructError, it appends the message to the existing one.
func Errorf(err error, format string, args ...any) error {
	var ws = &StructError{}
	if errors.As(err, &ws) {
		ws.msg = ws.msg + ": " + fmt.Sprintf(format, args...)
		return ws
	}
	ws.error = err
	ws.stack = callers()
	ws.msg = fmt.Sprintf(format, args...)
	return ws
}
