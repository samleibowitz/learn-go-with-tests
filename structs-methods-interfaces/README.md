# Structs, Method, and Interfaces

## Structs

Structs are a way of structuring data. You've seen them before in Ruby, and here they so far seem pretty similar. Note the syntax:

```go
type Rectangle struct {
	Width  float64
	Height float64
}
```

 ## Methods

 It's so weird that you don't define method _inside_ a class, but rather identifying the type of a receiver. So to extend our example from before:

 ```go
 type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return (r.Width * r.Height)
}
```
