package main

import (
	"fmt"
)

// square записывает в канал квадрат числа
func square(n int, ch chan<- int) {
	ch <- n * n
}

func main() {
	// создаем и инициализируем массив
	arr := [5]int{2, 4, 6, 8, 10}

	// создаем буферизованный канал
	ch := make(chan int, len(arr))

	// создаем переменную для записи суммы квадратов чисел
	var result int

	// запускаем горутины, которые будут записывать квадрат числа в канал
	for _, n := range arr {
		go square(n, ch)
	}

	// читаем из канала результаты работы горутин
	for i := 0; i < cap(ch); i++ {
		result += <-ch
	}

	// закрываем канал
	close(ch)

	// печатаем результат
	fmt.Println(result)
}
