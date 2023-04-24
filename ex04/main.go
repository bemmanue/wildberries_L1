package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// work читает данные из канала и печатает результат в терминал
func work(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		n, ok := <-ch
		if ok != true {
			return
		}
		fmt.Println(n)
	}
}

func main() {
	// создаем переменную для обозначения количества воркеров
	var n int

	// определяем флаг, чтобы можно было задавать количество воркеров при запуске программы
	flag.IntVar(&n, "n", 2, "number of workers")
	flag.Parse()

	// создаем канал для записи и чтения числовых значений
	ch := make(chan int)

	// создаем канал для получения сигнала завершения
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT)

	// создаем структуру WaitGroup, чтобы заставить главную горутину ожидать завершения запущенных в ней горутин
	wg := sync.WaitGroup{}
	wg.Add(n)

	// запускаем горутины, считывающие данные из канала ch
	for i := 0; i < n; i++ {
		go work(ch, &wg)
	}

	// запускаем горутину, записывающую в канал ch рандомные числовые значения
	go func() {
		for {
			select {
			case ch <- rand.Intn(100):
				time.Sleep(time.Second)
			case <-done: // ожидаем получения сигнала завершения
				close(ch)
				return
			}
		}
	}()

	// ожидаем завершения запущенных горутин
	wg.Wait()
}
