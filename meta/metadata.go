package meta

import (
	"runtime"
	"strconv"
	"strings"
)

// Caller return file and line number information
// based on [runtime.Caller]
// but in string form with at most 3-tier directories
func Caller() string {
	_, file, line, _ := runtime.Caller(1)
	buf := &strings.Builder{}
	writeTrimmedPath(buf, file)
	buf.WriteByte(':')
	buf.WriteString(strconv.Itoa(line))
	return buf.String()
}

// CallerSkip is similar to Caller but with skip specified
//
//	note: meta.CallerSkip(1) is equivalent to runtime.Caller(0)
func CallerSkip(skip int) string {
	_, file, line, _ := runtime.Caller(skip)
	buf := &strings.Builder{}
	writeTrimmedPath(buf, file)
	buf.WriteByte(':')
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
