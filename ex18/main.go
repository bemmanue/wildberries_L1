package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mutex sync.Mutex
	value int
}

func (c *Counter) Increase() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.value++
}

func main() {
	var counter Counter
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(c *Counter) {
			for j := 0; j < 1000; j++ {
				c.Increase()
			}
			wg.Done()
		}(&counter)
	}

	wg.Wait()

	fmt.Println(counter.value)
}
