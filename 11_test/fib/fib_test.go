package fib

import "testing"

func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		fib(n)
	}
}

func BenchmarkFib10(b *testing.B) {
	benchmarkFib(b, 10)
}
func BenchmarkFib100(b *testing.B) {
	benchmarkFib(b, 100)
}
func BenchmarkFib1000(b *testing.B) {
	benchmarkFib(b, 1000)
}
func BenchmarkFib10000(b *testing.B) {
	benchmarkFib(b, 10000)
}
