package main

import "testing"

func BenchmarkSingleSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Search("C:\\")
	}
}
