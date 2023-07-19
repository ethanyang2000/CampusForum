package utils

type Counter struct {
	count int
}

func (c *Counter) Gen() int {
	c.count += 1
	return c.count - 1
}

func NewCounter() *Counter {
	return &Counter{
		count: 0,
	}
}
