package codes

import (
	"testing"
)

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calculate(3, 4, Add)
	}
}

func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Compute(3, 4, "+")
	}
}
