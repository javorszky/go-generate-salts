package main

import "testing"

func BenchmarkRandStringBytesMaskImpr8x64(b *testing.B) {
	// run the randStringBytesMaskImpr function b.N times
	for n := 0; n < b.N; n++ {
		randStringBytesMaskImpr(64)
		randStringBytesMaskImpr(64)
		randStringBytesMaskImpr(64)
		randStringBytesMaskImpr(64)
		randStringBytesMaskImpr(64)
		randStringBytesMaskImpr(64)
		randStringBytesMaskImpr(64)
		randStringBytesMaskImpr(64)
	}
}

func BenchmarkRandStringBytesMaskImpr512(b *testing.B) {
	// run the randStringBytesMaskImpr function b.N times
	for n := 0; n < b.N; n++ {
		randStringBytesMaskImpr(512)
	}
}

func BenchmarkGenerateSaltsWP512(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateSaltsWPEfficient()
	}
}

func BenchmarkGenerateSaltsEnv512(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateSaltsEnvEfficient()
	}
}

func BenchmarkGenerateSaltsJSON512(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateSaltsJSONEfficient()
	}
}

func BenchmarkSrcInt63Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			randStringBytesMaskImpr(512)
		}
	})
}
