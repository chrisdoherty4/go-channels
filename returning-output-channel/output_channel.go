package main

import "fmt"

/*
func foo() <-chan int {
	output := make(chan int)

	go func() {
		// Do something that writes to `output`
	}()

	return output
}
*/

func feed(max int) <-chan int {
	output := make(chan int)

	go func() {
		defer close(output)
		for i := 0; i < max; i++ {
			output <- i
		}
	}()

	return output
}

func main() {
	for i := range feed(10) {
		fmt.Println(i)
	}
}
