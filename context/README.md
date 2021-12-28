# Context

The [Context type](https://pkg.go.dev/context)  allows you to scope values to a specific request which may be used across API boundaries or processes. Any time you accept an incoming request, you should create a context, and outgoing calls should accept a context. In this example we're going to use a context to help us manage processes which may take a long time to run.

(An aside: you might wanna check out [this awesome Go blog post about the Context type](https://go.dev/blog/context). )

# Explain to me like I'm dumb what's going on here?
## The simple case
Let's start with the very basics, our initial `Store` interface.

```go
type Store interface {
	Fetch() string
}
```

This is pretty straightforward, right? A `Store` is something that supports a `Fetch` function, which returns a string. `Fetch` takes no parameters, it's just a convenient placeholder for sticking some stuff.

```go
func Server(store Store) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, store.Fetch())
    }
}
```

Our `Server` function take a `Store` as a parameter and sends back an `http.Handler` function. (Remember that handler? We used it [back in our greet example](../select/README.md).) The handler function takes a `ResponseWriter` and a `Request`, and writes some stuff out to the response - specifically, the formatted output of `store.Fetch()`.

## The more advanced case

We add a second test to simulate what happens if the store takes a long time to respond. This test case adds a `SpyStore` stub which adds a delay to our `Fetch` call.

```go
func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}
```

Then, after we create our request, we tell it to cancel five milliseconds after the request has been initiated, to simulate the user canceling the operation.

```go
      request := httptest.NewRequest(http.MethodGet, "/", nil)

      cancellingCtx, cancel := context.WithCancel(request.Context())
      time.AfterFunc(5 * time.Millisecond, cancel)
      request = request.WithContext(cancellingCtx)
```

To see the solution for implmenting that cancel function, check out [context.go](./context.go), but it can be summed up the following steps:

1. First, the Server gets the context from the `http.Request` that was passed to it.
2. Then it creates a channel to communicate over and spins off the call to `store.Fetch()` into a goroutine.
3. Finally, it waits to see which happens first - either the channel gets the result back from the store (in which case we format), or the context says that it's done, in which case we call `store.Cancel()`

