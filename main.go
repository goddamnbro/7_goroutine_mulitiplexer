package main

import (
	"fmt"
)

func MergeTwoChannels(ch1, ch2 chan int) chan int {
	to := make(chan int, 100)

	go func() {
		for value := range ch1 {
			to <- value
		}
	}()

	go func() {
		for value := range ch2 {
			to <- value
		}
	}()

	// I'm not sure about this
	go func() {
		for {
			if len(ch1) == 0 && len(ch2) == 0 {
				close(to)
				break
			}
		}
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
