package main

import "testing"

// from salt.go
func BenchmarkRandStringRunes(b *testing.B) {
	// run the RandStringRunes function b.N times
	for n := 0; n < b.N; n++ {
		RandStringRunes(64)
	}
}

func BenchmarkRandStringBytes(b *testing.B) {
	// run the RandStringBytes function b.N times
	for n := 0; n < b.N; n++ {
		RandStringBytes(64)
	}
}
