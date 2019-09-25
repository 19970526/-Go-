package benchmark_test

import (
	"fmt"
	"testing"
)

func BenchmarkFor(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 10; i++ {
			fmt.Println("hello")
		}
	}
	b.StopTimer()
}
