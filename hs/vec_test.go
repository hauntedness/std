package hs_test

import (
	"cmp"
	"fmt"
	"testing"

	"github.com/hauntedness/std/hs"
)

func TestVec_Append(t *testing.T) {
	vec := hs.NewWith(1, 2, 3)
	vec2 := vec.Append(4, 5, 6)
	if vec != vec2 {
		t.Fatal("append failed.")
	}
	expected := []int{1, 2, 3, 4, 5, 6}
	if !vec.Equal((*hs.Vec[int])(&expected), hs.Eq) {
		t.Fatal("compare failed.")
	}
}

func TestVec_Reduce(t *testing.T) {
	vec := hs.NewWith(1, 6, 3).Append(4, 5, 3)
	maximum := func(a, b int) int {
		return max(a, b)
	}
	ret := vec.Reduce(1, maximum)
	if ret != 6 {
		t.Fatal("reduce failed.")
	}
}

func TestVec_Reverse(t *testing.T) {
	vec := hs.NewWith(1, 2, 3, 4, 5, 6).Reverse()
	expected := []int{6, 5, 4, 3, 2, 1}
	if !vec.Equal((*hs.Vec[int])(&expected), hs.Eq) {
		t.Fatal("reverse failed.")
	}
}

func TestVec_Sort(t *testing.T) {
	vec := hs.NewWith(2, 1, 5, 4, 3, 6).Sort(cmp.Compare)
	expected := []int{1, 2, 3, 4, 5, 6}
	if !vec.Equal((*hs.Vec[int])(&expected), hs.Eq) {
		t.Fatal("sort failed.")
	}
}

func TestVec_BinarySearch(t *testing.T) {
	vec := hs.NewWith(1, 2, 3, 4, 6, 6)
	pos, ok := vec.BinarySearch(4, cmp.Compare)
	if !ok || pos != 3 {
		t.Fatal("BinarySearch failed.")
	}

	pos, ok = vec.BinarySearch(5, cmp.Compare)
	if ok || pos != 4 {
		t.Fatal("BinarySearch failed.")
	}
}

func TestVec_Pipe(t *testing.T) {
	vec0 := hs.NewWith(1, 2, 3, 4, 5, 6)
	vec1 := vec0.Pipe(func(i int) (int, bool) {
		return i + 1, true
	})
	expected := []int{2, 3, 4, 5, 6, 7}
	if !vec1.Equal((*hs.Vec[int])(&expected), hs.Eq) {
		t.Fatal("Pipe failed.")
	}
}

func TestString(t *testing.T) {
	vec := hs.NewWith(1, 2, 3, 4, 5, 6)
	if vec.String() != fmt.Sprint(&[]int{1, 2, 3, 4, 5, 6}) {
		t.Fatalf("String failed.")
	}
	vec = nil
	if vec.String() != fmt.Sprint((*[]int)(nil)) {
		t.Fatalf("String failed.")
	}
	_ = vec.String()
}
