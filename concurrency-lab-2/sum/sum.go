package main

import (
	"fmt"
	"sync"
)

func main() {
	sum := 0
	var wg sync.WaitGroup
	channel := make(chan int, 1000)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			channel <- 1
			wg.Done()
		}()
	}
	wg.Wait()
	L:
	for {
		select {
		case x := <-channel:
			sum += x
		default:
			break L
		}
	}

	fmt.Println(sum)
}
