package hv

import (
	"fmt"
)

// Error construct stack [TracedError] by err and message.
func Error(err error, message string) error {
	return &TracedError{error: err, stack: callers(), msg: message}
}

// Err formats according to a format specifier and returns the string
// as a value that satisfies error.
// Err also records the stack trace at the point it was called.
func Err(format string, args ...any) error {
	return &TracedError{error: fmt.Errorf(format, args...), stack: callers(), msg: "..."}
}

// Errf wraps an error into a [TracedError], appending a formatted message.
// If the error is already a [TracedError], it appends the message to the existing one.
func Errf(err error, format string, args ...any) error {
	if ws, ok := err.(*TracedError); ok {
		ws.msg = ws.msg + ": " + fmt.Sprintf(format, args...)
		return ws
	}
	var ws = &TracedError{}
	ws.error = err
	ws.stack = callers()
	ws.msg = fmt.Sprintf(format, args...)
	return ws
}
