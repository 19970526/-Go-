package reflect_test

import (
	"reflect"
	"testing"
)

type Student struct {
	name string `json:"student_name"`
	age  int    `json:"student_age"`
}

func (s *Student) GetInformations() (string, int) {
	return s.name, s.age
}

func TestReflectStruct(t *testing.T) {
	obj := Student{
		name: "ricky",
		age:  22,
	}

	t.Log(reflect.ValueOf(obj).FieldByName("name"))

	obj_value, ok := reflect.TypeOf(obj).FieldByName("name")
	if !ok {
		t.Error("没有")
	} else {
		t.Log(obj_value.Tag.Get("json"))
	}

}
