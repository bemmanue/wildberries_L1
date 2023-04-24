package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	<-time.After(duration)
}

func sleep2(duration time.Duration) {
	timer := time.NewTimer(duration)

	<-timer.C
}

func main() {
	fmt.Println("Ожидаем 5 секунд")
	sleep(5 * time.Second)
	fmt.Println("Время истекло")

	fmt.Println("Ожидаем 5 секунд")
	sleep2(5 * time.Second)
	fmt.Println("Время истекло")
}
