# Genearting salts for WordPress with go

## To deploy

1. clone repo
2. install packages with `govendor install +local`
3. run it with `go run salt.go` or compile with `go build salt.go` and run with `./salt`

## Benchmark

`go test -benchmem -bench=.`

## Once it's running

`/` gets you standard snippet that you can drop into your `wp-config.php` file

`/env` gets you format that you can drop into your `.env` file
