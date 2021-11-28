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

Our second version iterates over the number of fields in the interface

```go
func walk(x interface{}, fn func(input string)) {
  // The value of x here is the actual value of the passed variable.
	val := reflect.ValueOf(x)

  // Here we iterate over its fields...
	for i := 0; i < val.NumField(); i++ {
    // ...and call the passed function on each 
		field := val.Field(i)
		fn(field.String())
	}
}
```

### Version 3: works for multiple string and non-string (but still flat) values


Version 3 lets us throw in an int or somthing and assures us that it'll be skipped over.

```go
func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

    // field.Kind() gets us the kind of field it is, durr.
		if field.Kind() == reflect.String {
			fn(field.String())
		}
	}
}
```

### Version 4: works for nested structs!

Note in this version that if we're passing a struct, we have to cast it to an Interface.

```go
func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		if field.Kind() == reflect.Struct {
			walk(field.Interface(), fn)
		}
		if field.Kind() == reflect.String {
			fn(field.String())
		}
	}
}
```