package herr

import (
	"errors"
	"fmt"
	"runtime"
)

const depth = 32

// New create error with message.
func New(message string) error {
	var pcs = make([]uintptr, depth)
	n := runtime.Callers(2, pcs)
	var st stack = pcs[0:n]
	return &TracedError{error: errors.New(message), stack: st, mc: nil}
}

// Format create error with format and args.
func Format(format string, args ...any) error {
	var pcs = make([]uintptr, depth)
	n := runtime.Callers(2, pcs)
	var st stack = pcs[0:n]
	return &TracedError{error: fmt.Errorf(format, args...), stack: st, mc: nil}
}

// With wrap err as [TracedError] with formatted message.
//
// With doesn't reuse underlying TracedError if exists.
func With(err error, format string, args ...any) error {
	var pcs = make([]uintptr, depth)
	n := runtime.Callers(2, pcs)
	var st stack = pcs[0:n]
	return &TracedError{error: err, stack: st, mc: &mc{msg: fmt.Sprintf(format, args...)}}
}

// WithMsg wrap err as [TracedError] with message.
//
// WithMsg doesn't reuse underlying TracedError if exists.
func WithMsg(err error, message string) error {
	var pcs = make([]uintptr, depth)
	n := runtime.Callers(2, pcs)
	var st stack = pcs[0:n]
	return &TracedError{error: err, stack: st, mc: &mc{msg: message}}
}

// Wrap construct stack [TracedError] by err and message.
//
//	Wrap try to use existing stack and err.
//
// If the err is already a [TracedError], it mutate its message and return the original err.
func Wrap(err error, message string) error {
	if ws, ok := err.(*TracedError); ok {
		ws.mc = &mc{msg: message, parent: ws.mc}
		return ws
	}
	var pcs = make([]uintptr, depth)
	n := runtime.Callers(2, pcs)
	var st stack = pcs[0:n]
	return &TracedError{error: err, stack: st, mc: &mc{msg: message}}
}

// Wrapf wraps an error into a [TracedError], appending a formatted message.
//
//	Wrapf try to use existing stack and err.
//
// If the err is already a [TracedError], it mutate its message and return the original err.
func Wrapf(err error, format string, args ...any) error {
	if ws, ok := err.(*TracedError); ok {
		ws.mc = &mc{msg: fmt.Sprintf(format, args...), parent: ws.mc}
		return ws
	}
	var pcs = make([]uintptr, depth)
	n := runtime.Callers(2, pcs)
	var st stack = pcs[0:n]
	return &TracedError{error: err, stack: st, mc: &mc{msg: fmt.Sprintf(format, args...)}}
}
