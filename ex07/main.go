package main

import (
	"fmt"
	"sync"
)

type Data struct {
	mutex sync.Mutex
	data  map[int]int
}

func NewData() *Data {
	return &Data{
		data: map[int]int{},
	}
}

func (msm *Data) Write(key int) {
	msm.mutex.Lock()
	defer msm.mutex.Unlock()

	msm.data[key]++
}

func main() {
	data := NewData()
	n := 5

	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(key int) {
			data.Write(key)
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println(data.data)
}
