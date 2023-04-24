package main

import (
	"fmt"
	"math/rand"
	"time"
)

func read(ch <-chan int) {
	for {
		fmt.Printf("Value was read: %d\n", <-ch)
		time.Sleep(time.Second)
	}
}

func write(ch chan<- int) {
	for {
		n := rand.Intn(100)
		ch <- n
		fmt.Printf("Value was written: %d\n", n)
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int)

	go read(ch)
	go write(ch)

	time.Sleep(5 * time.Second)
}
