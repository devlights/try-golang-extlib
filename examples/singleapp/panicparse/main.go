/*
maruel/panicparse の動作確認用のプログラム

２秒するとパニックします。
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(3)

	go func(ch chan<- int) {
		defer wg.Done()
		for i := range 100 {
			ch <- i
			time.Sleep(500 * time.Millisecond)
		}
	}(ch)
	go func(ch <-chan int) {
		defer wg.Done()
		for i := range ch {
			fmt.Println(i)
		}
	}(ch)
	go func(ch chan<- int) {
		defer wg.Done()
		defer close(ch)
		time.Sleep(2 * time.Second)
	}(ch)

	wg.Wait()

	return nil
}
