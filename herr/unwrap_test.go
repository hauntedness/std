package herr_test

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/hauntedness/std/herr"
)

type customError struct {
	code int
	msg  string
}

func (c *customError) Error() string {
	return fmt.Sprintf("code=%d msg=%s", c.code, c.msg)
}

func TestMatch_Nil(t *testing.T) {
	if herr.Match(nil, func(err error) bool { return true }) {
		t.Fatal("expected Match with nil err to return false")
	}

	err := errors.New("some error")
	if herr.Match(err, nil) {
		t.Fatal("expected Match with nil fn to return false")
	}
}

func TestMatch_SingleError(t *testing.T) {
	err := errors.New("connection reset by peer")

	matched := herr.Match(err, func(e error) bool {
		return strings.Contains(e.Error(), "connection reset")
	})
	if !matched {
		t.Fatalf("expected Match to find 'connection reset', but got false")
	}

	notMatched := herr.Match(err, func(e error) bool {
		return strings.Contains(e.Error(), "timeout")
	})
	if notMatched {
		t.Fatalf("expected Match not to find 'timeout', but got true")
	}
}

func TestMatch_TracedErrorChain(t *testing.T) {
	root := errors.New("disk full")
	err1 := herr.With(root, "failed to write log")
	err2 := herr.Wrap(err1, "service worker aborted")

	// Match root error
	if !herr.Match(err2, func(e error) bool { return errors.Is(e, root) }) {
		t.Fatal("expected Match to find root error in TracedError chain")
	}

	// Match middle error message
	if !herr.Match(err2, func(e error) bool { return strings.Contains(e.Error(), "failed to write log") }) {
		t.Fatal("expected Match to find middle error in TracedError chain")
	}

	// Non-matching condition
	if herr.Match(err2, func(e error) bool { return strings.Contains(e.Error(), "unrelated") }) {
		t.Fatal("expected Match to return false for unrelated predicate")
	}
}

func TestMatch_FmtWrapped(t *testing.T) {
	base := io.ErrUnexpectedEOF
	wrapped := fmt.Errorf("read payload: %w", base)

	if !herr.Match(wrapped, func(e error) bool { return errors.Is(e, io.ErrUnexpectedEOF) }) {
		t.Fatal("expected Match to find wrapped io.ErrUnexpectedEOF")
	}
}

func TestMatch_JoinedErrors(t *testing.T) {
	errA := errors.New("task A failed")
	errB := &customError{code: 404, msg: "task B not found"}
	joined := errors.Join(errA, errB)

	// Check matching errA
	if !herr.Match(joined, func(e error) bool { return e.Error() == "task A failed" }) {
		t.Fatal("expected Match to find task A in joined errors")
	}

	// Check matching custom error by type and property
	if !herr.Match(joined, func(e error) bool {
		var ce *customError
		if errors.As(e, &ce) {
			return ce.code == 404
		}
		return false
	}) {
		t.Fatal("expected Match to find customError with code 404 in joined errors")
	}

	// Non-matching
	if herr.Match(joined, func(e error) bool { return strings.Contains(e.Error(), "task C") }) {
		t.Fatal("expected Match to return false for non-existent error")
	}
}
