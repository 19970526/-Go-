package switch_test

import "testing"

func TestSwitch(t *testing.T) {
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			t.Log(i, "：偶数")
		case i%2 == 1:
			t.Log(i, "：奇数")
		default:
			t.Log(i, "：不知道是啥")
		}
	}
}
