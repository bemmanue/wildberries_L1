package main

import (
	"context"
	"sync"
	"time"
)

// work1 ожидает, когда канал из контекста закроется и завершается
func work1(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	<-ctx.Done()
}

// work2 также ожидает закрытия канала done и завершается
func work2(done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	<-done
}

// work3 читает данные из канала и завершается в случае, если канал закрывается
func work3(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		_, ok := <-ch
		if ok != true {
			return
		}
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	// 1 вариант - использование контекста
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	go work1(ctx, &wg)

	// 2 вариант - паттерн на основе канала done
	done := make(chan struct{})
	go work2(done, &wg)
	time.Sleep(time.Second)
	close(done)

	// 3 вариант - завершение из горутины в случае закрытия канала
	ch := make(chan int)
	go work3(ch, &wg)
	for i := 0; i < 1000; i++ {
		ch <- 1
	}
	close(ch)

	wg.Wait()
}
