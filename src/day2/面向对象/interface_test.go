package class_test

import "testing"

//定义接口
type P interface {
	GetName() string
}

type Student2 struct {
	name string
}

func (s *Student2) GetName() string {
	return s.name
}

func TestInterface(t *testing.T) {
	var s P

	s = &Student2{name: "James"}

	t.Log(s.GetName())
}
