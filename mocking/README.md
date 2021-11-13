# Mocking

We start of with a discussion about slicing up a feature into a series of small tests so that you can implement them individually. 

## The Mock in our example.

Here we're using DI to do our mocking, which... I don't love. In Ruby with Rspec, we'd have done something like

```ruby
  # util.rb
  class Util
    def sleep(duration)
      sleep(duration)
    end
  end

  # some spec somewhere
  allow(Util).to receive(:sleep).with(45)
```

Go is using full on dependency injection:

```go
# in countdown.go
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
...

# in test_countdown.go
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

```

I don't love this, but I guess it might just be a mindset change. IMO, if you're using DI _strictly_ for the sake of implementing test mocks, you're making your code hard to follow unnecessarily, given the existence of stuff like `allow` in Rspec.  

(And yes, I know that Rspec for Ruby, not Go, but you get my drift. )

## Spies

DI is also how we wind up being able to spy. We create a type that can accept calls to `.Sleep` and `.Write`, which satisfy the requirements for both of `Countdown`'s arguments. Then we just pass it twice, using it to record every time one of its methods gets called.

```go
type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}
```

It's... kinda fiendish. Again, this si a trip for someone like me who use used to Rspec giving us access to stuff like

```ruby
expect(foo).to receive(:bar).at_least(3).times
```

And now, some direct quotes:

> Mocking requires no magic and is relatively simple; using a framework can make mocking seem more complicated than it is. We don't use automocking in this chapter so that we get:
> * a better understanding of how to mock
> * practise implementing interfaces
> 
> In collaborative projects there is value in auto-generating mocks. In a team, a mock generation tool codifies consistency around the test doubles. This will avoid inconsistently written test doubles which can translate to inconsistently written tests.
> You should only use a mock generator that generates test doubles against an interface. Any tool that overly dictates how tests are written, or that use lots of 'magic', can get in the sea.

Huh. Well, that's a different approach.

## A little more generic foo

"Test Double" is the generic term for something that stands for somethign else in a test. Mocks, spies, stubs, etc all have a variety of meanings that can sorta get blended into each other in practice, but for a good discussion of what everything means, checkout out [this Martin Fowler post](https://martinfowler.com/bliki/TestDouble.html)
## Odds and Ends

I notice that Sleep seems to take milliseconds as its parameter by default. Huh.

