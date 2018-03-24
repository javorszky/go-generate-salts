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

func BenchmarkRandStringBytesRmndr(b *testing.B) {
	// run the RandStringBytesRmndr function b.N times
	for n := 0; n < b.N; n++ {
		RandStringBytesRmndr(64)
	}
}

func BenchmarkRandStringBytesMask(b *testing.B) {
	// run the RandStringBytesMask function b.N times
	for n := 0; n < b.N; n++ {
		RandStringBytesMask(64)
	}
}

func BenchmarkRandStringBytesMaskImpr(b *testing.B) {
	// run the RandStringBytesMaskImpr function b.N times
	for n := 0; n < b.N; n++ {
		RandStringBytesMaskImpr(64)
	}
}

func BenchmarkRandStringBytesMaskImprSrc(b *testing.B) {
	// run the RandStringBytesMaskImprSrc function b.N times
	for n := 0; n < b.N; n++ {
		RandStringBytesMaskImprSrc(64)
	}
}
