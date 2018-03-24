# Genearting salts for WordPress with go

## To deploy

1. clone repo
2. install packages with `govendor install +local`
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
$ go test -benchmem -bench=.
goos: darwin
goarch: amd64
pkg: heroku-dotenv
BenchmarkRandStringBytesMaskImpr8x64-4    300000     5199 ns/op   1024 B/op     16 allocs/op
BenchmarkRandStringBytesMaskImpr512-4     300000     4687 ns/op   1024 B/op      2 allocs/op
BenchmarkGenerateSaltsWP512-4             200000     8511 ns/op   4193 B/op     52 allocs/op
BenchmarkGenerateSaltsEnv512-4            200000     6769 ns/op   3249 B/op     28 allocs/op
BenchmarkGenerateSaltsJSON512-4           200000     5150 ns/op   1360 B/op      4 allocs/op
BenchmarkSrcInt63Parallel-4               200000    11216 ns/op   1024 B/op      2 allocs/op
PASS
ok      heroku-dotenv    9.803s
```

## Profiling

1. run app
2. siege one of the endpoints continuously (`siege -c 250 -t2m -v http://localhost:8090` works nicely)
3. _while_ siege is running, in another terminal, start the profile: `go tool pprof http://localhost:8090`
4. enjoy your data! See [https://blog.golang.org/profiling-go-programs](https://blog.golang.org/profiling-go-programs) for more info on profiling
5. 48.9% syscall ;)

## Thank you!

* https://stackoverflow.com/a/31832326/2862802 for the fantastic thread / explanations
* [@djavorszky](https://github.com/djavorszky) for rubber ducking and pushing me to do the benchmarks
