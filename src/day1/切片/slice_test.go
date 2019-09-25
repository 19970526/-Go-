package slice_test

import "testing"

func TestSlice1(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1, 2, 3, 4, 5, 6, 7)
	t.Log(len(s0), cap(s0))
}

func TestSlice2(t *testing.T) {
	s1 := make([]int, 3, 5)
	t.Log(s1, cap(s1))
	s1 = append(s1, 1, 2)
	t.Log(s1, cap(s1))
	s1 = append(s1, 3)
	t.Log(s1, cap(s1))
}
