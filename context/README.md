# Context

The [Context type](https://pkg.go.dev/context)  allows you to scope values to a specific request which may be used across API boundaries or processes. Any time you accept an incoming request, you should create a context, and outgoing calls should accept a context. In this example we're going to use a context to help us manage processes which may take a long time to run.

(An aside: you might wanna check out [this awesome Go blog post about the Context type](https://go.dev/blog/context). )

## Explain to me like I'm dumb what's going on here?

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

