# Concurrency

Basic concurrency in Go uses the concept of a _goroutine_ . A goroutine is  a non-blocking, lightweight thread. In order to make thread safe data updates possible, go includes the concept of a _channel_.  Channels enable you to send data from one goroutine to another without having to worry about locking


```go

package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
    results := make(map[string]bool)
    // A channel is a typed conduit, so we create one here which is designed to 
    // pass our result type.
    resultChannel := make(chan result)

    // for each URL that we're checking
    for _, url := range urls {
        // We call an anonymous function (our goroutine)
        go func(u string) {
            // which pushers a result consisting of the URL and the output of the WebsiteChcker
            resultChannel <- result{u, wc(u)}
        }(url)
    }

    // Then for every URL in our list
    for i := 0; i < len(urls); i++ {
        // we receive a result from our resultChannel and put it in our results map.
        r := <-resultChannel
        results[r.string] = r.bool
    }

    return results
}
```

## Testing race conditions

Channels make short work of basic concurrency problems, but a way to check to make sure that you don't have race conditions is to run your tests using the `race` option. For example, here's what happens when we don't use channels for thread safety.:

```shell
❯ go test -race
==================
WARNING: DATA RACE
Write at 0x00c00009e330 by goroutine 9:
  runtime.mapassign_faststr()
      /usr/local/go/src/runtime/map_faststr.go:202 +0x0
  arrays-and-slices/concurrency.CheckWebsites.func1()
      /Users/sleibowitz/src/samleibowitz/learn-go-with-tests/concurrency/concurrency.go:14 +0x70
  arrays-and-slices/concurrency.CheckWebsites·dwrap·1()
      /Users/sleibowitz/src/samleibowitz/learn-go-with-tests/concurrency/concurrency.go:15 +0x58

Previous write at 0x00c00009e330 by goroutine 8:
  runtime.mapassign_faststr()
      /usr/local/go/src/runtime/map_faststr.go:202 +0x0
  arrays-and-slices/concurrency.CheckWebsites.func1()
      /Users/sleibowitz/src/samleibowitz/learn-go-with-tests/concurrency/concurrency.go:14 +0x70
  arrays-and-slices/concurrency.CheckWebsites·dwrap·1()
      /Users/sleibowitz/src/samleibowitz/learn-go-with-tests/concurrency/concurrency.go:15 +0x58

Goroutine 9 (running) created at:
  arrays-and-slices/concurrency.CheckWebsites()
      /Users/sleibowitz/src/samleibowitz/learn-go-with-tests/concurrency/concurrency.go:13 +0x1f8
  arrays-and-slices/concurrency.TestCheckWebsites()
      /Users/sleibowitz/src/samleibowitz/learn-go-with-tests/concurrency/concurrency_test.go:29 +0x165
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1259 +0x22f
  testing.(*T).Run·dwrap·21()
      /usr/local/go/src/testing/testing.go:1306 +0x47

Goroutine 8 (finished) created at:
  arrays-and-slices/concurrency.CheckWebsites()
      /Users/sleibowitz/src/samleibowitz/learn-go-with-tests/concurrency/concurrency.go:13 +0x1f8
  arrays-and-slices/concurrency.TestCheckWebsites()
      /Users/sleibowitz/src/samleibowitz/learn-go-with-tests/concurrency/concurrency_test.go:29 +0x165
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1259 +0x22f
  testing.(*T).Run·dwrap·21()
      /usr/local/go/src/testing/testing.go:1306 +0x47
==================
==================
WARNING: DATA RACE
Read at 0x00c000180088 by goroutine 7:
  reflect.typedmemmove()
      /usr/local/go/src/runtime/mbarrier.go:178 +0x0
  reflect.copyVal()
      /usr/local/go/src/reflect/value.go:1651 +0x6e
  reflect.Value.MapIndex()
      /usr/local/go/src/reflect/value.go:1537 +0x274
  reflect.deepValueEqual()
      /usr/local/go/src/reflect/deepequal.go:140 +0x1125
  reflect.DeepEqual()
      /usr/local/go/src/reflect/deepequal.go:218 +0x404
  arrays-and-slices/concurrency.TestCheckWebsites()
      /Users/sleibowitz/src/samleibowitz/learn-go-with-tests/concurrency/concurrency_test.go:31 +0x184
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1259 +0x22f
  testing.(*T).Run·dwrap·21()
      /usr/local/go/src/testing/testing.go:1306 +0x47

Previous write at 0x00c000180088 by goroutine 8:
  arrays-and-slices/concurrency.CheckWebsites.func1()
      /Users/sleibowitz/src/samleibowitz/learn-go-with-tests/concurrency/concurrency.go:14 +0x7c
  arrays-and-slices/concurrency.CheckWebsites·dwrap·1()
      /Users/sleibowitz/src/samleibowitz/learn-go-with-tests/concurrency/concurrency.go:15 +0x58

Goroutine 7 (running) created at:
  testing.(*T).Run()
      /usr/local/go/src/testing/testing.go:1306 +0x726
  testing.runTests.func1()
      /usr/local/go/src/testing/testing.go:1598 +0x99
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1259 +0x22f
  testing.runTests()
      /usr/local/go/src/testing/testing.go:1596 +0x7ca
  testing.(*M).Run()
      /usr/local/go/src/testing/testing.go:1504 +0x9d1
  main.main()
      _testmain.go:45 +0x22b

Goroutine 8 (finished) created at:
  arrays-and-slices/concurrency.CheckWebsites()
      /Users/sleibowitz/src/samleibowitz/learn-go-with-tests/concurrency/concurrency.go:13 +0x1f8
  arrays-and-slices/concurrency.TestCheckWebsites()
      /Users/sleibowitz/src/samleibowitz/learn-go-with-tests/concurrency/concurrency_test.go:29 +0x165
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1259 +0x22f
  testing.(*T).Run·dwrap·21()
      /usr/local/go/src/testing/testing.go:1306 +0x47
==================
--- FAIL: TestCheckWebsites (2.00s)
    testing.go:1152: race detected during execution of test
FAIL
```
