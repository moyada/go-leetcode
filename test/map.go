package main

import "fmt"

func main() {
	m := map[int]string{1: "base.go", 2: "array.go", 3: "slice.go", 4: "map.go"}

	for k, v := range m {
		fmt.Printf("key = %d, value = %v.\n", k, v)
	}

	_, exist := m[5]
	if !exist {
		m[5] = "func.go"
	}
	delete(m, 4)
	fmt.Println(m)
}
