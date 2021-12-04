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

## The Assignment

**Write a function `walk(x interface{}, fn func(string))` which takes a struct x and calls fn for all strings fields found inside. difficulty level: recursively.**

### Version 1: works for a single field.

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

    // cast the struct to an empty Interface. (A little weird, but okay)
		if field.Kind() == reflect.Struct {
			walk(field.Interface(), fn)
		}
		if field.Kind() == reflect.String {
			fn(field.String())
		}
	}
}
```

### Version 5: works on pointers!

Made it work with just the addition of:

```go
  // Reflect's Kind() function returns a string with 
	// the type of the referent. reflect.Ptr is just 
	// a constant.
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
```

We refactored it to move the pointer dereferencing stuff into its own function.

### Version 6: Okay, but what if we get a slice?

We added this bit:
```go
  // if the thing we got passed is a slice
	if val.Kind() == reflect.Slice {
		// iterate over the slice and then
		for i := 0; i < val.Len(); i++ {
			// cast each element of thee slice to an empty Interface and call walk() on it.
			walk(val.Index(i).Interface(), fn)
		}
		return
	}
``` 

Which is gross, but still works. We eventually refactor it to make it less gross, though. We did it with some crazy shit.

```go
func walk(x interface{}, fn func(input string)) {
    val := getValue(x)

    numberOfValues := 0
		// Here, we're building a general purpose method to get values
		// from somthing inside a container. We create a varable called 
		// getField, which holds a reference to a function
    var getField func(int) reflect.Value

    switch val.Kind() {
    case reflect.String:
        fn(val.String())
    case reflect.Struct:
        numberOfValues = val.NumField()
				// if we're walking a Struct, getField will behave likee val.Field. (I mean,
				// it will actually BE val.Field)
        getField = val.Field
    case reflect.Slice:
        numberOfValues = val.Len()
				// And if it's a Slice, it'll behave like val.Index
        getField = val.Index
    }

    for i:=0; i< numberOfValues; i++ {
			  // Now we just call walk() with our magically adapted function getField.
        walk(getField(i).Interface(), fn)
    }
}
```

Adding support for Arrays isn't eveen worth mentioning.

### Version 7: Okay, what about Maps?

Our first iteration is kinda gross, because it relies on special behavior for the Map case - in that case, it's calling walk right in the switch statement. But it works and we're gonna commit it like that for now.

```golang
func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	numberOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field
	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		getField = val.Index
  // Here's the grody bit.
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	}

	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), fn)
	}
}
```

### Version 8: Let's clean up that map behavior a bit.

Here we're just fixing a test because we're not guaranteed the ordering of elements inside the map. We created a helper method called assertContains().

```golang
func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
```
