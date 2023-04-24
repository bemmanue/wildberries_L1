package main

import "fmt"

// writeNum записывает в канал числа из среза и затем закрывает канал
func writeNum(arr []int, ch chan<- int) {
	for _, n := range arr {
		ch <- n
	}

	close(ch)
}

// writePow читает числа из канала ch1 и записывает квадрат числа в канал ch2
func writePow(ch1 <-chan int, ch2 chan<- int) {
	for {
		n, ok := <-ch1

		// завершается в случае закрытия 1 канала
		if ok != true {
			close(ch2)
			return
		}

		// записывает квадрат числа в канал
		ch2 <- n * n
	}
}

func main() {
	arr := []int{1, 4, 9, 3, 2, 7, 8, 6}

	ch1 := make(chan int) // канал для записи и чтения чисел из среза arr
	ch2 := make(chan int) // канал для записи и чтения квадрата числа

	// запускаем горутины
	go writeNum(arr, ch1)
	go writePow(ch1, ch2)

	for {
		n, ok := <-ch2

		// завершается в случае, когда канал ch2 закрыт
		// (все дочерние горутины к этому моменту уже завершены)
		if ok != true {
			return
		}

		fmt.Println(n)
	}
}
