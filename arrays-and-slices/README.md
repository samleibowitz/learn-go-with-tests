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

We can use `len`  to get the length of a slice (wouldn't have to for an array, since they're of fixed length in Go).  And we can use `make` to generate a new one, like the following:

```go
    lengthOfNumbers := len(numbersToSum)
    sums := make([]int, lengthOfNumbers)
```

`make` is a built-in function that works to construct slices, maps, or chans. 

We can use `append` to add stuff to the end of a slice, although it is technically returning a copy of the original slice (with the new thingy tacked on)
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

## More about Testing and Comparisons

We used the `reflect` package to be able to compare two slices (using `reflect.DeepEqual`). It's not type safe, so in theory we could do `reflect.DeepEqual("freddie", 6)` and the compiler won't complain.




