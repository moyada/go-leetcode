package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

const (
	bit  = 1 << (10 * iota)
	KB   = 1 << (10 * iota)
	MB   = 1 << (10 * iota)
	GB   = 1 << (10 * iota)
	name = "intro"
)

var (
	max, step = 10, rand.Intn(10)
)

func main() {
	now := time.Now()

	defer fmt.Println(now)

	defer fmt.Println("Hello", now.Weekday(), runtime.GOOS)

	var count = 0
	for i := 0; i <= max; i = i + step {
		fmt.Printf("increase %d\n", i)
		count++
	}

	switch {
	case count == 1:
		fmt.Println("count == 1")
	case count > 1:
		fmt.Println("count > 1")
		fallthrough
	case count == 3:
		fmt.Println("count == 3")
	case count > 3:
		fmt.Println("count == 3")
	default:
		fmt.Println("...")
	}

	fmt.Println(name, KB)
}
