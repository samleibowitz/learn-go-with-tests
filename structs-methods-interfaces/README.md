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
## Interfaces

An interface is Go's solution to maintaining type-safety, while still being flexible enough to support methods that accept multiple types. In other words, it's a way to support polymorphism without having to resort to class hierarchies. You can think of this like a slightly more explicit version of duck typing. So when we declare an interface like this:

```go
type Shape interface {
	Area() float64
}
```

...we're saying "A `Shape` is a type which supports an `Area` method of type `float64`." 

So when we call it with something like `shape.Area()`, we're literally saying "I may not know specifically what type `shape` is, but I know it has to support the Area method."

