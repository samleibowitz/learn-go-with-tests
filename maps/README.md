# Maps

Thinks Hashes or Dictionaries, depending on your language of choice.

## Weird quirky stuff

Maps never have to be dereferenced. That's because the type is itself "a pointer to a runtime.hmap structure", which sounds weird to me. Like, what if you want it frozen? Weird.

Anyway, maps can have a nil value, but if you try to assign something to it it'll blow up.

```go
var m map[string]string // DON'T DO THIS

//Do one of these instead:
var dictionary = map[string]string{}
// OR
var dictionary = make(map[string]string)
```

Yeah, it's weird. Anyway, more about maps [here](https://blog.golang.org/go-maps-in-action)

## Other stuff about this section.

At one point we wind up with a switch statement to handle whether we return an error or not in our `Add` method:

```go
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
```

This really bugs me for reasons I'm not sure I understand. I suppose I'm in the habit of thinking that switch statements are a code smell.

