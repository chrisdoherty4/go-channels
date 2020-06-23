package main

import "fmt"

func add(input <-chan int, addend int) <-chan int {
	output := make(chan int)

	go func() {
		for {
			select {
			case v := <-input:
				output <- v + addend
			}
		}
	}()

	return output
}

func multiply(input <-chan int, multiplier int) <-chan int {
	output := make(chan int)

	go func() {
		for {
			select {
			case v := <-input:
				output <- v * multiplier
			}
		}
	}()

	return output
}

func main() {
	input := make(chan int)
	output := add(add(add(input, 1), 2), 4)

	input <- 1
	fmt.Println(<-output)
}
