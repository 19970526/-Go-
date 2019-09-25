package compare_test

import (
	"reflect"
	"testing"
)

func TestCompare(t *testing.T) {
	var array1 = [3]int{1, 2, 3}
	var array2 = [3]int{2, 3, 4}

	t.Log("array1 == array2?", reflect.DeepEqual(array1, array2))

	var map1 = map[string]int{
		"a": 1,
		"b": 2,
	}
	var map2 = map[string]int{
		"c": 3,
		"d": 4,
	}

	t.Log("map1 == map2?", reflect.DeepEqual(map1, map2))
}
