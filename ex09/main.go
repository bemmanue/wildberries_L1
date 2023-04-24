package main

import "fmt"

func writeNum(arr []int, ch chan<- int) {
	for _, n := range arr {
		ch <- n
	}

	close(ch)
}

func writePow(ch1 <-chan int, ch2 chan<- int) {
	for {
		n, ok := <-ch1
		if ok != true {
			close(ch2)
			return
		}
		ch2 <- n * n
	}
}

func main() {
	arr := []int{1, 4, 9, 3, 2, 7, 8, 6}

	ch1 := make(chan int)
	ch2 := make(chan int)

	go writeNum(arr, ch1)
	go writePow(ch1, ch2)

	for {
		n, ok := <-ch2
		if ok != true {
			return
		}
		fmt.Println(n)
	}
}
