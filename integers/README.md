# 2. Integers

Okay look, this bettter be more or less what you'd expect.

## 2.1 Write the test first

Had to remind myself to do `go mod init integers`, whoops.

Notice that multiple function arguments of the same time can be instantiated with a comma, as in `func Add(x, y int)`.

The addition of "examples" into test suites is prettty slick:

```go
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
```
