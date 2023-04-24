package main

import (
	"fmt"
	"math/rand"
	"time"
)

// read читает данные из канала и печатает результат в терминал
func read(ch <-chan int) {
	for {
		fmt.Printf("Value was read: %d\n", <-ch)
		time.Sleep(time.Second) // замедляем для удобного отслеживания работы
	}
}

// write записывает в канал рандомные числовые значаения
func write(ch chan<- int) {
	for {
		n := rand.Intn(100)
		ch <- n
		fmt.Printf("Value was written: %d\n", n)
		time.Sleep(time.Second) // замедляем для удобного отслеживания работы
	}
}

func main() {
	// создаем канал для чтения и записи числовых значений
	ch := make(chan int)

	go read(ch)  // запускаем горутину, читающую из канала
	go write(ch) // запускаем горутину, записывающую в канал

	// ждем 5 секунд для завершения программы
	time.Sleep(5 * time.Second)
}
