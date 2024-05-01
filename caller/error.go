package caller

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
)

func Error(err error) error {
	pc, _, line, _ := runtime.Caller(1)
	buf := &bytes.Buffer{}
	// apply function name to result
	fn := runtime.FuncForPC(pc)
	if fn != nil {
		buf.WriteString(fn.Name())
		buf.WriteByte(':')
		buf.WriteString(strconv.Itoa(line))
	}
	buf.WriteString(": ")
	return &withcaller{name: buf.String(), err: err}
}

// Errorf is like fmt.Errorf but append caller name to the error
func Errorf(format string, args ...any) error {
	pc, _, line, _ := runtime.Caller(1)
	buf := &bytes.Buffer{}
	// apply function name to result
	fn := runtime.FuncForPC(pc)
	if fn != nil {
		buf.WriteString(fn.Name())
		buf.WriteByte(':')
		buf.WriteString(strconv.Itoa(line))
	}
	buf.WriteString(": ")
	return &withcaller{name: buf.String(), err: fmt.Errorf(format, args...)}
}

type withcaller struct {
	name string
	err  error
}

func (e *withcaller) Error() string {
	return e.name + e.err.Error()
}

func (e *withcaller) Unwrap() error {
	return e.err
}
