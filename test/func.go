package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("over.")
	}()

	fun1()
	f2 := fun2(10)
	fmt.Println(f2)
	f3 := fun3(10)
	fmt.Println(f3)
	f4m, f4s := fun4(3, 4, 5, 10, 13)
	fmt.Println(f4m, f4s)

	f5s := []int{1, 2, 3}
	fun5(f5s)
	fmt.Println(f5s)

	f6 := 10
	fun6(&f6)
	fmt.Println(f6)

	f7 := fun7
	fmt.Println(f7(10))

	f8 := fun8(10)
	fmt.Println(f8(5))

	fun9()
}

func fun1() {
	fmt.Println("func1")
}

func fun2(a int) int {
	return a + 1
}

func fun3(a int) (b int) {
	b = a + 1
	return
}

func fun4(s ...int) (max int, sum int) {
	max, sum = s[0], 0
	for _, v := range s {
		sum = sum + v
		if max < v {
			max = v
		}
	}
	return
}

func fun5(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}

func fun6(s *int) {
	*s = *s * 2
}

func fun7(a int) int {
	x, y := 1, 1

	for ; a > 0; a-- {
		x = x + y
		y = x
	}
	return x
}

func fun8(x int) func(int) int {
	return func(y int) int {
		fmt.Printf("closure %p\n", &x)
		return x + y
	}
}

func fun9() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("interrupt")
		}
	}()

	fmt.Println("start")
	panic("error occur")

	fmt.Println("done")
}
