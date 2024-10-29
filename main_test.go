package main

import "testing"

// 6.782s
// 6778759411 ns/op
func BenchmarkMyFn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BruteForce("Janko")
	}
}

// 8.635s
// 8632277364 ns/op
func BenchmarkGPT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BruteForceChatGPT("Janko")
	}
}
