# Genearting salts for WordPress with go

## To deploy

1. clone repo
2. install packages with `go mod tidy && go mod vendor`
3. run it with `go run salt.go` or compile with `go build salt.go` and run with `./salt`

## Once it's running

`/` gets you standard snippet that you can drop into your `wp-config.php` file

`/env` gets you format that you can drop into your `.env` file

`/json` gets you the salts with key-value in JSON format

## Benchmark

We're only using the optimised fastest / safest versions of the function calls.

`go test -benchmem -bench=.`

As of 24th March 2018

```
❯ go test -benchmem -bench=.
goos: darwin
goarch: amd64
pkg: github.com/javorszky/go-generate-salts
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkRandStringBytesMaskImpr8x64-16    	  382617	      3176 ns/op	    1024 B/op	      16 allocs/op
BenchmarkRandStringBytesMaskImpr512-16     	  371102	      3018 ns/op	    1024 B/op	       2 allocs/op
BenchmarkGenerateSaltsWP512-16             	  178719	      6263 ns/op	    3698 B/op	      49 allocs/op
BenchmarkGenerateSaltsEnv512-16            	  212131	      5166 ns/op	    3138 B/op	      31 allocs/op
BenchmarkGenerateSaltsJSON512-16           	  304455	      3851 ns/op	    2065 B/op	       5 allocs/op
BenchmarkSrcInt63Parallel-16               	  149832	      7800 ns/op	    1024 B/op	       2 allocs/op
PASS
ok  	github.com/javorszky/go-generate-salts	8.360s
```

## Thank you!

* https://stackoverflow.com/a/31832326/2862802 for the fantastic thread / explanations
* [@djavorszky](https://github.com/djavorszky) for rubber ducking and pushing me to do the benchmarks
