package main

import "testing"

// from salt.go
func BenchmarkRandStringRunes8x64(b *testing.B) {
	// run the RandStringRunes function b.N times
	for n := 0; n < b.N; n++ {
		RandStringRunes(64)
		RandStringRunes(64)
		RandStringRunes(64)
		RandStringRunes(64)
		RandStringRunes(64)
		RandStringRunes(64)
		RandStringRunes(64)
		RandStringRunes(64)
	}
}

func BenchmarkRandStringRunes512(b *testing.B) {
	// run the RandStringRunes function b.N times
	for n := 0; n < b.N; n++ {
		RandStringRunes(512)
	}
}

func BenchmarkRandStringBytes8x64(b *testing.B) {
	// run the RandStringBytes function b.N times
	for n := 0; n < b.N; n++ {
		RandStringBytes(64)
		RandStringBytes(64)
		RandStringBytes(64)
		RandStringBytes(64)
		RandStringBytes(64)
		RandStringBytes(64)
		RandStringBytes(64)
		RandStringBytes(64)
	}
}

func BenchmarkRandStringBytes512(b *testing.B) {
	// run the RandStringBytes function b.N times
	for n := 0; n < b.N; n++ {
		RandStringBytes(512)
	}
}

func BenchmarkRandStringBytesRmndr8x64(b *testing.B) {
	// run the RandStringBytesRmndr function b.N times
	for n := 0; n < b.N; n++ {
		RandStringBytesRmndr(64)
		RandStringBytesRmndr(64)
		RandStringBytesRmndr(64)
		RandStringBytesRmndr(64)
		RandStringBytesRmndr(64)
		RandStringBytesRmndr(64)
		RandStringBytesRmndr(64)
		RandStringBytesRmndr(64)
	}
}

func BenchmarkRandStringBytesRmndr512(b *testing.B) {
	// run the RandStringBytesRmndr function b.N times
	for n := 0; n < b.N; n++ {
		RandStringBytesRmndr(512)
	}
}

func BenchmarkRandStringBytesMask8x64(b *testing.B) {
	// run the RandStringBytesMask function b.N times
	for n := 0; n < b.N; n++ {
		RandStringBytesMask(64)
		RandStringBytesMask(64)
		RandStringBytesMask(64)
		RandStringBytesMask(64)
		RandStringBytesMask(64)
		RandStringBytesMask(64)
		RandStringBytesMask(64)
		RandStringBytesMask(64)
	}
}

func BenchmarkRandStringBytesMask512(b *testing.B) {
	// run the RandStringBytesMask function b.N times
	for n := 0; n < b.N; n++ {
		RandStringBytesMask(512)
	}
}

func BenchmarkRandStringBytesMaskImpr8x64(b *testing.B) {
	// run the RandStringBytesMaskImpr function b.N times
	for n := 0; n < b.N; n++ {
		RandStringBytesMaskImpr(64)
		RandStringBytesMaskImpr(64)
		RandStringBytesMaskImpr(64)
		RandStringBytesMaskImpr(64)
		RandStringBytesMaskImpr(64)
		RandStringBytesMaskImpr(64)
		RandStringBytesMaskImpr(64)
		RandStringBytesMaskImpr(64)
	}
}

func BenchmarkRandStringBytesMaskImpr512(b *testing.B) {
	// run the RandStringBytesMaskImpr function b.N times
	for n := 0; n < b.N; n++ {
		RandStringBytesMaskImpr(512)
	}
}

func BenchmarkRandStringBytesMaskImprSrc8x64(b *testing.B) {
	// run the RandStringBytesMaskImprSrc function b.N times
	for n := 0; n < b.N; n++ {
		RandStringBytesMaskImprSrc(64)
		RandStringBytesMaskImprSrc(64)
		RandStringBytesMaskImprSrc(64)
		RandStringBytesMaskImprSrc(64)
		RandStringBytesMaskImprSrc(64)
		RandStringBytesMaskImprSrc(64)
		RandStringBytesMaskImprSrc(64)
		RandStringBytesMaskImprSrc(64)
	}
}

func BenchmarkRandStringBytesMaskImprSrc512(b *testing.B) {
	// run the RandStringBytesMaskImprSrc function b.N times
	for n := 0; n < b.N; n++ {
		RandStringBytesMaskImprSrc(512)
	}
}
