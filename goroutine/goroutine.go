package main

import (
	"fmt"
	"runtime"
	"sync"
	"unsafe"
)

func main() {

	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func(r int) {
		for i := 0; i < 10; i++ {
			fmt.Println("Routing", r, "iteration", i)
		}

		wg.Done()
	}(1)

	go func(r int) {
		for i := 0; i < 10; i++ {
			fmt.Println("Routing", r, "iteration", i)
		}

		wg.Done()
	}(2)

	wg.Wait()
}

type hchan struct {
	// Fields removed
	sendx uint
	recvx uint
	recvq waitq
	sendq waitq
	lock  mutex
}

type waitq struct {
	first *sudog
	last  *sudog
}

type sudog struct {
	// Fields removed
	next *sudog
	prev *sudog
	elem unsafe.Pointer
}
