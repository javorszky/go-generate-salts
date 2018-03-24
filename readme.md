# Genearting salts for WordPress with go

## To deploy

1. clone repo
2. install packages with `govendor install +local`
3. run it with `go run salt.go` or compile with `go build salt.go` and run with `./salt`

## Once it's running

`/` gets you standard snippet that you can drop into your `wp-config.php` file

`/env` gets you format that you can drop into your `.env` file

## Benchmark

`go test -benchmem -bench=.`

As of 24th March 2018

```
$ go test -benchmem -bench=.
goos: darwin
goarch: amd64
pkg: heroku-dotenv
BenchmarkRandStringRunes8x64-4                100000     21809 ns/op    2688 B/op      16 allocs/op
BenchmarkRandStringRunes512-4                 100000     21222 ns/op    2624 B/op       2 allocs/op
BenchmarkRandStringBytes8x64-4                100000     17721 ns/op    1024 B/op      16 allocs/op
BenchmarkRandStringBytes512-4                 100000     17212 ns/op    1024 B/op       2 allocs/op
BenchmarkRandStringBytesRmndr8x64-4           100000     13477 ns/op    1024 B/op      16 allocs/op
BenchmarkRandStringBytesRmndr512-4            100000     13306 ns/op    1024 B/op       2 allocs/op
BenchmarkRandStringBytesMask8x64-4            100000     19478 ns/op    1024 B/op      16 allocs/op
BenchmarkRandStringBytesMask512-4             100000     18988 ns/op    1024 B/op       2 allocs/op
BenchmarkRandStringBytesMaskImpr8x64-4        300000      5179 ns/op    1024 B/op      16 allocs/op
BenchmarkRandStringBytesMaskImpr512-4         300000      4756 ns/op    1024 B/op       2 allocs/op
BenchmarkRandStringBytesMaskImprSrc8x64-4     300000      3979 ns/op    1024 B/op      16 allocs/op
BenchmarkRandStringBytesMaskImprSrc512-4      500000      3600 ns/op    1024 B/op       2 allocs/op
BenchmarkGenerateSaltsWP8x64-4                200000      7933 ns/op    4193 B/op      66 allocs/op
BenchmarkGenerateSaltsWP512-4                 200000      7249 ns/op    4193 B/op      52 allocs/op
BenchmarkGenerateSaltsEnv8x64-4               200000      6288 ns/op    3248 B/op      42 allocs/op
BenchmarkGenerateSaltsEnv512-4                200000      5955 ns/op    3248 B/op      28 allocs/op
PASS
ok      heroku-dotenv     27.816s
```

## Thank you!

* https://stackoverflow.com/a/31832326/2862802 for the fantastic thread / explanations
* @djavorszky for rubber ducking and pushing me to do the benchmarks
