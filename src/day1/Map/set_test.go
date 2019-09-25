package map_test

import "testing"

func TestSet(t *testing.T) {
	m := map[int]bool{}
	m[1] = true
	n := 1
	if m[n] {
		t.Logf("Key%d is existing", n)
	} else {
		t.Logf("Key%d is not existing", n)
	}
}
