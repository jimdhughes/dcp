package main

import "log"

// ConstantStack is a stack that implements push, pop, and max
// and returns values in constant time
type ConstantStack struct {
	values []int
	maximums []int
}

func (c *ConstantStack) push(i int) {
	idx := len(c.values) -1
	c.values = append(c.values, i)
	if idx == -1 {
		c.maximums = append(c.maximums, i)
		return
	}
	if c.maximums[idx] > i {
		c.maximums = append(c.maximums, c.maximums[idx])
		return
	}
	c.maximums = append(c.maximums, i)
	return
}

func (c *ConstantStack) pop() int {
	idx := len(c.values) -1
	i := c.values[idx]
	c.values = c.values[:len(c.values)-1]
	c.maximums = c.maximums[:idx]
	return i
}

func (c *ConstantStack) max() int {
	idx := len(c.values) -1
	return c.maximums[idx]
}

func main() {
	log.Println("Hello, World")
}