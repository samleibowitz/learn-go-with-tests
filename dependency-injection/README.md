# Dependency Injection

DI isn't exactly a language feature, it's a technique, so I'm curious to see how the author handles it in this section.

## A brief aside on Dependency Injection

`Fprintf` is a classic use of dependency injection in a bunch of different languages. 

```go
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
```

Here we're passing in a variable of type `io.Writer` because that's where our outbput is going to go. In theory, it could be a byte buffer, or a file, or a websocket - `Greet()` doesn't care, it's taking advantage of the fact that they'll all have the same interface as `io.Writer`.

We even ran the thing as a simple web service!  (Hey, remember how to actually run youir go program? In this case, it was `go run greet.go`.)

Anyway, the [Go standard library](https://pkg.go.dev/std) has all kinds of examples of DI and you should check it out because it's important and good and will make you a better developer and also butter your toast for you. You may also want to check out [The Go Standard Library Cookbook](https://www.packtpub.com/product/go-standard-library-cookbook/9781788475273)
