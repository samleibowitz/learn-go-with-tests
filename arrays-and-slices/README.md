# Arrays and Slices

## Arrays

You know what these are, yeah? Except you don't, because Arrays in Go have a fixed length! Yeah, you heard me right. A function expecting an argument of type `[5]int` wll throw up if you try to give it a `[4]int`. Good times.

You can iterate over an Array using the `range` keyword. It returns the index of each item of the array as an Integer. So in this example, 

```go
	for _, number := range numbers {
		total += number
	}
```

the `_` indicates the index of the number, which we're just throwing away here. If you don't use `_`, you'd have to use that variable name or get a compiler error.


## Slices

Just like arrays, but can take a variable number of items. (All of the same type, though, dur.)

Referred to similarly, just without the number identifying the count in the type declaration:

```go
my_array := [5]int{0, 1, 2, 3, 4, 5}
my_slice := []int{0, 1, 2, 3}
```

## A couple of words about tests

The author introduces the built-in test coverage tool, which is run with the `-cover` argument:

```
❯ go test
PASS
ok      arrays-and-slices/arrays-and-slices     0.099s
❯ go test -cover
PASS
coverage: 100.0% of statements
ok      arrays-and-slices/arrays-and-slices     0.631s
```

Honestly not sure why you wouldn't run it with the -cover option all the time.
