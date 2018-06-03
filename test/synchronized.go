package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	synchronized()
}

func synchronized() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go countSync(&wg, i)
	}

	wg.Wait()
}

func countSync(wg *sync.WaitGroup, index int) {
	a := 0
	for i := 0; i < 1000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	wg.Done()
}
