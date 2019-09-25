package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("整型")
	case reflect.Float32, reflect.Float64:
		fmt.Println("浮点型")
	default:
		fmt.Println("其他")
	}
}

func TestReflect(t *testing.T) {
	CheckType(1)
}

func TestValue(t *testing.T) {
	var a int = 10
	t.Log(reflect.TypeOf(a), reflect.ValueOf(a))
	t.Log(reflect.ValueOf(a).Type(), reflect.ValueOf(a))
}
