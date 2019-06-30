package main

import (
	"fmt"
	"sync"
)

func MergeTwoChannels(ch1, ch2 chan int) chan int {
	to := make(chan int, 100)

	var group sync.WaitGroup

	// reading from 1st chan
	group.Add(1)
	go func() {
		for value := range ch1 {
			to <- value
		}
		group.Done()
	}()

	// reading from 2nd chan
	group.Add(1)
	go func() {
		for value := range ch2 {
			to <- value
		}
		group.Done()
	}()

	// just waiting for finishing other goroutines
	go func() {
		group.Wait()
		close(to)
	}()

	return to
}

func main() {
	fromCh1 := make(chan int, 10)
	fromCh2 := make(chan int, 10)

	// fill channels
	go func() {
		for i := 0; i < 10; i++ {
			fromCh1 <- i
			fromCh2 <- i
		}
		close(fromCh1)
		close(fromCh2)
	}()

	mergedCh := MergeTwoChannels(fromCh1, fromCh2)

	// just printing
	for {
		value, ok := <-mergedCh
		if !ok {
			break
		}
		fmt.Println(value)
	}
}
