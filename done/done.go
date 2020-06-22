package main

import (
	"fmt"
	"time"
)

func feed(max int, done <-chan struct{}) <-chan int {
	output := make(chan int)

	go func() {
		defer close(output)
		n := 0
		for {
			select {
			case <-done:
				fmt.Println("Received signal to close")
				return
			default:
				output <- n
				n++
				if n >= max {
					return
				}
			}
		}
	}()

	return output
}

func main() {
	done := make(chan struct{})

	go func() {
		time.Sleep(time.Nanosecond * 500)
		close(done)
	}()

	for i := range feed(1000, done) {
		fmt.Println(i)
	}
}
