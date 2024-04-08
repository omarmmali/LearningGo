package counter

import "sync"

// Counter is a simple counter
type Counter struct {
	mu sync.Mutex
	value int
}

// Inc increments the counter by 1
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns the current value of the counter
func (c *Counter) Value() int {
	return c.value
}