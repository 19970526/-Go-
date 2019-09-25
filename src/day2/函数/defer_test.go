package func_test

import (
	"testing"
)

// 可变长参数
func SumNumber(op ...int) int {
	res := 0
	for _, v := range op {
		res += v
	}
	return res
}

func TestSum(t *testing.T) {
	t.Log(SumNumber(1, 2, 3))
}

//defer延迟
func TestDefer(t *testing.T) {
	defer func() {
		t.Log("end")
	}()
	t.Log("start")
}
