package caller

import (
	"bytes"
	"runtime"
	"strconv"
	"strings"
)

// Name return caller function name and line number information
//
// based on [runtime.Caller]
func Name() string {
	pc, _, line, _ := runtime.Caller(1)
	buf := &bytes.Buffer{}
	// apply function name to result
	fn := runtime.FuncForPC(pc)
	if fn != nil {
		buf.WriteString(fn.Name())
		buf.WriteByte(':')
		buf.WriteString(strconv.Itoa(line))
	}
	return buf.String()
}

// NameSkip is similar to [Name] but with skip specified
//
//	note: NameSkip(1) is equivalent to runtime.Caller(0)
func NameSkip(skip int) string {
	pc, _, line, _ := runtime.Caller(skip)
	buf := &bytes.Buffer{}
	// apply function name to result
	fn := runtime.FuncForPC(pc)
	if fn != nil {
		buf.WriteString(fn.Name())
		buf.WriteByte(':')
		buf.WriteString(strconv.Itoa(line))
	}
	return buf.String()
}

// Nm return caller function abbr and line number information
//
// based on [runtime.Caller]
func Nm() string {
	pc, _, line, _ := runtime.Caller(1)
	buf := &bytes.Buffer{}
	// apply function name to result
	fn := runtime.FuncForPC(pc)
	if fn != nil {
		writeTrimmedPath(buf, fn.Name())
		buf.WriteByte(':')
		buf.WriteString(strconv.Itoa(line))
	}
	return buf.String()
}

// NmSkip is similar to [Nm] but with skip specified
//
//	note: NmSkip(1) is equivalent to runtime.Caller(0)
func NmSkip(skip int) string {
	pc, _, line, _ := runtime.Caller(skip)
	buf := &bytes.Buffer{}
	// apply function name to result
	fn := runtime.FuncForPC(pc)
	if fn != nil {
		writeTrimmedPath(buf, fn.Name())
		buf.WriteByte(':')
		buf.WriteString(strconv.Itoa(line))
	}
	return buf.String()
}

// Path return caller file name and line number information
//
//	note: Path(1) is equivalent to runtime.Caller(0)
func Path() string {
	_, file, line, _ := runtime.Caller(1)
	buf := &bytes.Buffer{}
	buf.WriteString(file)
	buf.WriteByte(':')
	buf.WriteString(strconv.Itoa(line))
	return buf.String()
}

// PathSkip is similar to [Path] but with skip specified
//
//	note: PathSkip(1) is equivalent to runtime.Caller(0)
func PathSkip(skip int) string {
	_, file, line, _ := runtime.Caller(skip)
	buf := &bytes.Buffer{}
	buf.WriteString(file)
	buf.WriteByte(':')
	buf.WriteString(strconv.Itoa(line))
	return buf.String()
}

// ShortPathSkip is similar to [PathSkip] but in short string form
//
//	note: ShortPathSkip(1) is equivalent to runtime.Caller(0)
func ShortPathSkip(skip int) string {
	_, file, line, _ := runtime.Caller(skip)
	buf := &bytes.Buffer{}
	writeTrimmedPath(buf, file)
	buf.WriteByte(':')
	buf.WriteString(strconv.Itoa(line))
	return buf.String()
}

// writeTrimmedPath write path with at most 3-tier directories
func writeTrimmedPath(buf *bytes.Buffer, path string) {
	idx := strings.LastIndexByte(path, '/')
	if idx == -1 {
		buf.WriteString(path)
		return
	}
	// Find the penultimate separator.
	idx = strings.LastIndexByte(path[:idx], '/')
	if idx == -1 {
		buf.WriteString(path)
		return
	}
	// Find the penultimate separator.
	idx2 := strings.LastIndexByte(path[:idx], '/')
	if idx == -1 {
		buf.WriteString(path[idx+1:])
		return
	}
	buf.WriteString(path[idx2+1:])
}
