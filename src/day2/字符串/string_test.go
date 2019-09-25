package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	s1 := "a,b,c"
	//分割
	sl := strings.Split(s1, ",")
	t.Logf("类型是：%[1]T，值是：%[1]v", sl)
	//合并
	s2 := strings.Join(sl, "-")
	t.Logf("类型是：%[1]T，值是：%[1]s", s2)
}

func TestChangeString(t *testing.T) {
	//字符串转数字
	s1, err := strconv.Atoi("10")
	if err == nil {
		t.Logf("类型是%[1]T,值是%[1]d", s1)
	}
	//数字转字符串
	s2 := strconv.Itoa(10)
	t.Logf("类型是%[1]T,值是%[1]s", s2)
}
