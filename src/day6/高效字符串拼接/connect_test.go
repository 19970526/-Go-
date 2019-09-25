package connect_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestConnect(t *testing.T) {
	var builder strings.Builder
	for i := 0; i < 10; i++ {
		builder.WriteString(strconv.Itoa(i))
	}
	result := builder.String()
	t.Log(result)
}
