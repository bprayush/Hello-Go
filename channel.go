package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

// Channel test
func Channel() {
	ch := make(chan int, 9)
	wg.Add(2)

	go func(ch <-chan int) {
		for i := range ch {
			fmt.Printf("%v %T \n", i, i)
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 1
		ch <- 2
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
}
