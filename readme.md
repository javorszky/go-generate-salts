# Genearting salts for WordPress with go

## To deploy

1. clone repo
2. install packages with `govendor install +local`
3. run it with `go run salt.go` or compile with `go build salt.go` and run with `./salt`

## Benchmark

`go test -benchmem -bench=.`

As of 24th March 2018

```
$ go test -benchmem -bench=.
goos: darwin
goarch: amd64
pkg: heroku-dotenv
BenchmarkRandStringRunes-4              	  500000	      2857 ns/op	     336 B/op	       2 allocs/op
BenchmarkRandStringBytes-4              	 1000000	      2315 ns/op	     128 B/op	       2 allocs/op
BenchmarkRandStringBytesRmndr-4         	 1000000	      1817 ns/op	     128 B/op	       2 allocs/op
BenchmarkRandStringBytesMask-4          	  500000	      2855 ns/op	     128 B/op	       2 allocs/op
BenchmarkRandStringBytesMaskImpr-4      	 2000000	       715 ns/op	     128 B/op	       2 allocs/op
BenchmarkRandStringBytesMaskImprSrc-4   	 3000000	       498 ns/op	     128 B/op	       2 allocs/op
PASS
ok  	heroku-dotenv	11.223s
```

## Once it's running

`/` gets you standard snippet that you can drop into your `wp-config.php` file

`/env` gets you format that you can drop into your `.env` file
