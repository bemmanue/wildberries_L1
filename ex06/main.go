package main

import (
	"context"
	"sync"
	"time"
)

func work1(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	<-ctx.Done()
}

func work2(done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	<-done
}

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

	//
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	go work1(ctx, &wg)

	//
	done := make(chan struct{})
	go work2(done, &wg)
	time.Sleep(time.Second)
	close(done)

	//
	ch := make(chan int)
	go work3(ch, &wg)
	for i := 0; i < 1000; i++ {
		ch <- 1
	}
	close(ch)

	wg.Wait()
}
