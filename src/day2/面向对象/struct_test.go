package class_test

import (
	"fmt"
	"testing"
)

// 定义结构体
type Student struct {
	Name   string
	Age    int
	Number int
}

// 方法
func (s *Student) GetStudentImformation() {
	fmt.Printf("姓名:%s,年龄:%d,学号:%d\n", s.Name, s.Age, s.Number)
}

func TestStudent(t *testing.T) {
	student1 := Student{Name: "陈瑞麒", Age: 22, Number: 18}
	student1.GetStudentImformation()
}
