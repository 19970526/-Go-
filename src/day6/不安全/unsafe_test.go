package unsafe_test

import (
	"testing"
	"unsafe"
)

type MyInt int

func TestUnsafe(t *testing.T) {
	var array1 = [3]int{1, 2, 3}
	var array2 = *(*[3]MyInt)(unsafe.Pointer(&array1))
	t.Logf("array1类型%[1]T,值%[1]v", array1)
	t.Logf("array2类型%[1]T,值%[1]v", array2)
}
