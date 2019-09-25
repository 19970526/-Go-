package map_test

import "testing"

func TestMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	t.Logf("m1 len is %d", len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("m2 len is %d", len(m2))
	m3 := make(map[int]int)
	t.Logf("m3 len is %d", len(m3))
}

func TestMap2(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	//m1[4] = 4
	// 不存在默认为0，下面方法判断0是赋值的还是不存在的
	if v, ok := m1[4]; ok {
		t.Logf("m1[4] is %d", v)
	} else {
		t.Logf("m1[4] is not")
	}
}
