package codes

import "testing"

var a = [10]int{2, 3, 5, 7, 11, 13, 19, 20, 22, 24}

func Benchmark3(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Linear(a, 13)
	}
}

func Benchmark4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Binary(a, 13)
	}
}
