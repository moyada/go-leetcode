package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// selector()

	// channel()

	// send()

	timeout()
}

func timeout() {
	c := make(chan bool)

	select {
	case v := <-c:
		fmt.Println(c)
	case <-time.After(3 * time.Second):
		fmt.Println("Time Out")
	}
}

func send() {
	i := make(chan int)
	go func() {
		for v := range i {
			fmt.Println(v)
		}
	}()

	for {
		select {
		case i <- 0:
		case i <- 1:
		}
	}
}

func selector() {
	c1, c2, o := make(chan int), make(chan string), make(chan bool, 2)

	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c1", v)

			case v, ok := <-c2:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c2", v)
			}
		}
	}()

	c1 <- 1
	c2 <- "hello"
	c1 <- 3
	c2 <- "88"
	close(c1)
	close(c2)

	for i := 0; i < 2; i++ {
		<-o
	}
}

func channel() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go countChannel(c, i)
	}

	for i := 0; i < 10; i++ {
		<-c
	}
	close(c)
}

func countChannel(c chan bool, index int) {
	a := 0
	for i := 0; i < 1000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	c <- true
}
