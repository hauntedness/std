package meta

import (
	"runtime"
	"strconv"
	"strings"
)

type CallerOption int

const (
	// apply function name to result
	WithoutFunctionName CallerOption = 1 << iota
	// apply file path to result
	WithFilePath
)

// Caller return file and line number information
// based on [runtime.Caller]
// but in string form with at most 3-tier directories
func Caller(options ...CallerOption) string {
	return CallerSkip(1, options...)
}

// CallerSkip is similar to Caller but with skip specified
//
//	note: meta.CallerSkip(1) is equivalent to runtime.Caller(0)
func CallerSkip(skip int, options ...CallerOption) string {
	var option CallerOption
	for _, op := range options {
		option |= op
	}
	pc, file, line, _ := runtime.Caller(skip)
	buf := &strings.Builder{}
	// apply function name to result
	if option&WithoutFunctionName != WithoutFunctionName {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			buf.WriteString(fn.Name())
			buf.WriteByte(':')
		}
	}
	// apply file path to result
	if option&WithFilePath == WithFilePath {
		writeTrimmedPath(buf, file)
		buf.WriteByte(':')
	}
	buf.WriteString(strconv.Itoa(line))
	return buf.String()
}

func writeTrimmedPath(buf *strings.Builder, path string) {
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
