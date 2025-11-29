package herr

import (
	"errors"
	"fmt"
)

// New create error with message.
func New(message string) error {
	return &TracedError{error: errors.New(message), stack: callers(), msg: message}
}

// Format create error with format and args.
func Format(format string, args ...any) error {
	return &TracedError{error: fmt.Errorf(format, args...), stack: callers(), msg: ""}
}

// With wrap err as [TracedError] with no msg.
func With(err error, message string) error {
	return &TracedError{error: err, stack: callers(), msg: message}
}

// Wrap construct stack [TracedError] by err and message.
//  Wrap try to use existing stack. If the err is already a [TracedError], it appends the message to the existing one.
func Wrap(err error, message string) error {
	if ws, ok := err.(*TracedError); ok {
		ws.msg = ws.msg + ": " + message
		return ws
	}
	return &TracedError{error: err, stack: callers(), msg: message}
}

// Wrapf wraps an error into a [TracedError], appending a formatted message.
//  Wrapf try to use existing stack. If the err is already a [TracedError], it appends the message to the existing one.
func Wrapf(err error, format string, args ...any) error {
	if ws, ok := err.(*TracedError); ok {
		ws.msg = ws.msg + ": " + fmt.Sprintf(format, args...)
		return ws
	}
	return &TracedError{error: err, stack: callers(), msg: fmt.Sprintf(format, args...)}
}
