package caller

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func Error(err error) error {
	pc, _, line, _ := runtime.Caller(1)
	buf := &strings.Builder{}
	// apply function name to result
	fn := runtime.FuncForPC(pc)
	if fn != nil {
		buf.WriteString(fn.Name())
		buf.WriteByte(':')
		buf.WriteString(strconv.Itoa(line))
	}
	return fmt.Errorf("%s: %w", buf.String(), err)
}
