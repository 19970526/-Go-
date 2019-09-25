package array_test

import "testing"

func TestArray(t *testing.T) {
	var a [3]int
	t.Log(a)
	a2 := [3]int{1, 2, 3}
	t.Log(a2)
}

func TestArrayTravel(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5}
	for i, j := range a {
		t.Log(i, j)
	}
}

func TestArraySecion(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5}
	t.Log(a[:3])
}
