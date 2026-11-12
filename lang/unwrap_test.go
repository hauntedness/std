package lang_test

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/hauntedness/std/lang"
)

type customError struct {
	code int
	msg  string
}

func (c *customError) Error() string {
	return fmt.Sprintf("code=%d msg=%s", c.code, c.msg)
}

func TestMatch_Nil(t *testing.T) {
	if lang.Match(nil, func(err error) bool { return true }) {
		t.Fatal("expected Match with nil err to return false")
	}

	err := errors.New("some error")
	if lang.Match(err, nil) {
		t.Fatal("expected Match with nil fn to return false")
	}
}

func TestMatch_SingleError(t *testing.T) {
	err := errors.New("connection reset by peer")

	matched := lang.Match(err, func(e error) bool {
		return strings.Contains(e.Error(), "connection reset")
	})
	if !matched {
		t.Fatalf("expected Match to find 'connection reset', but got false")
	}

	notMatched := lang.Match(err, func(e error) bool {
		return strings.Contains(e.Error(), "timeout")
	})
	if notMatched {
		t.Fatalf("expected Match not to find 'timeout', but got true")
	}
}

func TestMatch_TracedChainedError(t *testing.T) {
	root := errors.New("disk full")
	err1 := lang.ErrorWith(root, "failed to write log")
	err2 := lang.Error(err1, "service worker aborted")

	// Match root error
	if !lang.Match(err2, func(e error) bool { return errors.Is(e, root) }) {
		t.Fatal("expected Match to find root error in TracedError chain")
	}

	// Match middle error message
	if !lang.Match(err2, func(e error) bool { return strings.Contains(e.Error(), "failed to write log") }) {
		t.Fatal("expected Match to find middle error in TracedError chain")
	}

	// Non-matching condition
	if lang.Match(err2, func(e error) bool { return strings.Contains(e.Error(), "unrelated") }) {
		t.Fatal("expected Match to return false for unrelated predicate")
	}
}

func TestMatch_FmtWrapped(t *testing.T) {
	base := io.ErrUnexpectedEOF
	wrapped := fmt.Errorf("read payload: %w", base)

	if !lang.Match(wrapped, func(e error) bool { return errors.Is(e, io.ErrUnexpectedEOF) }) {
		t.Fatal("expected Match to find wrapped io.ErrUnexpectedEOF")
	}
}

func TestMatch_JoinedErrors(t *testing.T) {
	errA := errors.New("task A failed")
	errB := &customError{code: 404, msg: "task B not found"}
	joined := errors.Join(errA, errB)

	// Check matching errA
	if !lang.Match(joined, func(e error) bool { return e.Error() == "task A failed" }) {
		t.Fatal("expected Match to find task A in joined errors")
	}

	// Check matching custom error by type and property
	if !lang.Match(joined, func(e error) bool {
		if ce, ok := errors.AsType[*customError](e); ok {
			return ce.code == 404
		}

		return false
	}) {
		t.Fatal("expected Match to find customError with code 404 in joined errors")
	}

	// Non-matching
	if lang.Match(joined, func(e error) bool { return strings.Contains(e.Error(), "task C") }) {
		t.Fatal("expected Match to return false for non-existent error")
	}
}
