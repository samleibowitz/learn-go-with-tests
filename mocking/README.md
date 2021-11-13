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


## Odds and Ends

I notice that Sleep seems to take milliseconds as its parameter by default. Huh.

