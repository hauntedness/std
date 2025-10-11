package hv

import (
	"errors"
	"fmt"
)

// Err wraps an error into a [Error], recording the stack trace and a default message.
func Err(err error) error {
	return &Error{error: err, stack: callers(), msg: "..."}
}

// Errorx formats according to a format specifier and returns the string
// as a value that satisfies error.
// Errorx also records the stack trace at the point it was called.
func Errorx(format string, args ...any) error {
	return &Error{error: fmt.Errorf(format, args...), stack: callers(), msg: "..."}
}

// Errorf wraps an error into a [Error], appending a formatted message.
// If the error is already a [Error], it appends the message to the existing one.
func Errorf(err error, format string, args ...any) error {
	var ws = &Error{}
	if errors.As(err, &ws) {
		ws.msg = ws.msg + ": " + fmt.Sprintf(format, args...)
		return ws
	}
	ws.error = err
	ws.stack = callers()
	ws.msg = fmt.Sprintf(format, args...)
	return ws
}
