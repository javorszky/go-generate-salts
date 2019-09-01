package main

import "testing"

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

func BenchmarkGenerateSaltsWP512(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GenerateSaltsWPEfficient()
	}
}

func BenchmarkGenerateSaltsEnv512(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GenerateSaltsEnvEfficient()
	}
}

func BenchmarkGenerateSaltsJSON512(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GenerateSaltsJSONEfficient()
	}
}

func BenchmarkSrcInt63Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			RandStringBytesMaskImpr(512)
		}
	})
}
