package main

import (
	"fmt"
	"time"
)

// sleep завершает работу после чтения из канала, возвращаемого функцией time.After(duration)
func sleep(duration time.Duration) {
	<-time.After(duration)
}

// sleep завершает работу после чтения из канала, возвращаемого объектом time.NewTimer(duration)
func sleep2(duration time.Duration) {
	timer := time.NewTimer(duration)

	<-timer.C
}

func main() {
	// вариант 1
	fmt.Println("Ожидаем 5 секунд")
	sleep(5 * time.Second)
	fmt.Println("Время истекло")

	// вариант 2
	fmt.Println("Ожидаем 5 секунд")
	sleep2(5 * time.Second)
	fmt.Println("Время истекло")
}
