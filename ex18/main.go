package main

import (
	"fmt"
	"sync"
)

// Counter структура, которая хранит значение счетчика и mutex, ограничивающий доступ к нему
type Counter struct {
	value int
	mutex sync.Mutex
}

// Increase блокирует доступ к счетчику другим горутинам во время изменения его значения
func (c *Counter) Increase() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.value++
}

func main() {
	var counter Counter   // счетчик
	var wg sync.WaitGroup // структура для ожидания завешения горутин

	// запускаем горутины, увеличивающие значение счетчика
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(c *Counter) {
			for j := 0; j < 1000; j++ {
				c.Increase()
			}
			wg.Done()
		}(&counter)
	}

	// ожидаем завершения горутин
	wg.Wait()

	// проверяем значение счетчика
	fmt.Println(counter.value)
}
