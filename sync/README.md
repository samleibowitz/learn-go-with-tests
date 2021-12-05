# Sync

The assignment: make a threadsafe counter. 

We'll start by making one which is manifestly unsafe, then fix it.

## Version One: thread unsafe.

Okay, so this one I guessed.

## Version Two: make it threadsafe

It's worth briefly going over the test itself in this version.

`sync.WaitGroup` tracks a number of goroutines then waits for them to be done. Here, we're going to start up 1000 goroutines, each one of which is going to increment our counter by one. So the result SHOULD be 1000... but without a thread safe counter, it almost certainlyu won't work.

```golang
	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

    // our sync.WaitGroup is told to wait for 1000 goroutines to report that they're done.
		var wg sync.WaitGroup
		wg.Add(wantedCount)

    // A for loop that spawns 1000 goroutines
		for i := 0; i < wantedCount; i++ {
			go func() {
        // increment the counter, then tell the waitgroup we're done
				counter.Inc()
				wg.Done()
			}()
		}
    // wait on all the threads to complete.
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
```

We use a `sync.Mutex` to ensure that two threads aren't trying to update the value at the same time.

```golang
type Counter struct {
	mu  sync.Mutex
	val int
}
// ...

func (c *Counter) Inc() {
	c.mu.Lock()
	c.val += 1
	c.mu.Unlock()
}
```

The addition of the mutex means that we'll get warnings any time we pass it by value, so from now on, Counters get passed by reference:

```
./counter_test.go:15:20: call of assertCounter copies lock value: arrays-and-slices/sync.Counter contains sync.Mutex
./counter_test.go:33:20: call of assertCounter copies lock value: arrays-and-slices/sync.Counter contains sync.Mutex
./counter_test.go:37:38: assertCounter passes lock by value: arrays-and-slices/sync.Counter contains sync.Mutex
```

## Additional stuff

1. `go vet` is the bees knees, and you should use it regularly.
2. Channels and Mutexes are two different ways of managing concurrency. Use channels to manage who owns what data, and mutexes to manage state.
3. Thing _very_ hard before using an embedded type like:
   ```golang
   type Counter struct {
     sync.Mutex
     value int
   }

   func (c *Counter) Inc() {
     c.Lock()
     defer c.Unlock()
     c.value++
   }
   ```
   ...as it means that now random consumers of your type can call it like `myCounter.Unlock()`. Super shady.

