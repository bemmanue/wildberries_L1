package main

import (
	"fmt"
	"sync"
)

// Data структура, хранящая map и mutex для конкурентного доступа к map
type Data struct {
	data  map[int]int
	mutex sync.Mutex
}

// NewData конструктор Data
func NewData() *Data {
	return &Data{
		data: map[int]int{},
	}
}

// Write с помощью mutex блокирует доступ к map другим горутинам
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

	// запускаем горутины, записывающие в канал
	for i := 0; i < n; i++ {
		go func(key int) {
			data.Write(key) // обеспечивает безопасную операцию в конкурентной среде
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println(data.data)
}
