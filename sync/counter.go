package main

import "sync"

type Counter struct {
	mu  sync.Mutex
	val int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	c.val += 1
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	return c.val
}
