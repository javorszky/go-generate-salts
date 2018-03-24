package main

import "testing"

// from salt.go
func BenchmarkRandStringRunes(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		RandStringRunes(64)
	}
}
