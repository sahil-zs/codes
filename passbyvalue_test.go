package codes

import "testing"

func Benchmark11(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Passbyvaluee(5, 6)
	}
}

func Benchmark12(b *testing.B) {
	a := 5
	c := 6
	for i := 0; i < b.N; i++ {
		Passbyreferencee(&a, &c)
	}
}
