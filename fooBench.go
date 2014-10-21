package foo

import "testing"

func BenchmarkFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Do work here
	}
}

func BenchmarkBar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Do different work here
	}
}