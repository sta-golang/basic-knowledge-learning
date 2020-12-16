package main

import "testing"

func BenchmarkSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}
