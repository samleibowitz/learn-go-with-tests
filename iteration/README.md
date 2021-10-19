# Iteration

So it turns out that there's no `while`, `do`, or `until` in Go, so you gotta use `for`. (They claim this is a good thing. I am suspicious of their reasoning, but whatevs.)

## Red

Be in the habit of making the test fail first. (But make it a good failure. A compilation error because you didn't write the source file yet isn't that useful)

## Write enough code to make it pass

Notice Go's `for` syntax. Pretty boring, right?

```go
for i := 0; i < 5; i++ {
  repeated = repeated + character
}
```


## Refactor

Okay, Go has a `+=` operator, whee.

## Benchmarking

Okay, what? This is cool. Check out `BenchmarkRepeat`. I ran it and got the following output:

```
Running tool: /usr/local/go/bin/go test -benchmem -run=^$ -bench ^BenchmarkRepeat$ repeat

goos: darwin
goarch: amd64
pkg: repeat
cpu: Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz
BenchmarkRepeat-12    	11202793	       106.4 ns/op	      16 B/op	       4 allocs/op
PASS
ok  	repeat	1.898s
```

On average, it took 106.4 nanoseconds to run, averaged over the course of... uh... 11,202,793 repetitions. That's a lot.

