package main

import (
	"fmt"
	"time"
)

// Guarded Command Language

func main() {
	c := make(chan int)

	go func() {
		fmt.Println("Sleeping input")
		time.Sleep(time.Second * 3)
		fmt.Println("Writing input")
		c <- 1
	}()

	fmt.Println("Waiting for inpout")
	fmt.Println("Input received:", <-c)
}
