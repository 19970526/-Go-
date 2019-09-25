package func_test

import (
	"math/rand"
	"testing"
)

func returnValue() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFunc(t *testing.T) {
	a, b := returnValue()
	t.Log(a, b)
}
