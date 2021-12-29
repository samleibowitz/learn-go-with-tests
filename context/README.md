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

## Refactoring to best practices

The author quotes some great stuff here that I'm gonna also quote in full:

> Incoming requests to a server should create a Context, and outgoing calls to servers should accept a Context. The chain of function calls between them must propagate the Context, optionally replacing it with a derived Context created using WithCancel, WithDeadline, WithTimeout, or WithValue. When a Context is canceled, all Contexts derived from it are also canceled.

> At Google, we require that Go programmers pass a Context parameter as the first argument to every function on the call path between incoming and outgoing requests. This allows Go code developed by many different teams to interoperate well. It provides simple control over timeouts and cancelation and ensures that critical values like security credentials transit Go programs properly.

So my commentary here: there is a level at which I'm pretty uncomfortable, at least initially, with this rule. Where is the bottom here? Exactly how far down the stack do you have to go before you don't have to propagate this context object? Does that mean that anything that doesn't accept that context parameter is therefore ineligble for consideration in a Google project?

But! Let's figure out how this thing would work.

Responsibility on termination is no longer on the `Server`, but rather on the `Store`:

```go
func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}
```

There's something that seems deeply unsettling about this to me - we started off by creating a Server function and testing its happy path, then added handling the unhappy path, then refactored the unhappy path behavior out of the Server and into the function it gets passed. That's deeply weird, because our `SpyStore` type is now inherently more complicated than the thing that we're using it to test! In real life I'd never have written the second version of this test, since the Store object is doing all the work.

# And a bit more besides

At the end, the author quotes Michal Strba's [Context should go away for Go 2
[Context should go away for Go 2](https://faiface.github.io/post/context-should-go-away-go2/) in order to throw cold water on the idea of using `context.Value` everwhere. But it turns out that, as you may have guessed from the title of his article, he thinks that propagating contexts everywhere kinda sucks.... the same criticism I had previously. Huh.