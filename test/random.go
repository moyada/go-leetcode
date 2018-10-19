package main

import (
	"fmt"
	"math"
)

const (
	intro = `hello world`
)

var (
	time = 0
)

func main() {
	// var name string = `xyk`

	fmt.Println(intro)
	fmt.Println("next %g", math.Nextafter(3, 5))
	fmt.Println("next %g", math.Nextafter(3, 5))
	// fmt.Println(len(name))
	// fmt.Println(name[0])
	// time++
}

func randomArr(length int) (arr []int) {

	if length <= 0 {
		// arr = [0]int{}
		return arr
	}

	// arr = [length]int{}
	return arr

	// if (length <= 0) {
	//            return new int[0];
	//        }
	//        Random random = new Random();
	//        int[] result = new int[length];
	//        //i从0开始，下面的初始化可以去掉，但是这样节省一次遍历,可以在性能和可读性权衡
	//        result[0] = 0;
	//        for (int i = 1; i < length; i++) {
	//            Integer randomIndex = random.nextInt(i + 1);
	//            result[i] = result[randomIndex];
	//            result[randomIndex] = i;
	//        }
	//        return result;
}
