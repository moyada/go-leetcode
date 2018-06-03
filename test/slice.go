package main

import "fmt"

func main() {
	s := make([]int, 4)
	s[0] = 1
	sa := s[:3]

	sa = append(sa, 5)
	fmt.Println(sa)
	fmt.Println(s)

	sc := make([]int, 2)
	copy(sc, s)
	sc = append(sc, 2)

	fmt.Println(s)
	fmt.Println(sc)

	var list []int
	for _, v := range s {
		list = append(list, v)
	}
	fmt.Println(list)
}
