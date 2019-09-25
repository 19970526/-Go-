package emtry_interface_test

import (
	"fmt"
	"testing"
)

// 判断空接口类型：方法1
func EmtryInterface1(a interface{}) {
	if t, ok := a.(int); ok {
		fmt.Printf("%d是整型", t)
	} else {
		fmt.Printf("%v不是整型", t)
	}
}

// 判断空接口类型：方法1
func EmtryInterface2(a interface{}) {

	switch a.(type) {
	case string:
		fmt.Printf("%s是字符串类型", a)
	case int:
		fmt.Printf("%d是整型", a)
	case bool:
		fmt.Printf("%v是布尔型", a)
	default:
		fmt.Println("未知类型")
	}
}


func TestEmtryInterface(t *testing.T) {
	EmtryInterface1(10)
	EmtryInterface2("abc")
}
