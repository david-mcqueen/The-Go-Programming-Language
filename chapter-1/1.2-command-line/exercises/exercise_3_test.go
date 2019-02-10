package exercises

import "testing"

func Benchmark_3_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exercise3_1()
	}
	// 85000 ns/op
}

func Benchmark_3_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exercise3_2()
	}

	// ~80000 ns/op
}
