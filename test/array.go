package main

import "fmt"

func main() {
	a := 1
	b := [...]int{1, 2}
	c := [2]int{0: 2, 1: 3}
	fmt.Println(b == c)
	// 指向数组的指针
	var p1 *[len(b)]int = &b
	// 指针数组
	var p2 = [...]*int{&a}
	fmt.Println(p1)
	fmt.Println(p2)

	for i, v := range b {
		fmt.Printf("index %d, value %v.\n", i, v)
	}
}
