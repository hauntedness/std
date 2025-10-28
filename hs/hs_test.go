package hs_test

import (
	"testing"

	"github.com/hauntedness/std/hs"
)

func TestLoc(t *testing.T) {
	res := hs.Loc([]int{1, 2, 3}, 1, -1)
	if len(res) != 1 || res[0] != 2 {
		t.Fatal("loc failed")
	}
}

func TestTail(t *testing.T) {
	res := hs.Tail([]int{1, 2, 3}, -2)
	if len(res) != 2 || res[0] != 2 {
		t.Fatal("Tail failed")
	}
}
