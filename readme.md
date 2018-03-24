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
BenchmarkRandStringRunes8x64-4               100000    22526 ns/op   2688 B/op     16 allocs/op
BenchmarkRandStringRunes512-4                100000    22128 ns/op   2624 B/op      2 allocs/op
BenchmarkRandStringBytes8x64-4               100000    18103 ns/op   1024 B/op     16 allocs/op
BenchmarkRandStringBytes512-4                100000    17739 ns/op   1024 B/op      2 allocs/op
BenchmarkRandStringBytesRmndr8x64-4          100000    14510 ns/op   1024 B/op     16 allocs/op
BenchmarkRandStringBytesRmndr512-4           100000    14293 ns/op   1024 B/op      2 allocs/op
BenchmarkRandStringBytesMask8x64-4           100000    20262 ns/op   1024 B/op     16 allocs/op
BenchmarkRandStringBytesMask512-4            100000    19774 ns/op   1024 B/op      2 allocs/op
BenchmarkRandStringBytesMaskImpr8x64-4       300000     5157 ns/op   1024 B/op     16 allocs/op
BenchmarkRandStringBytesMaskImpr512-4        300000     4651 ns/op   1024 B/op      2 allocs/op
BenchmarkRandStringBytesMaskImprSrc8x64-4    300000     3952 ns/op   1024 B/op     16 allocs/op
BenchmarkRandStringBytesMaskImprSrc512-4     300000     3886 ns/op   1024 B/op      2 allocs/op
PASS
ok  	heroku-dotenv	22.007s
```

## Thank you!

* https://stackoverflow.com/a/31832326/2862802 for the fantastic thread / explanations
* @djavorszky for rubber ducking and pushing me to do the benchmarks
