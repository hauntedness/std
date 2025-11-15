package herr

import (
	"fmt"
)

// New create error with ErrNew and message, it's same to With(ErrNew, message).
func New(message string) error {
	return &TracedError{error: ErrNew, stack: callers(), msg: message}
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
