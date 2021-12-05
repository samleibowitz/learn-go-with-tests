package main

type Counter struct {
	val int
}

func (c *Counter) Inc() {
	c.val += 1
}

func (c Counter) Value() int {
	return c.val
}
