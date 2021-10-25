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

