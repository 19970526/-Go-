package error_test

import (
	"errors"
	"testing"
)

func ErrorTest(a int) (int, error) {
	if a < 0 {
		return 0, errors.New("输入不能小于0")
	} else {
		return a, nil
	}
}

func TestError(t *testing.T) {
	a, err := ErrorTest(-1)
	if err != nil {
		t.Log(err)
	} else {
		t.Log(a)
	}

}
