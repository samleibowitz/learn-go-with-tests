# Reflection

> Reflection in computing is the ability of a program to examine its own structure, particularly through types; it's a form of metaprogramming. It's also a great source of confusion.

Rails (and a lot of Ruby programming generally) make heavy use of reflection, so this should be right up my alley. Maybe.

> In short only use reflection if you really need to.

Oh.

## Interfaces

```go
// From the io Package, here's the definition of io.Reader
// Any type that implements the below method has the Reader interface.
type Reader interface {
    Read(p []byte) (n int, err error)
}

// This is important because in the following examples, 
var r io.Reader
r = os.Stdin
r = bufio.NewReader(r)
r = new(bytes.Buffer)
// r is ALWAYS of type io.Reader, regardless of what the type is that any of those methods return.

```

## Problem discussion

### Version 1: works for a single field.

The task is, **write a function `walk(x interface{}, fn func(string))` which takes a struct x and calls fn for all strings fields found inside. difficulty level: recursively.**

Our first version just assumed there's a single field and that it's a stringifiable one.

```go
func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	field := val.Field(0)
	fn(field.String())
}
```

### Version 2: works for multiple string fields


```go
func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fn(field.String())
	}
}
```