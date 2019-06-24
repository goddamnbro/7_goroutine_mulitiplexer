package main

import (
	"fmt"
	"sync"
)

func MergeTwoChannels(ch1, ch2 chan int) chan int {
	to := make(chan int, 100)

	var group sync.WaitGroup

	group.Add(1)
	go func() {
		for value := range ch1 {
			to <- value
		}
		group.Done()
	}()

	group.Add(1)
	go func() {
		for value := range ch2 {
			to <- value
		}
		group.Done()
	}()

	group.Wait()
	return to
}

func PrintChannelValues(ch chan int) []int {
	var values []int
	for {
		value, ok := <-ch
		if !ok {
			break
		}
		values = append(values, value)
	}
	return values
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
	close(mergedCh)

	fmt.Println("-->", PrintChannelValues(mergedCh))
}
