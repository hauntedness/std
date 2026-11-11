package lo

import (
	"errors"
	"fmt"
	"runtime"
)

const depth = 32

// NewError create error with message.
func NewError(message string) error {
	var pcs = make([]uintptr, depth)

	n := runtime.Callers(2, pcs)

	var st stack = pcs[0:n]

	return &ErrorChain{error: errors.New(message), stack: st, mc: nil}
}

// NewErrorf create error with format and args.
func NewErrorf(format string, args ...any) error {
	var pcs = make([]uintptr, depth)

	n := runtime.Callers(2, pcs)

	var st stack = pcs[0:n]

	return &ErrorChain{error: fmt.Errorf(format, args...), stack: st, mc: nil}
}

// WrapErrorf wrap err as [ErrorChain] with formatted message.
//
// WrapErrorf doesn't reuse underlying Error if exists.
func WrapErrorf(err error, format string, args ...any) error {
	var pcs = make([]uintptr, depth)

	n := runtime.Callers(2, pcs)

	var st stack = pcs[0:n]

	return &ErrorChain{error: err, stack: st, mc: &mc{msg: fmt.Sprintf(format, args...)}}
}

// WrapError wrap err as [ErrorChain] with message.
//
// WrapError doesn't reuse underlying ErrorChain if exists.
func WrapError(err error, message string) error {
	var pcs = make([]uintptr, depth)

	n := runtime.Callers(2, pcs)

	var st stack = pcs[0:n]

	return &ErrorChain{error: err, stack: st, mc: &mc{msg: message}}
}

// Error construct chained error by err and message.
//
//	[Error] try to use existing stack and err.
//
// If the err is already a [ErrorChain], it mutate its message and return the original err.
func Error(err error, message string) error {
	//nolint:all //do not use errors.As this is only for direct fast path.
	if ws, ok := err.(*ErrorChain); ok {
		ws.mc = &mc{msg: message, parent: ws.mc}
		return ws
	}

	var pcs = make([]uintptr, depth)

	n := runtime.Callers(2, pcs)

	var st stack = pcs[0:n]

	return &ErrorChain{error: err, stack: st, mc: &mc{msg: message}}
}

// Errorf wraps an error into a [ErrorChain], appending a formatted message.
//
//	Errorf try to use existing stack and err.
//
// If the err is already a [ErrorChain], it mutate its message and return the original err.
func Errorf(err error, format string, args ...any) error {
	//nolint:all //do not use errors.As this is only for direct fast path.
	if ws, ok := err.(*ErrorChain); ok {
		ws.mc = &mc{msg: fmt.Sprintf(format, args...), parent: ws.mc}
		return ws
	}

	var pcs = make([]uintptr, depth)

	n := runtime.Callers(2, pcs)

	var st stack = pcs[0:n]

	return &ErrorChain{error: err, stack: st, mc: &mc{msg: fmt.Sprintf(format, args...)}}
}
