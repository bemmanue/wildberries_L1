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
	var n int
	flag.IntVar(&n, "n", 2, "number of workers")
	flag.Parse()

	ch := make(chan int)
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT)

	wg := sync.WaitGroup{}
	wg.Add(n)

	for i := 0; i < n; i++ {
		go work(ch, &wg)
	}

	go func() {
		for {
			select {
			case ch <- rand.Intn(100):
				time.Sleep(time.Second)
			case <-done:
				close(ch)
				return
			}
		}
	}()

	wg.Wait()
}
