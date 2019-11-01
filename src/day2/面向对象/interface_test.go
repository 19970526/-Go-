package class_test

import (
	"fmt"
	"testing"
)

//定义接口
type P interface {
	GetName(job string) string
}

type Student2 struct {
	name string
}

func (s *Student2) GetName(job string) string {
	return fmt.Sprintf("I'm a %s,my name is %s", job, s.name)
}

func TestInterface(t *testing.T) {
	var s P

	s = &Student2{name: "James"}

	t.Log(s.GetName("student"))
}
