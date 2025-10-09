package he

import (
	"errors"
	"fmt"
)

// Error wraps an error into a StructError, recording the stack trace and a default message.
func Error(err error) error {
	return &StructError{error: err, stack: callers(), msg: "..."}
}

// Err formats according to a format specifier and returns the string
// as a value that satisfies error.
// Err also records the stack trace at the point it was called.
func Err(format string, args ...any) error {
	return &StructError{error: fmt.Errorf(format, args...), stack: callers(), msg: "..."}
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
