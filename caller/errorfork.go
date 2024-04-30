package caller

import (
	"fmt"
	"io"
)

func Error(err error, msg string) error {
	if err == nil {
		return nil
	}
	return &withStack{
		msg:   msg,
		err:   err,
		stack: callers(),
	}
}

func Errorf(err error, format string, args ...any) error {
	if err == nil {
		return nil
	}
	return &withStack{
		msg:   fmt.Sprintf(format, args...),
		err:   err,
		stack: callers(),
	}
}

type withStack struct {
	msg string
	err error
	*stack
}

func (w *withStack) Error() string {
	return w.msg + ": " + w.err.Error()
}

func (w *withStack) Unwrap() error { return w.err }

func (w *withStack) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v", w.Unwrap())
			w.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, w.Error())
	case 'q':
		fmt.Fprintf(s, "%q", w.Error())
	}
}
