# Select

Okay. So more about goroutines, only now we're synchronizing them beyond just having a thread-safe channel they can pump results into.

(that word will never look right to me. goroutines. gorountines. gorouTINEs. goroutines.)

I'm just gonna note, being willfully obstinate for a second, that "enough code to make it pass" isn't what the author says it is for that first go-round. "enough code to make it pass" would be:

```go
func Racer(first, second string) string {
	return second
}
```

But whatevs.


## On mocking stuff

Okay, so let me walk you through this.

[httptest](https://pkg.go.dev/net/http/httptest) is a package which provides utilities for HTTP testing.  The `NewRequest` function returns an incoming server request.

[Golang's http class](https://pkg.go.dev/net/http) provides a HandlerFunc type. It's an Adapter (more on those later) which lets you use any old function as an HTTP handler.

So what the hell are we doing here?

```go
    slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(20 * time.Millisecond)
        w.WriteHeader(http.StatusOK)
    }))

    fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    }))
```

  `slowServer` and `fastServer` are both tests for HTTP services. We create both of them and pass it a HandlerFunc, which we initialize with anonymous functions. The anonmous functions take an `http.ResponseWriter` and a reference to an `http.Request` as parameters.

  In the case of `fastServer`, the anonymous function returns a 200 status.  The `slowServer` does the same thing, but only after a pause of 20 milliseconds.

  In other words, we are actually writing a genuine web server here which is actually going to return a real status 200. Kinda cool.

  ## defer

  `defer` is kind of a lifecycle method that calls a function at the end of the containing function. We use it to clean up after our little HTTP servers in the test.

  ## select

  [select](https://golangdocs.com/select-statement-in-golang) is a bit of a weird one. Think of it like `switch`, but instead of checking a variable for variables, it's waiting for the first thing to update a channel.

  ```go
func Racer(a, b string) (winner string) {
  // select will choose the first thread that sends a message to a channel.  
	select {
    case <-ping(a):
      return a
    case <-ping(b):
      return b
	}
}

func ping(url string) chan struct{} {
  // the channel is defined here
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
```

## Timeouts

We use another `case` in our `select` to provide a timeout value:

```go
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
```

I don't hate this, but I _do_ dislike that we now have a happy path test that calls one function, and an unhappy path test that calls another. Yes, we're still getting all our test coverage, but I'm not 100% thrilled with that kind of behavior.

