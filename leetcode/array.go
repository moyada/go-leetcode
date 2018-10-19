package main

import (
	"fmt"
	"sort"
)

func main() {
	//arr := []int{0, 1, 0, 3, 0, 9}
	//fmt.Println(arr)
	//fmt.Println(removeDuplicates(arr))
	//fmt.Println(maxProfit(arr))
	//rotate(arr, 6)
	//fmt.Println(containsDuplicate(arr))
	//fmt.Println(singleNumber(arr))
	//fmt.Println(intersect(arr, []int{2, 2}))
	//fmt.Println(plusOne(arr))
	//moveZeroes(arr)
	//fmt.Println(arr)

	//matrix := [][]int{
	//	{5, 1, 9,11},
	//	{2, 4, 8,10},
	//	{13, 3, 6, 7},
	//	{15,14,12,16}}
	//rotateMatrix(matrix)

	//board := [][]byte{
	//	{'.','.','.','.','5','.','.','1','.'},
	//	{'.','4','.','3','.','.','.','.','.'},
	//	{'.','.','.','.','.','3','.','.','1'},
	//	{'8','.','.','.','.','.','.','2','.'},
	//	{'.','.','2','.','7','.','.','.','.'},
	//	{'.','1','5','.','.','.','.','.','.'},
	//	{'.','.','.','.','.','2','.','.','.'},
	//	{'.','2','.','9','.','.','.','.','.'},
	//	{'.','.','4','.','.','.','.','.','.'}}
	//fmt.Println(isValidSudoku(board))
	fmt.Println(getFactorial(3))
}

/**
给定一个有序数组，你需要原地删除其中的重复内容，使每个元素只出现一次,并返回新的长度。

不要另外定义一个数组，您必须通过用 O(1) 额外内存原地修改输入的数组来做到这一点。
 */
func removeDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	tmp := nums[0]
	index := 1
	for i := 1; i < length; i = i + 1 {
		// 与上一个非重复数值不同
		if tmp != nums[i] {
			nums[index] = nums[i]
			tmp = nums[i]
			index += 1
		}
	}
	return index
}

/**
假设有一个数组，它的第 i 个元素是一个给定的股票在第 i 天的价格。

设计一个算法来找到最大的利润。你可以完成尽可能多的交易（多次买卖股票）。然而，你不能同时参与多个交易（你必须在再次购买前出售股票）。
 */
func maxProfit(prices []int) int {
	length := len(prices)
	if length <= 0 {
		return 0
	}

	count := 0
	for i := 1; i < length; i++ {
		if prices[i] - prices[i-1] > 0 {
			count += prices[i] - prices[i-1];
		}
	}
	return count
}

/**
将包含 n 个元素的数组向右旋转 k 步。

例如，如果  n = 8 ,  k = 3，给定数组  [1,2,3,4,5,6,7,8]  ，向右旋转后的结果为 [6,7,8,1,2,3,4,5]。
 */
func rotate(nums []int, k int)  {
	length := len(nums)
	if length <= 1 || k == 0 || k == length {
		return
	}

	var tmp1, tmp2, nextIndex int
	last := nums[length -1]
	for index := 0; last == nums[length -1]; {
		nextIndex = index
		tmp1 = nums[index]
		for true {
			tmp2 = nums[(nextIndex + k) % length]
			nums[(nextIndex + k) % length] = tmp1
			nextIndex = (nextIndex + k) % length
			tmp1 = tmp2
			if index == nextIndex {
				index++
				break
			}
		}
	}

	fmt.Println(nums)
}

/**
给定一个整数数组，判断是否存在重复元素。

如果任何值在数组中出现至少两次，函数应该返回 true。如果每个元素都不相同，则返回 false。
 */
func containsDuplicate(nums []int) bool {
	length := len(nums)
	if length <= 1 {
		return false
	}

	sort.Ints(nums)
	for index := 1; index < length; index++ {
		if nums[index] == nums[index - 1] {
			return true
		}
	}
	return false
}

/**
给定一个整数数组，除了某个元素外其余元素均出现两次。请找出这个只出现一次的元素。
 */
func singleNumber(nums []int) int {
	for _, v := range nums[1:] {
		fmt.Println(v)
		nums[0] ^= v
	}
	return nums[0]
}

/**
给定两个数组，写一个方法来计算它们的交集。

例如:
给定 nums1 = [1, 2, 2, 1], nums2 = [2, 2], 返回 [2, 2].
 */
func intersect(nums1 []int, nums2 []int) []int {
	len1, len2 := len(nums1), len(nums2);
	if(len1 <= 0) {
		return []int{}
	}
	if(len2 <= 0) {
		return []int{}
	}

	if len1 > len2 {
		return findIntersect(nums2, nums1, len2, len1)
	} else {
		return findIntersect(nums1, nums2, len1, len2)
	}
}

func findIntersect(nums1 []int, nums2 []int, length int, length2 int) []int {
	var mark = make([]int, length2)
	var intersect = make([]int, length)
	index := 0
	sort.Ints(nums1)
	for _, v1 := range nums1 {
		for i, v2 := range nums2 {
			if v1 == v2 && mark[i] == 0 {
				intersect[index] = v1
				mark[i] = 1
				index++
				break
			}
		}
	}
	return intersect[:index]
}

/**
给定一个非负整数组成的非空数组，给整数加一。

可以假设整数不包含任何前导零，除了数字0本身。

最高位数字存放在列表的首位。
 */
func plusOne(digits []int) []int {
	length := len(digits)

	count := 0
	for i := length - 1; i >= 0; i-- {
		if isPlus(digits[i]) {
			digits[i] = 0
			count++
			continue
		}
		digits[i] += 1
		break
	}
	if count == length {
		digits = make([]int, length + 1)
		digits[0] = 1
		return digits
	}
	return digits
}

func isPlus(num int) bool {
	return num == 9
}


/**
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序
输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
 */
func moveZeroes(nums []int)  {
	index := 0

	for k, v := range nums {
		if v != 0 {
			if index == k {
				index++
				continue
			}
			nums[index] = v
			nums[k] = 0
			index++
		}
	}
}

/**
给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。

你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
 */
func twoSum(nums []int, target int) []int {
	length := len(nums)
	arr := make([]int, 2)
	for i := 0 ; i < length ; i++ {
		for j := i+1 ; j < length ; j++ {
			if nums[i] + nums[j] == target {
				arr[0] = i
				arr[1] = j
				break
			}
		}
	}
	return arr
}

/**
判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
 */
func isValidSudoku(board [][]byte) bool {
	mark := make([]bool, 9)

	var item byte
	for i := 0 ; i < 9 ; i++ {
		for j := 0 ; j < 9 ; j++ {
			item = board[i][j]
			if item == '.' {
				continue
			}
			if mark[item - '1'] == true {
				return false
			}
			mark[item - '1'] = true
		}
		mark = make([]bool, 9)
	}

	for i := 0 ; i < 9 ; i++ {
		for j := 0 ; j < 9 ; j++ {
			item = board[j][i]
			if item == '.' {
				continue
			}
			if mark[item - '1'] == true {
				return false
			}
			mark[item - '1'] = true
		}
		mark = make([]bool, 9)
	}

	for i := 0 ; i < 9 ; i = i+3 {
		for j := 0 ; j < 9 ; j = j+3 {
			for x := 0 ; x < 3 ; x++ {
				for y := 0 ; y < 3 ; y++ {
					item = board[i+x][j+y]
					if item == '.' {
						continue
					}
					if mark[item - '1'] == true {
						return false
					}
					mark[item - '1'] = true
				}
			}
			mark = make([]bool, 9)
		}
	}

	return true
}

/**

给定一个 n × n 的二维矩阵表示一个图像。

将图像顺时针旋转 90 度。
 */
func rotateMatrix(matrix [][]int)  {
	length := len(matrix)

	half := length/2
	max := length-1
	var subLen int
	var tmp int
	var index int
	for i := 0; i < half; i++ {
		subLen = max - i
		for j := i; j < subLen; j++ {
			index = max - j
			tmp = matrix[i][j]
			matrix[i][j] = matrix[index][i]

			matrix[index][i] = matrix[subLen][index]

			matrix[subLen][index] = matrix[j][subLen]

			matrix[j][subLen] = tmp
		}
	}
}

func getFactorial(num int) int {
	total := 1
	for ; num > 1 ; num-- {
		total *= num
	}
	return total
}